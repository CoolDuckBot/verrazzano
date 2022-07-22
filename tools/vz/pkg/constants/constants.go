// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package constants

// GlobalFlagKubeConfig - global flag for specifying the location of the kube config
const GlobalFlagKubeConfig = "kubeconfig"
const GlobalFlagKubeConfigHelp = "Path to the kubeconfig file to use"

// GlobalFlagContext - global flag for specifying which kube config context to use
const GlobalFlagContext = "context"
const GlobalFlagContextHelp = "The name of the kubeconfig context to use"

// GlobalFlagHelp - global help flag
const GlobalFlagHelp = "help"

// Flags that are common to more than one command
const (
	WaitFlag        = "wait"
	WaitFlagHelp    = "Wait for the command to complete and stream the logs to the console. The wait period is controlled by --timeout."
	WaitFlagDefault = true

	TimeoutFlag     = "timeout"
	TimeoutFlagHelp = "Limits the amount of time a command will wait to complete"

	VersionFlag            = "version"
	VersionFlagDefault     = "latest"
	VersionFlagInstallHelp = "The version of Verrazzano to install"
	VersionFlagUpgradeHelp = "The version of Verrazzano to upgrade to"

	DryRunFlag = "dry-run"

	SetFlag          = "set"
	SetFlagShorthand = "s"
	SetFlagHelp      = "Override a Verrazzano resource value (e.g. --set profile=dev).  This flag can be specified multiple times."

	OperatorFileFlag     = "operator-file"
	OperatorFileFlagHelp = "The path to the file for installing the Verrazzano platform operator. The default is derived from the version string."

	LogFormatFlag = "log-format"
	LogFormatHelp = "The format of the log output. Valid output formats are \"simple\" and \"json\"."

	FilenameFlag          = "filename"
	FilenameFlagShorthand = "f"
	FilenameFlagHelp      = "Path to file containing Verrazzano custom resource.  This flag can be specified multiple times to overlay multiple files.  Specifying \"-\" as the filename accepts input from stdin."
)

// VerrazzanoReleaseList - API for getting the list of Verrazzano releases
const VerrazzanoReleaseList = "https://api.github.com/repos/verrazzano/verrazzano/releases"

// VerrazzanoOperatorURL - URL for downloading Verrazzano releases
const VerrazzanoOperatorURL = "https://github.com/verrazzano/verrazzano/releases/download/%s/operator.yaml"

const VerrazzanoPlatformOperator = "verrazzano-platform-operator"

const VerrazzanoApplicationOperator = "verrazzano-application-operator"

const VerrazzanoMonitoringOperator = "verrazzano-monitoring-operator"

const VerrazzanoUninstall = "verrazzano-uninstall"

const VerrazzanoInstall = "verrazzano-install"

const VerrazzanoManagedCluster = "verrazzano-managed-cluster"

const VerrazzanoPlatformOperatorWait = 1

const OAMAppConfigurations = "applicationconfigurations"

const OAMMCAppConfigurations = "multiclusterapplicationconfigurations"

const OAMMCCompConfigurations = "multiclustercomponents"

const OAMComponents = "components"

const OAMMetricsTraits = "metricstraits"

const OAMIngressTraits = "ingresstraits"

const OAMProjects = "verrazzanoprojects"

const OAMManagedClusters = "verrazzanomanagedclusters"

// Analysis tool flags
const (
	DirectoryFlagName  = "capture-dir"
	DirectoryFlagValue = ""
	DirectoryFlagUsage = "Directory holding the captured data [Required]"

	ReportFileFlagName  = "report-file"
	ReportFileFlagValue = ""
	ReportFileFlagUsage = "Name of report output file. (default stdout)"

	ReportFormatFlagName  = "report-format"
	ReportFormatFlagValue = "simple"
	ReportFormatFlagUsage = "The format of the report output. Valid output format is \"simple\""
)

// Constants for bug report
const (
	BugReportFileFlagName     = "report-file"
	BugReportFileFlagValue    = ""
	BugReportFileFlagShort    = "r"
	BugReportFileFlagUsage    = "The report file to be created by bug-report command, as a .tar.gz file. Defaults to bug-report.tar.gz in the current directory."
	BugReportFileDefaultValue = "bug-report.tar.gz"

	BugReportIncludeNSFlagName  = "include-namespaces"
	BugReportIncludeNSFlagShort = "i"
	BugReportIncludeNSFlagUsage = "A comma separated list of additional namespaces to collect information from the cluster. This flag can be specified multiple times like --include-namespaces ns1 --include-namespaces ns..."

	BugReportDir = "bug-report"

	// File name for the log captured from the pod
	LogFile = "logs.txt"

	// File names for the various resources
	VzResource       = "verrazzano-resources.json"
	DeploymentsJSON  = "deployments.json"
	EventsJSON       = "events.json"
	PodsJSON         = "pods.json"
	ServicesJSON     = "services.json"
	ReplicaSetsJSON  = "replicasets.json"
	DaemonSetsJSON   = "daemonsets.json"
	IngressJSON      = "ingress.json"
	StatefulSetsJSON = "statefulsets.json"
	AppConfigJSON    = "application-configurations.json"
	ComponentJSON    = "components.json"
	IngressTraitJSON = "ingress-traits.json"
	MetricsTraitJSON = "metrics-traits.json"
	McAppConfigJSON  = "multicluster-application-configurations.json"
	McComponentJSON  = "multicluster-components.json"
	VzProjectsJSON   = "verrazzano-projects.json"
	VmcJSON          = "verrazzano-managed-clusters.json"

	// Indentation when the resource is marshalled as Json
	JSONIndent = "  "

	// The prefix used for the json.MarshalIndent
	JSONPrefix = ""

	// Top level directory for the bug report, keeping cluster-dump for now to support the analyze the command
	BugReportRoot = "cluster-dump"

	// Label for application
	AppLabel    = "app"
	K8SAppLabel = "k8s-app"
)