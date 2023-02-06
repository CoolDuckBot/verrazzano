// Copyright (c) 2022, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package operator

import (
	"fmt"
	"path/filepath"

	"github.com/verrazzano/verrazzano/pkg/k8s/ready"
	"github.com/verrazzano/verrazzano/pkg/vzcr"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/networkpolicies"
	clipkg "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/verrazzano/verrazzano/pkg/k8sutil"
	vzapi "github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1alpha1"
	installv1beta1 "github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1beta1"
	"github.com/verrazzano/verrazzano/platform-operator/constants"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/authproxy"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/certmanager"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/common"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/helm"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/nginx"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/spi"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/vmo"
	"github.com/verrazzano/verrazzano/platform-operator/internal/config"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// ComponentName is the name of the component
const ComponentName = "prometheus-operator"

// ComponentNamespace is the namespace of the component
const ComponentNamespace = constants.VerrazzanoMonitoringNamespace

// ComponentJSONName is the JSON name of the component in the CRD
const ComponentJSONName = "prometheusOperator"

const chartDir = "prometheus-community/kube-prometheus-stack"

const (
	prometheusHostName        = "prometheus.vmi.system"
	prometheusCertificateName = "system-tls-prometheus"

	istioPrometheus = "prometheus-server"
)

type prometheusComponent struct {
	helm.HelmComponent
}

func NewComponent() spi.Component {
	return prometheusComponent{
		helm.HelmComponent{
			ReleaseName:               ComponentName,
			JSONName:                  ComponentJSONName,
			ChartDir:                  filepath.Join(config.GetThirdPartyDir(), chartDir),
			ChartNamespace:            ComponentNamespace,
			IgnoreNamespaceOverride:   true,
			SupportsOperatorInstall:   true,
			SupportsOperatorUninstall: true,
			MinVerrazzanoVersion:      constants.VerrazzanoVersion1_3_0,
			ImagePullSecretKeyname:    "global.imagePullSecrets[0].name",
			ValuesFile:                filepath.Join(config.GetHelmOverridesDir(), "prometheus-operator-values.yaml"),
			// the dependency on the VMO is to ensure that a persistent volume is retained and the claim is released
			// so that persistent storage can be migrated to the new Prometheus
			Dependencies:            []string{networkpolicies.ComponentName, nginx.ComponentName, certmanager.ComponentName, vmo.ComponentName},
			AppendOverridesFunc:     AppendOverrides,
			GetInstallOverridesFunc: GetOverrides,
		},
	}
}

// IsEnabled returns true if the Prometheus Operator is enabled or if the component is not specified
// in the Verrazzano CR.
func (c prometheusComponent) IsEnabled(effectiveCR runtime.Object) bool {
	return vzcr.IsPrometheusOperatorEnabled(effectiveCR)
}

// IsReady checks if the Prometheus Operator deployment is ready
func (c prometheusComponent) IsReady(ctx spi.ComponentContext) bool {
	if c.HelmComponent.IsReady(ctx) {
		return isPrometheusOperatorReady(ctx)
	}
	return false
}

func (c prometheusComponent) IsAvailable(ctx spi.ComponentContext) (reason string, available vzapi.ComponentAvailability) {
	listOptions, err := prometheusOperatorListOptions()
	if err != nil {
		return err.Error(), vzapi.ComponentUnavailable
	}
	return (&ready.AvailabilityObjects{DeploymentSelectors: []clipkg.ListOption{listOptions}}).IsAvailable(ctx.Log(), ctx.Client())
}

// MonitorOverrides checks whether monitoring is enabled for install overrides sources
func (c prometheusComponent) MonitorOverrides(ctx spi.ComponentContext) bool {
	if ctx.EffectiveCR().Spec.Components.PrometheusOperator == nil {
		return false
	}
	if ctx.EffectiveCR().Spec.Components.PrometheusOperator.MonitorChanges != nil {
		return *ctx.EffectiveCR().Spec.Components.PrometheusOperator.MonitorChanges
	}
	return true
}

// PreInstall updates resources necessary for the Prometheus Operator Component installation
func (c prometheusComponent) PreInstall(ctx spi.ComponentContext) error {
	if err := preInstallUpgrade(ctx); err != nil {
		return err
	}
	return c.HelmComponent.PreInstall(ctx)
}

// PreUpgrade updates resources necessary for the Prometheus Operator Component installation
func (c prometheusComponent) PreUpgrade(ctx spi.ComponentContext) error {
	if err := preInstallUpgrade(ctx); err != nil {
		return err
	}
	return c.HelmComponent.PreUpgrade(ctx)
}

// PostInstall creates/updates associated resources after this component is installed
func (c prometheusComponent) PostInstall(ctx spi.ComponentContext) error {
	if err := postInstallUpgrade(ctx); err != nil {
		return err
	}

	// these need to be set for helm component post install processing
	c.IngressNames = c.GetIngressNames(ctx)
	c.Certificates = c.GetCertificateNames(ctx)

	return c.HelmComponent.PostInstall(ctx)
}

// PostInstall creates/updates associated resources after this component is upgraded
func (c prometheusComponent) PostUpgrade(ctx spi.ComponentContext) error {
	if err := postInstallUpgrade(ctx); err != nil {
		return err
	}

	return c.HelmComponent.PostUpgrade(ctx)
}

// ValidateInstall verifies the installation of the Verrazzano object
func (c prometheusComponent) ValidateInstall(vz *vzapi.Verrazzano) error {
	convertedVZ := installv1beta1.Verrazzano{}
	if err := common.ConvertVerrazzanoCR(vz, &convertedVZ); err != nil {
		return err
	}
	if err := checkExistingCNEPrometheus(vz); err != nil {
		return err
	}
	return c.validatePrometheusOperator(&convertedVZ)
}

// ValidateUpdate verifies the upgrade of the Verrazzano object
func (c prometheusComponent) ValidateUpdate(old *vzapi.Verrazzano, new *vzapi.Verrazzano) error {
	if c.IsEnabled(old) && !c.IsEnabled(new) {
		if vzcr.IsPrometheusEnabled(new) {
			return fmt.Errorf("disabling component %s is not allowed while %s is enabled", ComponentJSONName, prometheusName)
		}
	}
	convertedVZ := installv1beta1.Verrazzano{}
	if err := common.ConvertVerrazzanoCR(new, &convertedVZ); err != nil {
		return err
	}
	return c.validatePrometheusOperator(&convertedVZ)
}

// ValidateInstallV1Beta1 verifies the installation of the Verrazzano object
func (c prometheusComponent) ValidateInstallV1Beta1(vz *installv1beta1.Verrazzano) error {
	if err := checkExistingCNEPrometheus(vz); err != nil {
		return err
	}
	return c.validatePrometheusOperator(vz)
}

// ValidateUpdateV1Beta1 verifies the upgrade of the Verrazzano object
func (c prometheusComponent) ValidateUpdateV1Beta1(old *installv1beta1.Verrazzano, new *installv1beta1.Verrazzano) error {
	if c.IsEnabled(old) && !c.IsEnabled(new) {
		if vzcr.IsPrometheusEnabled(new) {
			return fmt.Errorf("disabling component %s is not allowed while %s is enabled", ComponentJSONName, prometheusName)
		}
	}
	return c.validatePrometheusOperator(new)
}

// GetIngressNames - gets the names of the ingresses associated with this component
func (c prometheusComponent) GetIngressNames(ctx spi.ComponentContext) []types.NamespacedName {
	var ingressNames []types.NamespacedName

	if vzcr.IsPrometheusEnabled(ctx.EffectiveCR()) {
		ns := ComponentNamespace
		if vzcr.IsAuthProxyEnabled(ctx.EffectiveCR()) {
			ns = authproxy.ComponentNamespace
		}
		ingressNames = append(ingressNames, types.NamespacedName{
			Namespace: ns,
			Name:      constants.PrometheusIngress,
		})
	}

	return ingressNames
}

// getCertificateNames - gets the names of the TLS ingress certificates associated with this component
func (c prometheusComponent) GetCertificateNames(ctx spi.ComponentContext) []types.NamespacedName {
	var certificateNames []types.NamespacedName

	if vzcr.IsPrometheusEnabled(ctx.EffectiveCR()) {
		ns := ComponentNamespace
		if vzcr.IsAuthProxyEnabled(ctx.EffectiveCR()) {
			ns = authproxy.ComponentNamespace
		}
		certificateNames = append(certificateNames, types.NamespacedName{
			Namespace: ns,
			Name:      prometheusCertificateName,
		})
	}

	return certificateNames
}

// checkExistingCNEPrometheus checks if Prometheus is already installed
// OLCNE Istio module may have Prometheus installed in istio-system namespace
func checkExistingCNEPrometheus(vz runtime.Object) error {
	if !vzcr.IsPrometheusEnabled(vz) {
		return nil
	}
	if err := k8sutil.ErrorIfDeploymentExists(constants.IstioSystemNamespace, istioPrometheus); err != nil {
		return err
	}
	if err := k8sutil.ErrorIfServiceExists(constants.IstioSystemNamespace, istioPrometheus); err != nil {
		return err
	}
	return nil
}
