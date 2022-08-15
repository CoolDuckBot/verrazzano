//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2020, 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	vmcontrollerv1 "github.com/verrazzano/verrazzano-monitoring-operator/pkg/apis/vmcontroller/v1"
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Acme) DeepCopyInto(out *Acme) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Acme.
func (in *Acme) DeepCopy() *Acme {
	if in == nil {
		return nil
	}
	out := new(Acme)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationOperatorComponent) DeepCopyInto(out *ApplicationOperatorComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationOperatorComponent.
func (in *ApplicationOperatorComponent) DeepCopy() *ApplicationOperatorComponent {
	if in == nil {
		return nil
	}
	out := new(ApplicationOperatorComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProxyComponent) DeepCopyInto(out *AuthProxyComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProxyComponent.
func (in *AuthProxyComponent) DeepCopy() *AuthProxyComponent {
	if in == nil {
		return nil
	}
	out := new(AuthProxyComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CA) DeepCopyInto(out *CA) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CA.
func (in *CA) DeepCopy() *CA {
	if in == nil {
		return nil
	}
	out := new(CA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerComponent) DeepCopyInto(out *CertManagerComponent) {
	*out = *in
	out.Certificate = in.Certificate
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerComponent.
func (in *CertManagerComponent) DeepCopy() *CertManagerComponent {
	if in == nil {
		return nil
	}
	out := new(CertManagerComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.Acme = in.Acme
	out.CA = in.CA
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceOperatorComponent) DeepCopyInto(out *CoherenceOperatorComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceOperatorComponent.
func (in *CoherenceOperatorComponent) DeepCopy() *CoherenceOperatorComponent {
	if in == nil {
		return nil
	}
	out := new(CoherenceOperatorComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentSpec) DeepCopyInto(out *ComponentSpec) {
	*out = *in
	if in.CertManager != nil {
		in, out := &in.CertManager, &out.CertManager
		*out = new(CertManagerComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.CoherenceOperator != nil {
		in, out := &in.CoherenceOperator, &out.CoherenceOperator
		*out = new(CoherenceOperatorComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplicationOperator != nil {
		in, out := &in.ApplicationOperator, &out.ApplicationOperator
		*out = new(ApplicationOperatorComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.AuthProxy != nil {
		in, out := &in.AuthProxy, &out.AuthProxy
		*out = new(AuthProxyComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.OAM != nil {
		in, out := &in.OAM, &out.OAM
		*out = new(OAMComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Console != nil {
		in, out := &in.Console, &out.Console
		*out = new(ConsoleComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = new(DNSComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenSearch != nil {
		in, out := &in.OpenSearch, &out.OpenSearch
		*out = new(OpenSearchComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Fluentd != nil {
		in, out := &in.Fluentd, &out.Fluentd
		*out = new(FluentdComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Grafana != nil {
		in, out := &in.Grafana, &out.Grafana
		*out = new(GrafanaComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressNginxComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Istio != nil {
		in, out := &in.Istio, &out.Istio
		*out = new(IstioComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.JaegerOperator != nil {
		in, out := &in.JaegerOperator, &out.JaegerOperator
		*out = new(JaegerOperatorComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Kiali != nil {
		in, out := &in.Kiali, &out.Kiali
		*out = new(KialiComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Keycloak != nil {
		in, out := &in.Keycloak, &out.Keycloak
		*out = new(KeycloakComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenSearchDashboards != nil {
		in, out := &in.OpenSearchDashboards, &out.OpenSearchDashboards
		*out = new(OpenSearchDashboardsComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeStateMetrics != nil {
		in, out := &in.KubeStateMetrics, &out.KubeStateMetrics
		*out = new(KubeStateMetricsComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusAdapter != nil {
		in, out := &in.PrometheusAdapter, &out.PrometheusAdapter
		*out = new(PrometheusAdapterComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusNodeExporter != nil {
		in, out := &in.PrometheusNodeExporter, &out.PrometheusNodeExporter
		*out = new(PrometheusNodeExporterComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusOperator != nil {
		in, out := &in.PrometheusOperator, &out.PrometheusOperator
		*out = new(PrometheusOperatorComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusPushgateway != nil {
		in, out := &in.PrometheusPushgateway, &out.PrometheusPushgateway
		*out = new(PrometheusPushgatewayComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Rancher != nil {
		in, out := &in.Rancher, &out.Rancher
		*out = new(RancherComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.RancherBackup != nil {
		in, out := &in.RancherBackup, &out.RancherBackup
		*out = new(RancherBackupComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.WebLogicOperator != nil {
		in, out := &in.WebLogicOperator, &out.WebLogicOperator
		*out = new(WebLogicOperatorComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Velero != nil {
		in, out := &in.Velero, &out.Velero
		*out = new(VeleroComponent)
		(*in).DeepCopyInto(*out)
	}
	if in.Verrazzano != nil {
		in, out := &in.Verrazzano, &out.Verrazzano
		*out = new(VerrazzanoComponent)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentSpec.
func (in *ComponentSpec) DeepCopy() *ComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentStatusDetails) DeepCopyInto(out *ComponentStatusDetails) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatusDetails.
func (in *ComponentStatusDetails) DeepCopy() *ComponentStatusDetails {
	if in == nil {
		return nil
	}
	out := new(ComponentStatusDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ComponentStatusMap) DeepCopyInto(out *ComponentStatusMap) {
	{
		in := &in
		*out = make(ComponentStatusMap, len(*in))
		for key, val := range *in {
			var outVal *ComponentStatusDetails
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ComponentStatusDetails)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentStatusMap.
func (in ComponentStatusMap) DeepCopy() ComponentStatusMap {
	if in == nil {
		return nil
	}
	out := new(ComponentStatusMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleComponent) DeepCopyInto(out *ConsoleComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleComponent.
func (in *ConsoleComponent) DeepCopy() *ConsoleComponent {
	if in == nil {
		return nil
	}
	out := new(ConsoleComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSComponent) DeepCopyInto(out *DNSComponent) {
	*out = *in
	if in.Wildcard != nil {
		in, out := &in.Wildcard, &out.Wildcard
		*out = new(Wildcard)
		**out = **in
	}
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(OCI)
		**out = **in
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(External)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSComponent.
func (in *DNSComponent) DeepCopy() *DNSComponent {
	if in == nil {
		return nil
	}
	out := new(DNSComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *External) DeepCopyInto(out *External) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new External.
func (in *External) DeepCopy() *External {
	if in == nil {
		return nil
	}
	out := new(External)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdComponent) DeepCopyInto(out *FluentdComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(OciLoggingConfiguration)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdComponent.
func (in *FluentdComponent) DeepCopy() *FluentdComponent {
	if in == nil {
		return nil
	}
	out := new(FluentdComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaComponent) DeepCopyInto(out *GrafanaComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaComponent.
func (in *GrafanaComponent) DeepCopy() *GrafanaComponent {
	if in == nil {
		return nil
	}
	out := new(GrafanaComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressNginxComponent) DeepCopyInto(out *IngressNginxComponent) {
	*out = *in
	if in.IngressClassName != nil {
		in, out := &in.IngressClassName, &out.IngressClassName
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressNginxComponent.
func (in *IngressNginxComponent) DeepCopy() *IngressNginxComponent {
	if in == nil {
		return nil
	}
	out := new(IngressNginxComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallOverrides) DeepCopyInto(out *InstallOverrides) {
	*out = *in
	if in.MonitorChanges != nil {
		in, out := &in.MonitorChanges, &out.MonitorChanges
		*out = new(bool)
		**out = **in
	}
	if in.ValueOverrides != nil {
		in, out := &in.ValueOverrides, &out.ValueOverrides
		*out = make([]Overrides, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallOverrides.
func (in *InstallOverrides) DeepCopy() *InstallOverrides {
	if in == nil {
		return nil
	}
	out := new(InstallOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceInfo) DeepCopyInto(out *InstanceInfo) {
	*out = *in
	if in.ConsoleURL != nil {
		in, out := &in.ConsoleURL, &out.ConsoleURL
		*out = new(string)
		**out = **in
	}
	if in.KeyCloakURL != nil {
		in, out := &in.KeyCloakURL, &out.KeyCloakURL
		*out = new(string)
		**out = **in
	}
	if in.RancherURL != nil {
		in, out := &in.RancherURL, &out.RancherURL
		*out = new(string)
		**out = **in
	}
	if in.OpenSearchURL != nil {
		in, out := &in.OpenSearchURL, &out.OpenSearchURL
		*out = new(string)
		**out = **in
	}
	if in.OpenSearchDashboardsURL != nil {
		in, out := &in.OpenSearchDashboardsURL, &out.OpenSearchDashboardsURL
		*out = new(string)
		**out = **in
	}
	if in.GrafanaURL != nil {
		in, out := &in.GrafanaURL, &out.GrafanaURL
		*out = new(string)
		**out = **in
	}
	if in.PrometheusURL != nil {
		in, out := &in.PrometheusURL, &out.PrometheusURL
		*out = new(string)
		**out = **in
	}
	if in.KialiURL != nil {
		in, out := &in.KialiURL, &out.KialiURL
		*out = new(string)
		**out = **in
	}
	if in.JaegerURL != nil {
		in, out := &in.JaegerURL, &out.JaegerURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceInfo.
func (in *InstanceInfo) DeepCopy() *InstanceInfo {
	if in == nil {
		return nil
	}
	out := new(InstanceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioComponent) DeepCopyInto(out *IstioComponent) {
	*out = *in
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.InjectionEnabled != nil {
		in, out := &in.InjectionEnabled, &out.InjectionEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioComponent.
func (in *IstioComponent) DeepCopy() *IstioComponent {
	if in == nil {
		return nil
	}
	out := new(IstioComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioIngressSection) DeepCopyInto(out *IstioIngressSection) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioIngressSection.
func (in *IstioIngressSection) DeepCopy() *IstioIngressSection {
	if in == nil {
		return nil
	}
	out := new(IstioIngressSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JaegerOperatorComponent) DeepCopyInto(out *JaegerOperatorComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JaegerOperatorComponent.
func (in *JaegerOperatorComponent) DeepCopy() *JaegerOperatorComponent {
	if in == nil {
		return nil
	}
	out := new(JaegerOperatorComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakComponent) DeepCopyInto(out *KeycloakComponent) {
	*out = *in
	in.MySQL.DeepCopyInto(&out.MySQL)
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakComponent.
func (in *KeycloakComponent) DeepCopy() *KeycloakComponent {
	if in == nil {
		return nil
	}
	out := new(KeycloakComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KialiComponent) DeepCopyInto(out *KialiComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KialiComponent.
func (in *KialiComponent) DeepCopy() *KialiComponent {
	if in == nil {
		return nil
	}
	out := new(KialiComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStateMetricsComponent) DeepCopyInto(out *KubeStateMetricsComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStateMetricsComponent.
func (in *KubeStateMetricsComponent) DeepCopy() *KubeStateMetricsComponent {
	if in == nil {
		return nil
	}
	out := new(KubeStateMetricsComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLComponent) DeepCopyInto(out *MySQLComponent) {
	*out = *in
	if in.VolumeSource != nil {
		in, out := &in.VolumeSource, &out.VolumeSource
		*out = new(v1.VolumeSource)
		(*in).DeepCopyInto(*out)
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLComponent.
func (in *MySQLComponent) DeepCopy() *MySQLComponent {
	if in == nil {
		return nil
	}
	out := new(MySQLComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAMComponent) DeepCopyInto(out *OAMComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAMComponent.
func (in *OAMComponent) DeepCopy() *OAMComponent {
	if in == nil {
		return nil
	}
	out := new(OAMComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCI) DeepCopyInto(out *OCI) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCI.
func (in *OCI) DeepCopy() *OCI {
	if in == nil {
		return nil
	}
	out := new(OCI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OciLoggingConfiguration) DeepCopyInto(out *OciLoggingConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OciLoggingConfiguration.
func (in *OciLoggingConfiguration) DeepCopy() *OciLoggingConfiguration {
	if in == nil {
		return nil
	}
	out := new(OciLoggingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSearchComponent) DeepCopyInto(out *OpenSearchComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]vmcontrollerv1.IndexManagementPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]OpenSearchNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSearchComponent.
func (in *OpenSearchComponent) DeepCopy() *OpenSearchComponent {
	if in == nil {
		return nil
	}
	out := new(OpenSearchComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSearchDashboardsComponent) DeepCopyInto(out *OpenSearchDashboardsComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSearchDashboardsComponent.
func (in *OpenSearchDashboardsComponent) DeepCopy() *OpenSearchDashboardsComponent {
	if in == nil {
		return nil
	}
	out := new(OpenSearchDashboardsComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSearchNode) DeepCopyInto(out *OpenSearchNode) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]vmcontrollerv1.NodeRole, len(*in))
		copy(*out, *in)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(OpenSearchNodeStorage)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSearchNode.
func (in *OpenSearchNode) DeepCopy() *OpenSearchNode {
	if in == nil {
		return nil
	}
	out := new(OpenSearchNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenSearchNodeStorage) DeepCopyInto(out *OpenSearchNodeStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenSearchNodeStorage.
func (in *OpenSearchNodeStorage) DeepCopy() *OpenSearchNodeStorage {
	if in == nil {
		return nil
	}
	out := new(OpenSearchNodeStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overrides) DeepCopyInto(out *Overrides) {
	*out = *in
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overrides.
func (in *Overrides) DeepCopy() *Overrides {
	if in == nil {
		return nil
	}
	out := new(Overrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAdapterComponent) DeepCopyInto(out *PrometheusAdapterComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAdapterComponent.
func (in *PrometheusAdapterComponent) DeepCopy() *PrometheusAdapterComponent {
	if in == nil {
		return nil
	}
	out := new(PrometheusAdapterComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusComponent) DeepCopyInto(out *PrometheusComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusComponent.
func (in *PrometheusComponent) DeepCopy() *PrometheusComponent {
	if in == nil {
		return nil
	}
	out := new(PrometheusComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusNodeExporterComponent) DeepCopyInto(out *PrometheusNodeExporterComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusNodeExporterComponent.
func (in *PrometheusNodeExporterComponent) DeepCopy() *PrometheusNodeExporterComponent {
	if in == nil {
		return nil
	}
	out := new(PrometheusNodeExporterComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusOperatorComponent) DeepCopyInto(out *PrometheusOperatorComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusOperatorComponent.
func (in *PrometheusOperatorComponent) DeepCopy() *PrometheusOperatorComponent {
	if in == nil {
		return nil
	}
	out := new(PrometheusOperatorComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusPushgatewayComponent) DeepCopyInto(out *PrometheusPushgatewayComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusPushgatewayComponent.
func (in *PrometheusPushgatewayComponent) DeepCopy() *PrometheusPushgatewayComponent {
	if in == nil {
		return nil
	}
	out := new(PrometheusPushgatewayComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RancherBackupComponent) DeepCopyInto(out *RancherBackupComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RancherBackupComponent.
func (in *RancherBackupComponent) DeepCopy() *RancherBackupComponent {
	if in == nil {
		return nil
	}
	out := new(RancherBackupComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RancherComponent) DeepCopyInto(out *RancherComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RancherComponent.
func (in *RancherComponent) DeepCopy() *RancherComponent {
	if in == nil {
		return nil
	}
	out := new(RancherComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySpec) DeepCopyInto(out *SecuritySpec) {
	*out = *in
	if in.AdminSubjects != nil {
		in, out := &in.AdminSubjects, &out.AdminSubjects
		*out = make([]rbacv1.Subject, len(*in))
		copy(*out, *in)
	}
	if in.MonitorSubjects != nil {
		in, out := &in.MonitorSubjects, &out.MonitorSubjects
		*out = make([]rbacv1.Subject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySpec.
func (in *SecuritySpec) DeepCopy() *SecuritySpec {
	if in == nil {
		return nil
	}
	out := new(SecuritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroComponent) DeepCopyInto(out *VeleroComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroComponent.
func (in *VeleroComponent) DeepCopy() *VeleroComponent {
	if in == nil {
		return nil
	}
	out := new(VeleroComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Verrazzano) DeepCopyInto(out *Verrazzano) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Verrazzano.
func (in *Verrazzano) DeepCopy() *Verrazzano {
	if in == nil {
		return nil
	}
	out := new(Verrazzano)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Verrazzano) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoComponent) DeepCopyInto(out *VerrazzanoComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoComponent.
func (in *VerrazzanoComponent) DeepCopy() *VerrazzanoComponent {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoList) DeepCopyInto(out *VerrazzanoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Verrazzano, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoList.
func (in *VerrazzanoList) DeepCopy() *VerrazzanoList {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerrazzanoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoSpec) DeepCopyInto(out *VerrazzanoSpec) {
	*out = *in
	in.Components.DeepCopyInto(&out.Components)
	in.Security.DeepCopyInto(&out.Security)
	if in.DefaultVolumeSource != nil {
		in, out := &in.DefaultVolumeSource, &out.DefaultVolumeSource
		*out = new(v1.VolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimSpecTemplates != nil {
		in, out := &in.VolumeClaimSpecTemplates, &out.VolumeClaimSpecTemplates
		*out = make([]VolumeClaimSpecTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoSpec.
func (in *VerrazzanoSpec) DeepCopy() *VerrazzanoSpec {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerrazzanoStatus) DeepCopyInto(out *VerrazzanoStatus) {
	*out = *in
	if in.VerrazzanoInstance != nil {
		in, out := &in.VerrazzanoInstance, &out.VerrazzanoInstance
		*out = new(InstanceInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make(ComponentStatusMap, len(*in))
		for key, val := range *in {
			var outVal *ComponentStatusDetails
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ComponentStatusDetails)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerrazzanoStatus.
func (in *VerrazzanoStatus) DeepCopy() *VerrazzanoStatus {
	if in == nil {
		return nil
	}
	out := new(VerrazzanoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeClaimSpecTemplate) DeepCopyInto(out *VolumeClaimSpecTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeClaimSpecTemplate.
func (in *VolumeClaimSpecTemplate) DeepCopy() *VolumeClaimSpecTemplate {
	if in == nil {
		return nil
	}
	out := new(VolumeClaimSpecTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMount.
func (in *VolumeMount) DeepCopy() *VolumeMount {
	if in == nil {
		return nil
	}
	out := new(VolumeMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebLogicOperatorComponent) DeepCopyInto(out *WebLogicOperatorComponent) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	in.InstallOverrides.DeepCopyInto(&out.InstallOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebLogicOperatorComponent.
func (in *WebLogicOperatorComponent) DeepCopy() *WebLogicOperatorComponent {
	if in == nil {
		return nil
	}
	out := new(WebLogicOperatorComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wildcard) DeepCopyInto(out *Wildcard) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wildcard.
func (in *Wildcard) DeepCopy() *Wildcard {
	if in == nil {
		return nil
	}
	out := new(Wildcard)
	in.DeepCopyInto(out)
	return out
}