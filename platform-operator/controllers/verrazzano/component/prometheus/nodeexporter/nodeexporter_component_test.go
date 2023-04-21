// Copyright (c) 2022, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package nodeexporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verrazzano/verrazzano/pkg/bom"
	vzapi "github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1alpha1"
	"github.com/verrazzano/verrazzano/platform-operator/apis/verrazzano/v1beta1"
	"github.com/verrazzano/verrazzano/platform-operator/controllers/verrazzano/component/spi"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

const profilesRelativePath = "../../../../../manifests/profiles"

// TestIsEnabled tests the IsEnabled function for the Prometheus Node-Exporter component
func TestIsEnabled(t *testing.T) {
	falseValue := false
	trueValue := true
	tests := []struct {
		name       string
		actualCR   vzapi.Verrazzano
		expectTrue bool
	}{
		{
			// GIVEN a default Verrazzano custom resource
			// WHEN we call IsReady on the Prometheus Node-Exporter component
			// THEN the call returns true (since by default, it is enabled when Prometheus is enabled)
			name:       "Test IsEnabled when using default Verrazzano CR",
			actualCR:   vzapi.Verrazzano{},
			expectTrue: true,
		},
		{
			// GIVEN a Verrazzano custom resource with the Prometheus Node-Exporter enabled
			// WHEN we call IsReady on the Prometheus Node-Exporter component
			// THEN the call returns true
			name: "Test IsEnabled when Prometheus Node-Exporter component set to enabled",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
					},
				},
			},
			expectTrue: true,
		},
		{
			// GIVEN a Verrazzano custom resource with the Prometheus Node-Exporter disabled
			// WHEN we call IsReady on the Prometheus Node-Exporter component
			// THEN the call returns false
			name: "Test IsEnabled when Prometheus Node-Exporter component set to disabled",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &falseValue,
						},
					},
				},
			},
			expectTrue: false,
		},
		{
			// GIVEN a Verrazzano custom resource with Prometheus disabled
			// AND Prometheus Node-Exporter is not specified
			// WHEN we call IsReady on the Prometheus Node-Exporter component
			// THEN the call returns false
			name: "Test IsEnabled when Prometheus is disabled and Node-Exporter component is not specified",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						Prometheus: &vzapi.PrometheusComponent{
							Enabled: &falseValue,
						},
					},
				},
			},
			expectTrue: false,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := spi.NewFakeContext(nil, &tests[i].actualCR, nil, false, profilesRelativePath)
			assert.Equal(t, tt.expectTrue, NewComponent().IsEnabled(ctx.EffectiveCR()))
		})
	}
}

// TestAppendOverrides tests whether the prometheus.monitor.enabled setting for enabling
// service monitor, is overridden and set to true when Prometheus Operator is also enabled, but
// not otherwise
func TestAppendOverrides(t *testing.T) {
	trueValue := true
	falseValue := false
	tests := []struct {
		name              string
		actualCR          vzapi.Verrazzano
		expectedOverrides []bom.KeyValue
	}{
		{
			// GIVEN a Verrazzano custom resource with the Prometheus Node-Exporter enabled, but not Prometheus Operator
			// WHEN we call AppendOverrides on the Prometheus Node-Exporter component
			// THEN prometheus.monitor.enabled is NOT set
			name: "Test AppendOverrides when Prometheus operator is not also enabled",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
						PrometheusOperator: &vzapi.PrometheusOperatorComponent{
							Enabled: &falseValue,
						},
					},
				},
			},
			expectedOverrides: []bom.KeyValue{},
		},
		{
			// GIVEN a Verrazzano custom resource with the Prometheus Node-Exporter AND Prometheus Operator enabled
			// WHEN we call AppendOverrides on the Prometheus Node-Exporter component
			// THEN prometheus.monitor.enabled is set to true
			name: "Test AppendOverrides when Prometheus operator is also enabled",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
						PrometheusOperator: &vzapi.PrometheusOperatorComponent{
							Enabled: &trueValue,
						},
					},
				},
			},
			expectedOverrides: []bom.KeyValue{
				{Key: "prometheus.monitor.enabled", Value: "true"},
			},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := spi.NewFakeContext(nil, &tests[i].actualCR, nil, false, profilesRelativePath)
			var err error
			kvs := make([]bom.KeyValue, 0)
			kvs, err = AppendOverrides(ctx, "", "", "", kvs)
			assert.NoError(t, err)
			assert.Len(t, kvs, len(tt.expectedOverrides))
			assert.Equal(t, tt.expectedOverrides, kvs)
		})
	}
}

// TestPostInstall tests the PostInstall component function
func TestPostInstall(t *testing.T) {
	// GIVEN the Prometheus Node Exporter is being installed
	// WHEN we call the PostInstall function
	// THEN no error is returned
	client := fake.NewClientBuilder().WithScheme(testScheme).Build()
	ctx := spi.NewFakeContext(client, &vzapi.Verrazzano{}, nil, false, profilesRelativePath)
	err := NewComponent().PostInstall(ctx)
	assert.NoError(t, err)
}

// TestPostUpgrade tests the PostUpgrade component function
func TestPostUpgrade(t *testing.T) {
	// GIVEN the Prometheus Node Exporter is being upgraded
	// WHEN we call the PostUpgrade function
	// THEN no error is returned
	client := fake.NewClientBuilder().WithScheme(testScheme).Build()
	ctx := spi.NewFakeContext(client, &vzapi.Verrazzano{}, nil, false, profilesRelativePath)
	err := NewComponent().PostUpgrade(ctx)
	assert.NoError(t, err)
}

// Test isReady when it's called with component context
func TestIsReady(t *testing.T) {
	c := fake.NewClientBuilder().Build()
	ctx := spi.NewFakeContext(c, &vzapi.Verrazzano{}, nil, false, profilesRelativePath)
	assert.False(t, NewComponent().IsReady(ctx))
}

// test Monitoroverrides method
func TestMonitorOverride(t *testing.T) {
	falseValue := false
	trueValue := true
	tests := []struct {
		name       string
		actualCR   vzapi.Verrazzano
		expectTrue bool
	}{
		{
			// GIVEN a default Verrazzano custom resource
			// WHEN we call MonitorOverride on the Prometheus Node Exporter
			// THEN the call returns false
			name:       "Test MonitorOverride when using default Verrazzano CR",
			actualCR:   vzapi.Verrazzano{},
			expectTrue: false,
		},
		{
			// GIVEN a Verrazzano custom resource with the Prometheus Node Exporter enabled
			// WHEN we call MonitorOverride on the Prometheus Node Exporter component
			// THEN the call returns true
			name: "Test MonitorOverride when Prometheus Node Exporter component set to enabled",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled:          &trueValue,
							InstallOverrides: vzapi.InstallOverrides{MonitorChanges: &trueValue},
						},
					},
				},
			},
			expectTrue: true,
		},
		{
			// GIVEN a Verrazzano custom resource with the Prometheus Node Exporter disabled
			// WHEN we call MonitorOverride on the Prometheus Node Exporter component
			// THEN the call returns true
			name: "Test MonitorOverride when Prometheus Node Exporter component set to disabled",
			actualCR: vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &falseValue,
						},
					},
				},
			},
			expectTrue: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := spi.NewFakeContext(nil, &tests[i].actualCR, nil, false, profilesRelativePath)
			assert.Equal(t, tt.expectTrue, NewComponent().MonitorOverrides(ctx))
		})
	}
}

// test preinstall for component class
func TestPreInstallcomponent(t *testing.T) {
	c := fake.NewClientBuilder().Build()
	ctx := spi.NewFakeContext(c, &vzapi.Verrazzano{}, nil, true)
	assert.Nil(t, NewComponent().PreInstall(ctx))
}

// TestValidateUpdate tests the validate update functions
func TestValidateUpdate(t *testing.T) {
	falseValue := false
	trueValue := true
	tests := []struct {
		name    string
		old     *vzapi.Verrazzano
		new     *vzapi.Verrazzano
		wantErr bool
	}{
		{
			// GIVEN the component is disabled
			// WHEN the component is enabled and we call the validate update function
			// THEN no error is returned
			name: "enable",
			old: &vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &falseValue,
						},
					},
				},
			},
			new: &vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			// GIVEN the component is enabled
			// WHEN the component is disabled and we call the validate update function
			// THEN an error is returned
			name: "disable",
			old: &vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
					},
				},
			},
			new: &vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &falseValue,
						},
					},
				},
			},
			wantErr: true,
		},
		{
			// GIVEN the component is enabled
			// WHEN the component is not changed and we call the validate update function
			// THEN no error is returned
			name: "no change",
			old: &vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
					},
				},
			},
			new: &vzapi.Verrazzano{
				Spec: vzapi.VerrazzanoSpec{
					Components: vzapi.ComponentSpec{
						PrometheusNodeExporter: &vzapi.PrometheusNodeExporterComponent{
							Enabled: &trueValue,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComponent()
			if err := c.ValidateUpdate(tt.old, tt.new); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}

			v1beta1New := &v1beta1.Verrazzano{}
			v1beta1Old := &v1beta1.Verrazzano{}
			err := tt.new.ConvertTo(v1beta1New)
			assert.NoError(t, err)
			err = tt.old.ConvertTo(v1beta1Old)
			assert.NoError(t, err)

			if err := c.ValidateUpdateV1Beta1(v1beta1Old, v1beta1New); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUpdateV1Beta1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
