// Copyright (c) 2020, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package helidon

import (
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	dump "github.com/verrazzano/verrazzano/tests/e2e/pkg/test/clusterdump"
	"io"
	"net/http"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/verrazzano/verrazzano/pkg/k8sutil"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg/test/framework"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg/test/framework/metrics"
)

const (
	longWaitTimeout          = 20 * time.Minute
	longPollingInterval      = 20 * time.Second
	shortPollingInterval     = 10 * time.Second
	shortWaitTimeout         = 5 * time.Minute
	imagePullWaitTimeout     = 40 * time.Minute
	imagePullPollingInterval = 30 * time.Second
	skipVerifications        = "Skip Verifications"
	helloHelidon             = "hello-helidon"
	nodeExporterJobName      = "node-exporter"
)

var (
	t                        = framework.NewTestFramework("helidon")
	generatedNamespace       = pkg.GenerateNamespace(helloHelidon)
	expectedPodsHelloHelidon = []string{"hello-helidon-deployment"}
	metricsTest              pkg.MetricsTest
)

var beforeSuite = t.BeforeSuiteFunc(func() {
	if !skipDeploy {
		start := time.Now()
		pkg.DeployHelloHelidonApplication(namespace, "", istioInjection, "", "")
		metrics.Emit(t.Metrics.With("deployment_elapsed_time", time.Since(start).Milliseconds()))
	}

	Eventually(func() bool {
		return pkg.ContainerImagePullWait(namespace, expectedPodsHelloHelidon)
	}, imagePullWaitTimeout, imagePullPollingInterval).Should(BeTrue())
	// Verify hello-helidon-deployment pod is running
	// GIVEN OAM hello-helidon app is deployed
	// WHEN the component and appconfig are created
	// THEN the expected pod must be running in the test namespace
	if !skipVerify {
		Eventually(helloHelidonPodsRunning, longWaitTimeout, longPollingInterval).Should(BeTrue())
	}

	kubeconfig, err := k8sutil.GetKubeConfigLocation()
	if err != nil {
		AbortSuite(fmt.Sprintf("Failed to get the Kubeconfig location for the cluster: %v", err))
	}
	metricsTest, err = pkg.NewMetricsTest(kubeconfig, map[string]string{})
	if err != nil {
		AbortSuite(fmt.Sprintf("Failed to create the Metrics test object: %v", err))
	}

	beforeSuitePassed = true
})

var _ = BeforeSuite(beforeSuite)

var failed = false
var beforeSuitePassed = false

var _ = t.AfterEach(func() {
	failed = failed || CurrentSpecReport().Failed()
})

var afterSuite = t.AfterSuiteFunc(func() {
	if failed || !beforeSuitePassed {
		dump.ExecuteBugReport(namespace)
	}
	if !skipUndeploy {
		start := time.Now()
		pkg.UndeployHelloHelidonApplication(namespace, "", "")
		metrics.Emit(t.Metrics.With("undeployment_elapsed_time", time.Since(start).Milliseconds()))
	}
})

var _ = AfterSuite(afterSuite)

var _ = t.Describe("Hello Helidon OAM App test", Label("f:app-lcm.oam",
	"f:app-lcm.helidon-workload"), func() {
	var host = ""
	var err error
	// Get the host from the Istio gateway resource.
	// GIVEN the Istio gateway for the hello-helidon namespace
	// WHEN GetHostnameFromGateway is called
	// THEN return the host name found in the gateway.
	t.BeforeEach(func() {
		Eventually(func() (string, error) {
			host, err = k8sutil.GetHostnameFromGateway(namespace, "")
			return host, err
		}, shortWaitTimeout, shortPollingInterval).Should(Not(BeEmpty()))
	})

	// Verify Hello Helidon app is working
	// GIVEN OAM hello-helidon app is deployed
	// WHEN the component and appconfig with ingress trait are created
	// THEN the application endpoint must be accessible
	t.Describe("for Ingress.", Label("f:mesh.ingress"), func() {
		t.It("Access /greet App Url w/o token and get RBAC denial", func() {
			if skipVerify {
				Skip(skipVerifications)
			}
			url := fmt.Sprintf("https://%s/greet", host)
			Eventually(func() bool {
				return appEndpointAccess(url, host, "", false)
			}, longWaitTimeout, longPollingInterval).Should(BeTrue())
		})

		t.It("Access /greet App Url with valid token", func() {
			if skipVerify {
				Skip(skipVerifications)
			}
			kc, err := pkg.NewKeycloakAdminRESTClient()
			Expect(err).To(BeNil())
			password := pkg.GetRequiredEnvVarOrFail("REALM_USER_PASSWORD")
			realmName := pkg.GetRequiredEnvVarOrFail("REALM_NAME")
			// check for realm
			_, err = kc.GetRealm(realmName)
			Expect(err).To(BeNil())
			var token string
			token, err = kc.GetToken(realmName, "testuser", password, "appsclient", t.Logs)
			Expect(err).To(BeNil())
			t.Logs.Debugf("Obtained token: %v", token)
			url := fmt.Sprintf("https://%s/greet", host)
			Eventually(func() bool {
				return appEndpointAccess(url, host, token, true)
			}, longWaitTimeout, longPollingInterval).Should(BeTrue())
		})
	})

	// Verify Prometheus scraped metrics
	// GIVEN OAM hello-helidon app is deployed
	// WHEN the component and appconfig without metrics-trait(using default) are created
	// THEN the application metrics must be accessible
	t.Describe("for Metrics.", Label("f:observability.monitoring.prom"), FlakeAttempts(5), func() {
		t.It("Retrieve Prometheus scraped metrics", func() {
			if skipVerify {
				Skip(skipVerifications)
			}
			pkg.Concurrently(
				func() {
					Eventually(appMetricsExists, longWaitTimeout, longPollingInterval).Should(BeTrue())
				},
				func() {
					Eventually(appComponentMetricsExists, longWaitTimeout, longPollingInterval).Should(BeTrue())
				},
				func() {
					Eventually(appConfigMetricsExists, longWaitTimeout, longPollingInterval).Should(BeTrue())
				},
				func() {
					Eventually(nodeExporterProcsRunning, longWaitTimeout, longPollingInterval).Should(BeTrue())
				},
				func() {
					Eventually(nodeExporterDiskIoNow, longWaitTimeout, longPollingInterval).Should(BeTrue())
				},
			)
		})
	})

	t.Context("Logging.", Label("f:observability.logging.es"), FlakeAttempts(5), func() {
		var indexName string
		Eventually(func() error {
			indexName, err = pkg.GetOpenSearchAppIndex(namespace)
			return err
		}, shortWaitTimeout, shortPollingInterval).Should(BeNil(), "Expected to get OpenSearch App Index")

		// GIVEN an application with logging enabled
		// WHEN the Opensearch index is retrieved
		// THEN verify that it is found
		t.It("Verify Opensearch index exists", func() {
			if skipVerify {
				Skip(skipVerifications)
			}
			Eventually(func() bool {
				return pkg.LogIndexFound(indexName)
			}, longWaitTimeout, longPollingInterval).Should(BeTrue(), "Expected to find log index for hello helidon")
		})

		// GIVEN an application with logging enabled
		// WHEN the log records are retrieved from the Opensearch index
		// THEN verify that at least one recent log record is found
		t.It("Verify recent Opensearch log record exists", func() {
			if skipVerify {
				Skip(skipVerifications)
			}
			Eventually(func() bool {
				return pkg.LogRecordFound(indexName, time.Now().Add(-24*time.Hour), map[string]string{
					"kubernetes.labels.app_oam_dev\\/name": helloHelidon,
					"kubernetes.container_name":            "hello-helidon-container",
				})
			}, longWaitTimeout, longPollingInterval).Should(BeTrue(), "Expected to find a recent log record")
			Eventually(func() bool {
				return pkg.LogRecordFound(indexName, time.Now().Add(-24*time.Hour), map[string]string{
					"kubernetes.labels.app_oam_dev\\/component": "hello-helidon-component",
					"kubernetes.labels.app_oam_dev\\/name":      helloHelidon,
					"kubernetes.container_name":                 "hello-helidon-container",
				})
			}, longWaitTimeout, longPollingInterval).Should(BeTrue(), "Expected to find a recent log record")
		})
	})

})

func helloHelidonPodsRunning() bool {
	result, err := pkg.PodsRunning(namespace, expectedPodsHelloHelidon)
	if err != nil {
		AbortSuite(fmt.Sprintf("One or more pods are not running in the namespace: %v, error: %v", namespace, err))
	}
	return result
}

func appEndpointAccess(url string, hostname string, token string, requestShouldSucceed bool) bool {
	req, err := retryablehttp.NewRequest("GET", url, nil)
	if err != nil {
		t.Logs.Errorf("Unexpected error=%v", err)
		return false
	}

	kubeconfigPath, err := k8sutil.GetKubeConfigLocation()
	if err != nil {
		t.Logs.Errorf("Unexpected error=%v", err)
		return false
	}

	httpClient, err := pkg.GetVerrazzanoHTTPClient(kubeconfigPath)
	if err != nil {
		t.Logs.Errorf("Unexpected error=%v", err)
		return false
	}

	if len(token) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	}

	req.Host = hostname
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Logs.Errorf("Unexpected error=%v", err)
		return false
	}
	bodyRaw, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Logs.Errorf("Unexpected error=%v", err)
		return false
	}
	if requestShouldSucceed {
		if resp.StatusCode != http.StatusOK {
			t.Logs.Errorf("Unexpected status code=%v", resp.StatusCode)
			return false
		}
		// HTTP Server headers should never be returned.
		for headerName, headerValues := range resp.Header {
			if strings.EqualFold(headerName, "Server") {
				t.Logs.Errorf("Unexpected Server header=%v", headerValues)
				return false
			}
		}
		bodyStr := string(bodyRaw)
		if !strings.Contains(bodyStr, "Hello World") {
			t.Logs.Errorf("Unexpected response body=%v", bodyStr)
			return false
		}
	} else {
		if resp.StatusCode == http.StatusOK {
			t.Logs.Errorf("Unexpected status code=%v", resp.StatusCode)
			return false
		}
	}
	return true
}

func appMetricsExists() bool {
	return metricsTest.MetricsExist("base_jvm_uptime_seconds", map[string]string{"app": helloHelidon})
}

func appComponentMetricsExists() bool {
	return metricsTest.MetricsExist("vendor_requests_count_total", map[string]string{"app_oam_dev_name": helloHelidon})
}

func appConfigMetricsExists() bool {
	return metricsTest.MetricsExist("vendor_requests_count_total", map[string]string{"app_oam_dev_component": "hello-helidon-component"})
}

func nodeExporterProcsRunning() bool {
	return metricsTest.MetricsExist("node_procs_running", map[string]string{"job": nodeExporterJobName})
}

func nodeExporterDiskIoNow() bool {
	return metricsTest.MetricsExist("node_disk_io_now", map[string]string{"job": nodeExporterJobName})
}
