//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2020, 2023, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartDependency) DeepCopyInto(out *ChartDependency) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartDependency.
func (in *ChartDependency) DeepCopy() *ChartDependency {
	if in == nil {
		return nil
	}
	out := new(ChartDependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartVersion) DeepCopyInto(out *ChartVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartVersion.
func (in *ChartVersion) DeepCopy() *ChartVersion {
	if in == nil {
		return nil
	}
	out := new(ChartVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Module) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleCondition) DeepCopyInto(out *ModuleCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleCondition.
func (in *ModuleCondition) DeepCopy() *ModuleCondition {
	if in == nil {
		return nil
	}
	out := new(ModuleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleDefinition) DeepCopyInto(out *ModuleDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleDefinition.
func (in *ModuleDefinition) DeepCopy() *ModuleDefinition {
	if in == nil {
		return nil
	}
	out := new(ModuleDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModuleDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleDefinitionCondition) DeepCopyInto(out *ModuleDefinitionCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleDefinitionCondition.
func (in *ModuleDefinitionCondition) DeepCopy() *ModuleDefinitionCondition {
	if in == nil {
		return nil
	}
	out := new(ModuleDefinitionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleDefinitionList) DeepCopyInto(out *ModuleDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Platform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleDefinitionList.
func (in *ModuleDefinitionList) DeepCopy() *ModuleDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ModuleDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModuleDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleDefinitionSpec) DeepCopyInto(out *ModuleDefinitionSpec) {
	*out = *in
	in.OperatorDefinition.DeepCopyInto(&out.OperatorDefinition)
	if in.ModuleDependencies != nil {
		in, out := &in.ModuleDependencies, &out.ModuleDependencies
		*out = make([]ChartDependency, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleDefinitionSpec.
func (in *ModuleDefinitionSpec) DeepCopy() *ModuleDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleDefinitionStatus) DeepCopyInto(out *ModuleDefinitionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleDefinitionStatus.
func (in *ModuleDefinitionStatus) DeepCopy() *ModuleDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ModuleDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleList) DeepCopyInto(out *ModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleList.
func (in *ModuleList) DeepCopy() *ModuleList {
	if in == nil {
		return nil
	}
	out := new(ModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleSpec) DeepCopyInto(out *ModuleSpec) {
	*out = *in
	if in.ChartName != nil {
		in, out := &in.ChartName, &out.ChartName
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(PlatformSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.TargetNamespace != nil {
		in, out := &in.TargetNamespace, &out.TargetNamespace
		*out = new(string)
		**out = **in
	}
	if in.Reconcile != nil {
		in, out := &in.Reconcile, &out.Reconcile
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleSpec.
func (in *ModuleSpec) DeepCopy() *ModuleSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleStatus) DeepCopyInto(out *ModuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ModuleCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleStatus.
func (in *ModuleStatus) DeepCopy() *ModuleStatus {
	if in == nil {
		return nil
	}
	out := new(ModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorDefinition) DeepCopyInto(out *OperatorDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorDefinition.
func (in *OperatorDefinition) DeepCopy() *OperatorDefinition {
	if in == nil {
		return nil
	}
	out := new(OperatorDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorDefinitionList) DeepCopyInto(out *OperatorDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Platform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorDefinitionList.
func (in *OperatorDefinitionList) DeepCopy() *OperatorDefinitionList {
	if in == nil {
		return nil
	}
	out := new(OperatorDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorDefinitionSpec) DeepCopyInto(out *OperatorDefinitionSpec) {
	*out = *in
	if in.CRDDependencies != nil {
		in, out := &in.CRDDependencies, &out.CRDDependencies
		*out = make([]ChartDependency, len(*in))
		copy(*out, *in)
	}
	if in.OperatorDependencies != nil {
		in, out := &in.OperatorDependencies, &out.OperatorDependencies
		*out = make([]ChartDependency, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorDefinitionSpec.
func (in *OperatorDefinitionSpec) DeepCopy() *OperatorDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorDefinitionStatus) DeepCopyInto(out *OperatorDefinitionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorDefinitionStatus.
func (in *OperatorDefinitionStatus) DeepCopy() *OperatorDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Platform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformCondition) DeepCopyInto(out *PlatformCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformCondition.
func (in *PlatformCondition) DeepCopy() *PlatformCondition {
	if in == nil {
		return nil
	}
	out := new(PlatformCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDefinition) DeepCopyInto(out *PlatformDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDefinition.
func (in *PlatformDefinition) DeepCopy() *PlatformDefinition {
	if in == nil {
		return nil
	}
	out := new(PlatformDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDefinitionList) DeepCopyInto(out *PlatformDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlatformDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDefinitionList.
func (in *PlatformDefinitionList) DeepCopy() *PlatformDefinitionList {
	if in == nil {
		return nil
	}
	out := new(PlatformDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDefinitionSpec) DeepCopyInto(out *PlatformDefinitionSpec) {
	*out = *in
	if in.CRDVersions != nil {
		in, out := &in.CRDVersions, &out.CRDVersions
		*out = make([]ChartVersion, len(*in))
		copy(*out, *in)
	}
	if in.OperatorVersions != nil {
		in, out := &in.OperatorVersions, &out.OperatorVersions
		*out = make([]ChartVersion, len(*in))
		copy(*out, *in)
	}
	if in.ModuleVersions != nil {
		in, out := &in.ModuleVersions, &out.ModuleVersions
		*out = make([]ChartVersion, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDefinitionSpec.
func (in *PlatformDefinitionSpec) DeepCopy() *PlatformDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformDefintionStatus) DeepCopyInto(out *PlatformDefintionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformDefintionStatus.
func (in *PlatformDefintionStatus) DeepCopy() *PlatformDefintionStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformDefintionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformList) DeepCopyInto(out *PlatformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Platform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformList.
func (in *PlatformList) DeepCopy() *PlatformList {
	if in == nil {
		return nil
	}
	out := new(PlatformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformSource) DeepCopyInto(out *PlatformSource) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformSource.
func (in *PlatformSource) DeepCopy() *PlatformSource {
	if in == nil {
		return nil
	}
	out := new(PlatformSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformSpec) DeepCopyInto(out *PlatformSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SourceType, len(*in))
		copy(*out, *in)
	}
	if in.Reconcile != nil {
		in, out := &in.Reconcile, &out.Reconcile
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformSpec.
func (in *PlatformSpec) DeepCopy() *PlatformSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformStatus) DeepCopyInto(out *PlatformStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformStatus.
func (in *PlatformStatus) DeepCopy() *PlatformStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceType) DeepCopyInto(out *SourceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceType.
func (in *SourceType) DeepCopy() *SourceType {
	if in == nil {
		return nil
	}
	out := new(SourceType)
	in.DeepCopyInto(out)
	return out
}
