// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/verrazzano/verrazzano/pkg/k8sutil"
	"github.com/verrazzano/verrazzano/platform-operator/constants"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	jaegerSpanIndexPrefix  = "verrazzano-jaeger-span"
	jaegerClusterNameLabel = "verrazzano_cluster"
	jaegerClusterName      = "local"
)

type JaegerTraceData struct {
	TraceID string `json:"traceID"`
	Spans   []struct {
		TraceID       string `json:"traceID"`
		SpanID        string `json:"spanID"`
		Flags         int    `json:"flags"`
		OperationName string `json:"operationName"`
		References    []struct {
			RefType string `json:"refType"`
			TraceID string `json:"traceID"`
			SpanID  string `json:"spanID"`
		} `json:"references"`
		StartTime int64 `json:"startTime"`
		Duration  int   `json:"duration"`
		Tags      []struct {
			Key   string      `json:"key"`
			Type  string      `json:"type"`
			Value interface{} `json:"value"`
		} `json:"tags"`
		Logs []struct {
			Timestamp int64 `json:"timestamp"`
			Fields    []struct {
				Key   string `json:"key"`
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"fields"`
		} `json:"logs"`
		ProcessID string      `json:"processID"`
		Warnings  interface{} `json:"warnings"`
	} `json:"spans"`
	Processes struct {
		P1 struct {
			ServiceName string `json:"serviceName"`
			Tags        []struct {
				Key   string `json:"key"`
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"tags"`
		} `json:"p1"`
	} `json:"processes"`
	Warnings interface{} `json:"warnings"`
}

type JaegerTraceDataWrapper struct {
	Data   []JaegerTraceData `json:"data"`
	Total  int               `json:"total"`
	Limit  int               `json:"limit"`
	Offset int               `json:"offset"`
	Errors interface{}       `json:"errors"`
}

//IsJaegerInstanceCreated checks whether the default Jaeger CR is created
func IsJaegerInstanceCreated() (bool, error) {
	deployments, err := ListDeployments(constants.VerrazzanoMonitoringNamespace)
	if err != nil {
		return false, err
	}
	if len(deployments.Items) > 0 {
		return true, nil
	}
	return false, nil
}

//JaegerSpanRecordFoundInOpenSearch checks if jaeger span records are found in OpenSearch storage
func JaegerSpanRecordFoundInOpenSearch(kubeconfigPath string, after time.Time, serviceName string) (bool, error) {
	indexName, err := GetJaegerSpanIndexName(kubeconfigPath)
	if err != nil {
		return false, err
	}
	fields := map[string]string{
		"process.serviceName": serviceName,
	}
	searchResult := querySystemElasticSearch(indexName, fields, kubeconfigPath)
	if len(searchResult) == 0 {
		Log(Info, fmt.Sprintf("Expected to find log record matching fields %v", fields))
		return false, nil
	}
	found := findJaegerSpanHits(searchResult, &after)
	if !found {
		Log(Error, fmt.Sprintf("Failed to find recent jaeger span record for service %s", serviceName))
	}
	return found, nil
}

//GetJaegerSpanIndexName returns the index name used in OpenSearch used for storage
func GetJaegerSpanIndexName(kubeconfigPath string) (string, error) {
	var jaegerIndices []string
	for _, indexName := range listSystemElasticSearchIndices(kubeconfigPath) {
		if strings.HasPrefix(indexName, jaegerSpanIndexPrefix) {
			jaegerIndices = append(jaegerIndices, indexName)
			break
		}
	}
	if len(jaegerIndices) > 0 {
		return jaegerIndices[0], nil
	}
	return "", fmt.Errorf("Jaeger Span index not found")
}

// IsJaegerMetricFound validates if the given jaeger metrics contain the labels with values specified as key-value pairs of the map
func IsJaegerMetricFound(kubeconfigPath, metricName string, kv map[string]string) bool {
	compMetrics, err := QueryMetricWithLabel(metricName, kubeconfigPath, jaegerClusterNameLabel, jaegerClusterName)
	if err != nil {
		return false
	}
	metrics := JTq(compMetrics, "data", "result").([]interface{})
	for _, metric := range metrics {
		metricFound := true
		for key, value := range kv {
			if Jq(metric, "metric", key) != value {
				metricFound = false
				break
			}
		}
		return metricFound
	}
	return false
}

//ListJaegerTraces lists all trace ids for a given service.
func ListJaegerTraces(kubeconfigPath string, serviceName string) []string {
	var traces []string
	url := fmt.Sprintf("%s/api/traces?service=%s", getJaegerURL(kubeconfigPath), serviceName)
	username, password, err := getJaegerUsernamePassword(kubeconfigPath)
	if err != nil {
		return traces
	}
	resp, err := getJaegerWithBasicAuth(url, "", username, password, kubeconfigPath)
	if err != nil {
		Log(Error, fmt.Sprintf("Error getting Jaeger traces: url=%s, error=%v", url, err))
		return traces
	}
	if resp.StatusCode != http.StatusOK {
		Log(Error, fmt.Sprintf("Error retrieving Jaeger traces: url=%s, status=%d", url, resp.StatusCode))
		return traces
	}
	Log(Debug, fmt.Sprintf("traces: %s", resp.Body))
	var jaegerTraceDataWrapper JaegerTraceDataWrapper
	json.Unmarshal(resp.Body, &jaegerTraceDataWrapper)
	for _, traceObj := range jaegerTraceDataWrapper.Data {
		traces = append(traces, traceObj.TraceID)
	}
	return traces
}

//ListServicesInJaeger lists the services whose traces are available in Jaeger
func ListServicesInJaeger(kubeconfigPath string) []string {
	var services []string
	url := fmt.Sprintf("%s/api/services", getJaegerURL(kubeconfigPath))
	username, password, err := getJaegerUsernamePassword(kubeconfigPath)
	if err != nil {
		return services
	}
	resp, err := getJaegerWithBasicAuth(url, "", username, password, kubeconfigPath)
	if err != nil {
		Log(Error, fmt.Sprintf("Error getting Jaeger traces: url=%s, error=%v", url, err))
		return services
	}
	if resp.StatusCode != http.StatusOK {
		Log(Error, fmt.Sprintf("Error retrieving Jaeger traces: url=%s, status=%d", url, resp.StatusCode))
		return services
	}
	Log(Debug, fmt.Sprintf("traces: %s", resp.Body))
	var serviceMap map[string][]string
	json.Unmarshal(resp.Body, &serviceMap)
	services = append(services, serviceMap["data"]...)
	return services
}

// DoesCronJobExist returns whether a cronjob with the given name and namespace exists for the cluster
func DoesCronJobExist(namespace string, name string) (bool, error) {
	cronjobs, err := ListCronJobNamesMatchingLabels(namespace, nil)
	if err != nil {
		Log(Error, fmt.Sprintf("Failed listing deployments in cluster for namespace %s: %v", namespace, err))
		return false, err
	}
	for _, cronJobName := range cronjobs {
		if strings.HasPrefix(cronJobName, name) {
			return true, nil
		}
	}
	return false, nil
}

// ListDeploymentsMatchingLabels returns the list of deployments in a given namespace matching the given labels for the cluster
func ListDeploymentsMatchingLabels(namespace string, matchLabels map[string]string) (*appsv1.DeploymentList, error) {
	// Get the Kubernetes clientset
	clientset, err := k8sutil.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}
	listOptions := metav1.ListOptions{}
	if matchLabels != nil {
		var selector labels.Selector
		for k, v := range matchLabels {
			selectorLabel, _ := labels.NewRequirement(k, selection.Equals, []string{v})
			selector = labels.NewSelector()
			selector = selector.Add(*selectorLabel)
		}
		listOptions.LabelSelector = selector.String()
	}
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), listOptions)
	if err != nil {
		Log(Error, fmt.Sprintf("Failed to list deployments in namespace %s: %v", namespace, err))
		return nil, err
	}
	return deployments, nil
}

// ListCronJobNamesMatchingLabels returns the list of cronjobs in a given namespace matching the given labels for the cluster
func ListCronJobNamesMatchingLabels(namespace string, matchLabels map[string]string) ([]string, error) {
	var cronjobNames []string
	// Get the Kubernetes clientset
	clientset, err := k8sutil.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}
	info, err := clientset.ServerVersion()
	if err != nil {
		return nil, err
	}
	majorVersion, err := strconv.Atoi(info.Major)
	if err != nil {
		return nil, err
	}
	if majorVersion > 1 {
		return nil, fmt.Errorf("Unknown major version %d", majorVersion)
	}
	minorVersion, err := strconv.Atoi(info.Minor)
	if err != nil {
		return nil, err
	}
	// For k8s version 1.20 and lesser, cronjobs are created under version batch/v1beta1
	// For k8s version greater than 1.20, cronjobs are created under version batch/v1
	if minorVersion <= 20 {
		cronJobs, err := listV1Beta1CronJobNames(clientset, namespace, fillLabelSelectors(matchLabels))
		if err != nil {
			return nil, err
		}
		for _, cronjob := range cronJobs {
			cronjobNames = append(cronjobNames, cronjob.Name)
		}
	} else {
		cronJobs, err := listV1CronJobNames(clientset, namespace, fillLabelSelectors(matchLabels))
		if err != nil {
			return nil, err
		}
		for _, cronjob := range cronJobs {
			cronjobNames = append(cronjobNames, cronjob.Name)
		}
	}
	return cronjobNames, nil
}

// fillLabelSelectors fills the match labels from map to be passed in list options
func fillLabelSelectors(matchLabels map[string]string) metav1.ListOptions {
	listOptions := metav1.ListOptions{}
	if matchLabels != nil {
		var selector labels.Selector
		for k, v := range matchLabels {
			selectorLabel, _ := labels.NewRequirement(k, selection.Equals, []string{v})
			selector = labels.NewSelector()
			selector = selector.Add(*selectorLabel)
		}
		listOptions.LabelSelector = selector.String()
	}
	return listOptions
}

// listV1CronJobNames lists the cronjob under batch/v1 api version for k8s version > 1.20
func listV1CronJobNames(clientset *kubernetes.Clientset, namespace string, listOptions metav1.ListOptions) ([]v1.CronJob, error) {
	var cronJobs []v1.CronJob
	cronJobList, err := clientset.BatchV1().CronJobs(namespace).List(context.TODO(), listOptions)
	if err != nil {
		Log(Error, fmt.Sprintf("Failed to list v1/cronjobs in namespace %s: %v", namespace, err))
		return cronJobs, err
	}
	return cronJobList.Items, nil
}

// listV1Beta1CronJobNames lists the cronjob under batch/v1beta1 api version for k8s version <= 1.20
func listV1Beta1CronJobNames(clientset *kubernetes.Clientset, namespace string, listOptions metav1.ListOptions) ([]v1beta1.CronJob, error) {
	var cronJobs []v1beta1.CronJob
	cronJobList, err := clientset.BatchV1beta1().CronJobs(namespace).List(context.TODO(), listOptions)
	if err != nil {
		Log(Error, fmt.Sprintf("Failed to list v1beta1/cronjobs in namespace %s: %v", namespace, err))
		return cronJobs, err
	}
	return cronJobList.Items, nil
}

// getJaegerWithBasicAuth access Jaeger with GET using basic auth, using a given kubeconfig
func getJaegerWithBasicAuth(url string, hostHeader string, username string, password string, kubeconfigPath string) (*HTTPResponse, error) {
	retryableClient, err := getJaegerClient(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return doReq(url, "GET", "", hostHeader, username, password, nil, retryableClient)
}

// getJaegerClient returns the Jaeger client which can be used for GET/POST operations using a given kubeconfig
func getJaegerClient(kubeconfigPath string) (*retryablehttp.Client, error) {
	client, err := GetVerrazzanoHTTPClient(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return client, err
}

// getJaegerURL returns Jaeger URL from the corresponding ingress resource using the given kubeconfig
func getJaegerURL(kubeconfigPath string) string {
	clientset, err := GetKubernetesClientsetForCluster(kubeconfigPath)
	if err != nil {
		Log(Error, fmt.Sprintf("Failed to get clientset for cluster %v", err))
		return ""
	}
	ingressList, _ := clientset.NetworkingV1().Ingresses(VerrazzanoNamespace).List(context.TODO(), metav1.ListOptions{})
	for _, ingress := range ingressList.Items {
		if ingress.Name == "verrazzano-jaeger" {
			Log(Info, fmt.Sprintf("Found Jaeger Ingress %v, host %s", ingress.Name, ingress.Spec.Rules[0].Host))
			return fmt.Sprintf("https://%s", ingress.Spec.Rules[0].Host)
		}
	}
	return ""
}

// getJaegerUsernamePassword returns the username/password for connecting to Jaeger
func getJaegerUsernamePassword(kubeconfigPath string) (username, password string, err error) {
	password, err = GetVerrazzanoPasswordInCluster(kubeconfigPath)
	if err != nil {
		return "", "", err
	}
	return "verrazzano", password, err
}

// findJaegerSpanHits returns the number of span hits that are older than the given time
func findJaegerSpanHits(searchResult map[string]interface{}, after *time.Time) bool {
	hits := Jq(searchResult, "hits", "hits")
	if hits == nil {
		Log(Info, "Expected to find hits in span record query results")
		return false
	}
	Log(Info, fmt.Sprintf("Found %d records", len(hits.([]interface{}))))
	if len(hits.([]interface{})) == 0 {
		Log(Info, "Expected span record query results to contain at least one hit")
		return false
	}
	if after == nil {
		return true
	}
	for _, hit := range hits.([]interface{}) {
		timestamp := Jq(hit, "_source", "startTimeMillis")
		t := time.UnixMilli(int64(timestamp.(float64)))
		if t.After(*after) {
			Log(Info, fmt.Sprintf("Found recent record: %f", timestamp))
			return true
		}
		Log(Info, fmt.Sprintf("Found old record: %f", timestamp))
	}
	return true
}