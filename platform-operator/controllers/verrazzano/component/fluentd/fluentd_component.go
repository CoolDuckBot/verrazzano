// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package fluentd

import (
	"context"
	"fmt"
	"github.com/verrazzano/verrazzano/pkg/k8s/ready"
	"github.com/verrazzano/verrazzano/pkg/vzcr"
	"github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1beta1"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/networkpolicies"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	"path/filepath"

	"github.com/verrazzano/verrazzano/pkg/k8s/resource"
	"github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1alpha1"
	vzconst "github.com/verrazzano/verrazzano/platform-operator/constants"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/helm"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/spi"
	"github.com/verrazzano/verrazzano/platform-operator/internal/config"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// ComponentName is the name of the component
	ComponentName = "fluentd"

	// ComponentNamespace is the namespace of the component
	ComponentNamespace = vzconst.VerrazzanoSystemNamespace

	// ComponentJSONName is the JSON name of the verrazzano component in CRD
	ComponentJSONName = "fluentd"

	// HelmChartReleaseName is the helm chart release name
	HelmChartReleaseName = "fluentd"

	// HelmChartDir is the name of the helm chart directory
	HelmChartDir = "verrazzano-fluentd"

	// vzImagePullSecretKeyName is the Helm key name for the VZ chart image pull secret
	vzImagePullSecretKeyName = "global.imagePullSecrets[0]"
)

var (
	// For Unit test purposes
	writeFileFunc = os.WriteFile
)

// fluentdComponent represents an Fluentd component
type fluentdComponent struct {
	helm.HelmComponent
}

// Verify that fluentdComponent implements Component
var _ spi.Component = fluentdComponent{}

// NewComponent returns a new Fluentd component
func NewComponent() spi.Component {
	return fluentdComponent{
		helm.HelmComponent{
			ReleaseName:               HelmChartReleaseName,
			JSONName:                  ComponentJSONName,
			ChartDir:                  filepath.Join(config.GetHelmChartsDir(), HelmChartDir),
			ChartNamespace:            ComponentNamespace,
			IgnoreNamespaceOverride:   true,
			SupportsOperatorInstall:   true,
			SupportsOperatorUninstall: true,
			ImagePullSecretKeyname:    vzImagePullSecretKeyName,
			AppendOverridesFunc:       appendOverrides,
			Dependencies:              []string{networkpolicies.ComponentName},
			GetInstallOverridesFunc:   GetOverrides,
			AvailabilityObjects: &ready.AvailabilityObjects{
				DaemonsetNames: []types.NamespacedName{
					{
						Name:      ComponentName,
						Namespace: ComponentNamespace,
					},
				},
			},
		},
	}
}

// ValidateInstall checks if the specified Verrazzano CR is valid for this component to be installed
func (c fluentdComponent) ValidateInstall(vz *v1alpha1.Verrazzano) error {
	vzV1Beta1 := &v1beta1.Verrazzano{}

	if err := vz.ConvertTo(vzV1Beta1); err != nil {
		return err
	}

	return c.ValidateInstallV1Beta1(vzV1Beta1)
}

// ValidateUpdate checks if the specified new Verrazzano CR is valid for this component to be updated
func (c fluentdComponent) ValidateUpdate(old *v1alpha1.Verrazzano, new *v1alpha1.Verrazzano) error {
	oldBeta := &v1beta1.Verrazzano{}
	newBeta := &v1beta1.Verrazzano{}

	if err := old.ConvertTo(oldBeta); err != nil {
		return err
	}

	if err := new.ConvertTo(newBeta); err != nil {
		return err
	}

	return c.ValidateUpdateV1Beta1(oldBeta, newBeta)
}

// ValidateInstall checks if the specified Verrazzano CR is valid for this component to be installed
func (c fluentdComponent) ValidateInstallV1Beta1(vz *v1beta1.Verrazzano) error {
	if err := validateFluentd(vz); err != nil {
		return err
	}
	return c.HelmComponent.ValidateInstallV1Beta1(vz)
}

// ValidateUpdate checks if the specified new Verrazzano CR is valid for this component to be updated
func (c fluentdComponent) ValidateUpdateV1Beta1(old *v1beta1.Verrazzano, new *v1beta1.Verrazzano) error {
	// Do not allow disabling active components
	if err := c.checkEnabled(old, new); err != nil {
		return err
	}
	if err := validateFluentd(new); err != nil {
		return err
	}
	return c.HelmComponent.ValidateUpdateV1Beta1(old, new)
}

func (c fluentdComponent) checkEnabled(old *v1beta1.Verrazzano, new *v1beta1.Verrazzano) error {
	// Do not allow disabling of any component post-install for now
	if c.IsEnabled(old) && !c.IsEnabled(new) {
		return fmt.Errorf("disabling component %s is not allowed", ComponentJSONName)
	}
	return nil
}

// PostInstall - post-install, clean up temp files
func (c fluentdComponent) PostInstall(ctx spi.ComponentContext) error {
	cleanTempFiles(ctx)
	return c.HelmComponent.PostInstall(ctx)
}

// PostUpgrade Fluentd component post-upgrade processing
func (c fluentdComponent) PostUpgrade(ctx spi.ComponentContext) error {
	ctx.Log().Debugf("Fluentd component post-upgrade")
	cleanTempFiles(ctx)
	return c.HelmComponent.PostUpgrade(ctx)
}

// PreInstall Fluentd component pre-install processing; create and label required namespaces, copy any
// required secrets
func (c fluentdComponent) PreInstall(ctx spi.ComponentContext) error {
	if err := loggingPreInstall(ctx); err != nil {
		return ctx.Log().ErrorfNewErr("Failed copying logging secrets for Verrazzano: %v", err)
	}
	if err := checkSecretExists(ctx); err != nil {
		return err
	}
	return c.HelmComponent.PreInstall(ctx)
}

func (c fluentdComponent) Reconcile(ctx spi.ComponentContext) error {
	installed, err := c.IsInstalled(ctx)
	if err != nil {
		return err
	}
	if installed {
		err = c.Install(ctx)
	}
	return err
}

// Install Fluentd component install processing
func (c fluentdComponent) Install(ctx spi.ComponentContext) error {
	if err := c.HelmComponent.Install(ctx); err != nil {
		return err
	}
	return nil
}

// PreUpgrade Fluentd component pre-upgrade processing
func (c fluentdComponent) PreUpgrade(ctx spi.ComponentContext) error {
	if err := fluentdPreUpgrade(ctx, ComponentNamespace); err != nil {
		return err
	}
	if err := checkSecretExists(ctx); err != nil {
		return err
	}
	return c.HelmComponent.PreUpgrade(ctx)
}

// Uninstall Fluentd to handle upgrade case where Fluentd was not its own helm chart.
// In that case, we need to delete the Fluentd resources explicitly
func (c fluentdComponent) Uninstall(context spi.ComponentContext) error {
	installed, err := c.HelmComponent.IsInstalled(context)
	if err != nil {
		return err
	}

	// If the helm chart is installed, then uninstall
	if installed {
		return c.HelmComponent.Uninstall(context)
	}

	// Attempt to delete the Fluentd resources
	rs := getFluentdManagedResources()
	for _, r := range rs {
		err := resource.Resource{
			Name:      r.NamespacedName.Name,
			Namespace: r.NamespacedName.Namespace,
			Client:    context.Client(),
			Object:    r.Obj,
			Log:       context.Log(),
		}.Delete()
		if err != nil {
			return err
		}
	}

	return nil
}

// Upgrade Fluentd component upgrade processing
func (c fluentdComponent) Upgrade(ctx spi.ComponentContext) error {
	return c.HelmComponent.Install(ctx)
}

// IsReady component check
func (c fluentdComponent) IsReady(ctx spi.ComponentContext) bool {
	if c.HelmComponent.IsReady(ctx) {
		return c.isFluentdReady(ctx)
	}
	return false
}

// IsInstalled component check
func (c fluentdComponent) IsInstalled(ctx spi.ComponentContext) (bool, error) {
	daemonSet := &appsv1.DaemonSet{}
	err := ctx.Client().Get(context.TODO(), types.NamespacedName{Namespace: ComponentNamespace, Name: ComponentName}, daemonSet)
	if errors.IsNotFound(err) {
		return false, nil
	}
	if err != nil {
		ctx.Log().Errorf("Failed to get %s/%s daemonSet: %v", ComponentNamespace, ComponentName, err)
		return false, err
	}
	return true, nil
}

// IsEnabled fluentd-specific enabled check for installation
func (c fluentdComponent) IsEnabled(effectiveCR runtime.Object) bool {
	return vzcr.IsFluentdEnabled(effectiveCR)
}

// MonitorOverrides checks whether monitoring of install overrides is enabled or not
func (c fluentdComponent) MonitorOverrides(ctx spi.ComponentContext) bool {
	if ctx.EffectiveCR().Spec.Components.Fluentd != nil {
		if ctx.EffectiveCR().Spec.Components.Fluentd.MonitorChanges != nil {
			return *ctx.EffectiveCR().Spec.Components.Fluentd.MonitorChanges
		}
		return true
	}
	return false
}

// GetOverrides returns install overrides for a component
func GetOverrides(object runtime.Object) interface{} {
	if effectiveCR, ok := object.(*v1alpha1.Verrazzano); ok {
		if effectiveCR.Spec.Components.Fluentd != nil {
			return effectiveCR.Spec.Components.Fluentd.ValueOverrides
		}
		return []v1alpha1.Overrides{}
	}
	effectiveCR := object.(*v1beta1.Verrazzano)
	if effectiveCR.Spec.Components.Fluentd != nil {
		return effectiveCR.Spec.Components.Fluentd.ValueOverrides
	}
	return []v1beta1.Overrides{}
}
