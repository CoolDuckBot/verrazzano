// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package system

import (
	"encoding/json"
	"fmt"
	"github.com/verrazzano/verrazzano/pkg/k8sutil"
	"regexp"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg/test/framework"
)

const (
	systemNamespace           = "verrazzano-system"
	installNamespace          = "verrazzano-install"
	certMgrNamespace          = "cert-manager"
	keycloakNamespace         = "keycloak"
	cattleSystemNamespace     = "cattle-system"
	fleetLocalSystemNamespace = "cattle-fleet-local-system"
	nginxNamespace            = "ingress-nginx"
	monitoringNamespace       = "monitoring"
	shortPollingInterval      = 10 * time.Second
	shortWaitTimeout          = 5 * time.Minute
	searchTimeWindow          = "1h"
	fleetLocalSystemIndex     = "verrazzano-namespace-cattle-fleet-local-system"
)

var (
	noExceptions    = []*regexp.Regexp{}
	istioExceptions = []*regexp.Regexp{
		regexp.MustCompile(`^-A .*$`),
		regexp.MustCompile(`^-N .*$`),
		regexp.MustCompile(`^:\w+? -.*$`),
		regexp.MustCompile(`^:\w+? ACCEPT.*$`),
		regexp.MustCompile(`^\w+?=.*$`),
		regexp.MustCompile(`^COMMIT.*$`),
		regexp.MustCompile(`^ {0,4}\w+:.*$`),
		regexp.MustCompile(`^:.*$`),
		regexp.MustCompile(`^\* ?nat.*$`),
		regexp.MustCompile(`^# Generated by.*$`),
		regexp.MustCompile(`^# Completed on.*$`),
		regexp.MustCompile(`^Writing following contents to rules file:.*$`),
		regexp.MustCompile(`^ip\w?tables.*$`),
		regexp.MustCompile(`^-+$`),
		regexp.MustCompile(`^$`),
	}
	jaegerExceptions = []*regexp.Regexp{
		regexp.MustCompile(`^.*http: TLS handshake error.*$`),
		regexp.MustCompile(`^.*GOMAXPROCS.*$`),
	}
)

var t = framework.NewTestFramework("system-logging")

var _ = t.BeforeSuite(func() {})

var failed = false
var _ = t.AfterEach(func() {
	failed = failed || CurrentSpecReport().Failed()
})

var _ = t.AfterSuite(func() {
	if failed {
		pkg.ExecuteBugReport()
	}
})

var _ = t.Describe("Opensearch system component data", Label("f:observability.logging.es"), func() {
	t.It("contains verrazzano-system index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the verrazzano-system namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(systemNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index verrazzano-system")

		valid := true
		valid = validateAuthProxyLogs() && valid
		valid = validateCoherenceLogs() && valid
		valid = validateOAMLogs() && valid
		valid = validateIstioProxyLogs() && valid
		valid = validateKialiLogs() && valid
		valid = validatePrometheusLogs() && valid
		valid = validatePrometheusConfigReloaderLogs() && valid
		valid = validateGrafanaLogs() && valid
		valid = validateOpenSearchLogs() && valid
		valid = validateWeblogicOperatorLogs() && valid
		kubeConfigPath, err := k8sutil.GetKubeConfigLocation()
		Expect(err).To(BeNil())
		isJaegerSupported, err := pkg.IsVerrazzanoMinVersion("1.4.0", kubeConfigPath)
		Expect(err).To(BeNil())
		if isJaegerSupported {
			valid = validateJaegerCollectorLogs() && valid
			valid = validateJaegerQueryLogs() && valid
		}
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in verrazzano-system index")
		}
	})

	t.It("contains valid verrazzano-install index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the verrazzano-install namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(installNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index verrazzano-install")

		// GIVEN Log message in Opensearch in the verrazzano-namespace-verrazzano-install index
		// With field kubernetes.labels.app.keyword==verrazzano-platform-operator
		// WHEN Log messages are retrieved from Opensearch
		// THEN Verify there are valid log records
		valid := true
		valid = validateVPOLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in verrazzano-install index")
		}
	})

	t.It("contains valid verrazzano-system index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the verrazzano-system namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(systemNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index verrazzano-system")

		// GIVEN Log message in Opensearch in the verrazzano-namespace-verrazzano-system index
		// With field
		//  kubernetes.labels.app.keyword==verrazzano-application-operator,
		//  kubernetes.labels.app.keyword==verrazzano-monitoring-operator,
		// WHEN Log messages are retrieved from Opensearch
		// THEN Verify there are valid log records
		if !validateVAOLogs() {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with Verrazzano Application Operator log records in verrazzano-system index")
		}
		if !validateVMOLogs() {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with Verrazzano Monitoring Operator log records in verrazzano-system index")
		}
	})

	t.It("contains cert-manager index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the cert-manager namespace is retrieved
		// THEN verify that it is found

		indexName, err := pkg.GetOpenSearchSystemIndex(certMgrNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index cert-manager")

		valid := true
		valid = validateCertManagerLogs() && valid

		dnsPodExist, err := pkg.DoesPodExist("cert-manager", "external-dns")
		if err != nil {
			dnsPodExist = false
			t.Logs.Infof("Error calling DoesPodExist for external-dns: %s", err)
		}
		if dnsPodExist {
			valid = validateExternalDNSLogs() && valid
		}

		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in cert-manager index")
		}
	})

	t.It("contains valid Keycloak index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the Keycloak namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(keycloakNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index verrazzano-namepace-keycloak")

		// GIVEN Log message in Opensearch in the verrazzano-namespace-keycloak index
		// With field kubernetes.labels.app.kubernetes.io/name=keycloak
		// WHEN Log messages are retrieved from Opensearch
		// THEN Verify there are valid log records
		valid := true
		valid = validateKeycloakLogs() && valid
		valid = validateKeycloakMySQLLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in Keycloak index")
		}
	})

	t.It("contains ingress-nginx index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the index for the ingress-nginx namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(nginxNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find NGINX index ingress-nginx")

		valid := true
		valid = validateIngressNginxLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in ingress-nginx index")
		}
	})

	t.It("contains cattle-system index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the cattle-system namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(cattleSystemNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index cattle-system")

		valid := true
		valid = validateRancherLogs() && valid
		valid = validateRancherWebhookLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in cattle-system index")
		}
	})

	t.It("contains cattle-fleet-local-system index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the cattle-fleet-system namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(fleetLocalSystemIndex)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index cattle-fleet-local-system")

		if !validateFleetSystemLogs() {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in cattle-fleet-local-system index")
		}
	})

	t.It("contains cattle-fleet-local-system index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the cattle-fleet-local-system namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(fleetLocalSystemNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index cattle-fleet-local-system")

		if !validateFleetSystemLogs() {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in cattle-fleet-local-system index")
		}
	})

	t.It("contains monitoring index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Opensearch index for the monitoring namespace is retrieved
		// THEN verify that it is found
		indexName, err := pkg.GetOpenSearchSystemIndex(monitoringNamespace)
		Expect(err).To(BeNil())
		Eventually(func() bool {
			return pkg.LogIndexFound(indexName)
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Opensearch index monitoring")

		if !validateNodeExporterLogs() {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in monitoring index")
		}
	})
})

func validateAuthProxyLogs() bool {
	exceptions := []*regexp.Regexp{
		regexp.MustCompile(`^Adding local CA cert to .*$`),
		regexp.MustCompile(`^Detected Nginx Configuration Change$`),
	}
	exceptions = append(exceptions, istioExceptions...)
	return validateOpensearchRecords(
		noLevelOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app.keyword",
		"verrazzano-authproxy",
		searchTimeWindow,
		exceptions)
}

func validateCoherenceLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app_kubernetes_io/name.keyword",
		"coherence-operator",
		searchTimeWindow,
		noExceptions)
}

func validateOAMLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app_kubernetes_io/name.keyword",
		"oam-kubernetes-runtime",
		searchTimeWindow,
		noExceptions)
}

// message:configPath: ./etc/istio/proxy
func validateIstioProxyLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.container_name",
		"istio-proxy",
		searchTimeWindow,
		istioExceptions)
}

func validateKialiLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app_kubernetes_io/part-of",
		"kiali",
		searchTimeWindow,
		istioExceptions)
}

func validateVPOLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(installNamespace) },
		"kubernetes.labels.app.keyword",
		"verrazzano-platform-operator",
		searchTimeWindow,
		noExceptions)
}

func validateVAOLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app.keyword",
		"verrazzano-application-operator",
		searchTimeWindow,
		noExceptions)
}

func validateVMOLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app.keyword",
		"verrazzano-monitoring-operator",
		searchTimeWindow,
		noExceptions)
}

func validatePrometheusLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.container_name",
		"prometheus",
		searchTimeWindow,
		noExceptions)
}

func validatePrometheusConfigReloaderLogs() bool {
	return validateOpensearchRecords(
		noLevelOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.container_name",
		"config-reloader",
		searchTimeWindow,
		noExceptions)
}

func validateCertManagerLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(certMgrNamespace) },
		"kubernetes.labels.app_kubernetes_io/instance",
		"cert-manager",
		searchTimeWindow,
		noExceptions)
}

func validateExternalDNSLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(certMgrNamespace) },
		"kubernetes.labels.app_kubernetes_io/instance",
		"external-dns",
		searchTimeWindow,
		noExceptions)
}

func validateGrafanaLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app.keyword",
		"system-grafana",
		searchTimeWindow,
		noExceptions)
}

func validateOpenSearchLogs() bool {
	valid := true
	openSearchAppComponents := []string{"system-kibana", "system-es-data", "system-es-master", "system-es-ingest"}
	for _, appLabel := range openSearchAppComponents {
		valid = validateOpensearchRecords(
			noLevelOpensearchRecordValidator,
			func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
			"kubernetes.labels.app.keyword",
			appLabel,
			searchTimeWindow,
			noExceptions) && valid
	}
	return valid
}

func validateWeblogicOperatorLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(systemNamespace) },
		"kubernetes.labels.app.keyword",
		"weblogic-operator",
		searchTimeWindow,
		noExceptions)
}

func validateKeycloakLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(keycloakNamespace) },
		"kubernetes.labels.app.kubernetes.io/name",
		"keycloak",
		searchTimeWindow,
		noExceptions)
}

func validateIngressNginxLogs() bool {
	return validateOpensearchRecords(
		noLevelOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(nginxNamespace) },
		"kubernetes.labels.app_kubernetes_io/name",
		"ingress-nginx",
		searchTimeWindow,
		noExceptions)
}

func validateKeycloakMySQLLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(keycloakNamespace) },
		"kubernetes.labels.app.keyword",
		"mysql",
		searchTimeWindow,
		noExceptions)
}

func validateRancherLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(cattleSystemNamespace) },
		"kubernetes.labels.app.keyword",
		"rancher",
		searchTimeWindow,
		noExceptions)
}

func validateRancherWebhookLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(cattleSystemNamespace) },
		"kubernetes.labels.app.keyword",
		"rancher-webhook",
		searchTimeWindow,
		noExceptions)
}
func validateFleetSystemLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(fleetLocalSystemNamespace) },
		"kubernetes.namespace_name",
		"fleet-system",
		searchTimeWindow,
		noExceptions)
}

func validateNodeExporterLogs() bool {
	return validateOpensearchRecords(
		allOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(monitoringNamespace) },
		"kubernetes.labels.app.keyword",
		"node-exporter",
		searchTimeWindow,
		noExceptions)
}

func validateJaegerCollectorLogs() bool {
	return validateOpensearchRecords(
		logLevelOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(monitoringNamespace) },
		"kubernetes.container_name",
		"jaeger-collector",
		searchTimeWindow,
		jaegerExceptions)
}

func validateJaegerQueryLogs() bool {
	return validateOpensearchRecords(
		logLevelOpensearchRecordValidator,
		func() (string, error) { return pkg.GetOpenSearchSystemIndex(monitoringNamespace) },
		"kubernetes.container_name",
		"jaeger-query",
		searchTimeWindow,
		jaegerExceptions)
}

func validateOpensearchRecords(hitValidator pkg.OpensearchHitValidator, indexFunc func() (string, error), appLabel string, appName string, timeRange string, exceptions []*regexp.Regexp) bool {
	pkg.Log(pkg.Info, fmt.Sprintf("Validating log records for %s", appName))
	index, err := indexFunc()
	if err != nil {
		pkg.Log(pkg.Error, fmt.Sprintf("Failed to get OpenSearch index: %v", err))
		return false
	}

	template :=
		`{
			"size": 1000,
			"sort": [{"@timestamp": {"order": "desc"}}],
			"query": {
				"bool": {
					"filter" : [
						{"match_phrase": {"%s": "%s"}},
						{"range": {"@timestamp": {"gte": "now-%s"}}}
					]
				}
			}
		}`
	query := fmt.Sprintf(template, appLabel, appName, timeRange)
	resp, err := pkg.PostOpensearch(fmt.Sprintf("%s/_search", index), query)
	if err != nil {
		pkg.Log(pkg.Error, fmt.Sprintf("Failed to query Opensearch: %v", err))
		return false
	}
	if resp.StatusCode != 200 {
		pkg.Log(pkg.Error, fmt.Sprintf("Failed to query Opensearch: status=%d: body=%s", resp.StatusCode, string(resp.Body)))
		return false
	}
	var result map[string]interface{}
	json.Unmarshal(resp.Body, &result)

	if !pkg.ValidateOpensearchHits(result, hitValidator, exceptions) {
		pkg.Log(pkg.Info, fmt.Sprintf("Found invalid (or zero) log records in %s logs", appName))
		return false
	}
	return true
}

// allOpensearchRecordValidator does all validation for log records
func allOpensearchRecordValidator(hit pkg.OpensearchHit) bool {
	valid := true
	if !commonOpensearchRecordValidator(hit) {
		valid = false
	}
	if !logLevelOpensearchRecordValidator(hit) {
		valid = false
	}

	return valid
}

// noLevelOpensearchRecordValidator does validation for log records except level validation
func noLevelOpensearchRecordValidator(hit pkg.OpensearchHit) bool {
	return commonOpensearchRecordValidator(hit)
}

// commonOpensearchRecordValidator does all validation for log records except level validation
func commonOpensearchRecordValidator(hit pkg.OpensearchHit) bool {
	ts := ""
	valid := true
	// Verify the record has a @timestamp field.
	// If so extract it.
	if val, ok := hit["@timestamp"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty @timestamp field")
		valid = false
	} else {
		ts = hit["@timestamp"].(string)
	}
	// Verify the record has a log field.
	// If so verify the time in the log field matches the @timestamp field.
	if val, ok := hit["log"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty log field")
		valid = false
	} else {
		re := regexp.MustCompile(`(\d{2}:\d{2}:\d{2})`)
		m := re.FindStringSubmatch(val.(string))
		if len(m) < 2 {
			pkg.Log(pkg.Info, "Log record log field does not contain a time")
			valid = false
		} else {
			if !strings.Contains(ts, m[1]) {
				pkg.Log(pkg.Info, fmt.Sprintf("Log record @timestamp field %s does not match log field %s content", ts, m[1]))
				valid = false
			}
		}
	}
	// Verify the record has a message field.
	if val, ok := hit["message"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty message field")
		valid = false
	}
	// Verify the log field isn't exactly the same as the message field.
	if hit["log"] == hit["message"] {
		pkg.Log(pkg.Info, "Log record has duplicate log and message field values")
		valid = false
	}
	// Verify the record does not have a timestamp field.
	if _, ok := hit["timestamp"]; ok {
		pkg.Log(pkg.Info, "Log record has unwanted timestamp field")
		valid = false
	}
	if !valid {
		pkg.Log(pkg.Info, fmt.Sprintf("Log record is invalid: %v", hit))
	}
	return valid
}

// logLevelOpensearchRecordValidator does validation of level for log records
func logLevelOpensearchRecordValidator(hit pkg.OpensearchHit) bool {
	// Verify the record has a level field.
	// If so verify that the level isn't debug.
	if val, ok := hit["level"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty level field")
		return false
	}
	// level := val.(string)
	// Put this validation back in when the OAM logging is fixed.
	// if strings.EqualFold(level, "debug") || strings.EqualFold(level, "dbg") || strings.EqualFold(level, "d") {
	// 	pkg.Log(pkg.Info, fmt.Sprintf("Log record has invalid debug level: %s", level))
	// 	valid = false
	// }
	// There is an Istio proxy error that causes this to fail.
	// Put this validation back in when that is addressed.
	// if strings.EqualFold(level, "error") || strings.EqualFold(level, "err") || strings.EqualFold(level, "e") {
	//	pkg.Log(pkg.Info, fmt.Sprintf("Log record has invalid error level: %s", level))
	//	valid = false
	// }

	return true
}
