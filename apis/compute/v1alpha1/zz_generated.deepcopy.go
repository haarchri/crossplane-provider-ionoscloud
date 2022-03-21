//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CubeServer) DeepCopyInto(out *CubeServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CubeServer.
func (in *CubeServer) DeepCopy() *CubeServer {
	if in == nil {
		return nil
	}
	out := new(CubeServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CubeServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CubeServerList) DeepCopyInto(out *CubeServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CubeServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CubeServerList.
func (in *CubeServerList) DeepCopy() *CubeServerList {
	if in == nil {
		return nil
	}
	out := new(CubeServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CubeServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CubeServerProperties) DeepCopyInto(out *CubeServerProperties) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
	out.Template = in.Template
	in.DasVolumeProperties.DeepCopyInto(&out.DasVolumeProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CubeServerProperties.
func (in *CubeServerProperties) DeepCopy() *CubeServerProperties {
	if in == nil {
		return nil
	}
	out := new(CubeServerProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CubeServerSpec) DeepCopyInto(out *CubeServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CubeServerSpec.
func (in *CubeServerSpec) DeepCopy() *CubeServerSpec {
	if in == nil {
		return nil
	}
	out := new(CubeServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DasVolumeProperties) DeepCopyInto(out *DasVolumeProperties) {
	*out = *in
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DasVolumeProperties.
func (in *DasVolumeProperties) DeepCopy() *DasVolumeProperties {
	if in == nil {
		return nil
	}
	out := new(DasVolumeProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Datacenter) DeepCopyInto(out *Datacenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Datacenter.
func (in *Datacenter) DeepCopy() *Datacenter {
	if in == nil {
		return nil
	}
	out := new(Datacenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Datacenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterConfig) DeepCopyInto(out *DatacenterConfig) {
	*out = *in
	if in.DatacenterIDRef != nil {
		in, out := &in.DatacenterIDRef, &out.DatacenterIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.DatacenterIDSelector != nil {
		in, out := &in.DatacenterIDSelector, &out.DatacenterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterConfig.
func (in *DatacenterConfig) DeepCopy() *DatacenterConfig {
	if in == nil {
		return nil
	}
	out := new(DatacenterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterList) DeepCopyInto(out *DatacenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Datacenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterList.
func (in *DatacenterList) DeepCopy() *DatacenterList {
	if in == nil {
		return nil
	}
	out := new(DatacenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatacenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterObservation) DeepCopyInto(out *DatacenterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterObservation.
func (in *DatacenterObservation) DeepCopy() *DatacenterObservation {
	if in == nil {
		return nil
	}
	out := new(DatacenterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterParameters) DeepCopyInto(out *DatacenterParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterParameters.
func (in *DatacenterParameters) DeepCopy() *DatacenterParameters {
	if in == nil {
		return nil
	}
	out := new(DatacenterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterSpec) DeepCopyInto(out *DatacenterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterSpec.
func (in *DatacenterSpec) DeepCopy() *DatacenterSpec {
	if in == nil {
		return nil
	}
	out := new(DatacenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterStatus) DeepCopyInto(out *DatacenterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterStatus.
func (in *DatacenterStatus) DeepCopy() *DatacenterStatus {
	if in == nil {
		return nil
	}
	out := new(DatacenterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRule) DeepCopyInto(out *FirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRule.
func (in *FirewallRule) DeepCopy() *FirewallRule {
	if in == nil {
		return nil
	}
	out := new(FirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleConfig) DeepCopyInto(out *FirewallRuleConfig) {
	*out = *in
	if in.FirewallRuleIDRef != nil {
		in, out := &in.FirewallRuleIDRef, &out.FirewallRuleIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.FirewallRuleIDSelector != nil {
		in, out := &in.FirewallRuleIDSelector, &out.FirewallRuleIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleConfig.
func (in *FirewallRuleConfig) DeepCopy() *FirewallRuleConfig {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleList) DeepCopyInto(out *FirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleList.
func (in *FirewallRuleList) DeepCopy() *FirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleObservation) DeepCopyInto(out *FirewallRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleObservation.
func (in *FirewallRuleObservation) DeepCopy() *FirewallRuleObservation {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleParameters) DeepCopyInto(out *FirewallRuleParameters) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
	in.ServerCfg.DeepCopyInto(&out.ServerCfg)
	in.NicCfg.DeepCopyInto(&out.NicCfg)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleParameters.
func (in *FirewallRuleParameters) DeepCopy() *FirewallRuleParameters {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleSpec) DeepCopyInto(out *FirewallRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleSpec.
func (in *FirewallRuleSpec) DeepCopy() *FirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleStatus) DeepCopyInto(out *FirewallRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleStatus.
func (in *FirewallRuleStatus) DeepCopy() *FirewallRuleStatus {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPBlock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlockList) DeepCopyInto(out *IPBlockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPBlock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlockList.
func (in *IPBlockList) DeepCopy() *IPBlockList {
	if in == nil {
		return nil
	}
	out := new(IPBlockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPBlockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlockObservation) DeepCopyInto(out *IPBlockObservation) {
	*out = *in
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlockObservation.
func (in *IPBlockObservation) DeepCopy() *IPBlockObservation {
	if in == nil {
		return nil
	}
	out := new(IPBlockObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlockParameters) DeepCopyInto(out *IPBlockParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlockParameters.
func (in *IPBlockParameters) DeepCopy() *IPBlockParameters {
	if in == nil {
		return nil
	}
	out := new(IPBlockParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlockSpec) DeepCopyInto(out *IPBlockSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlockSpec.
func (in *IPBlockSpec) DeepCopy() *IPBlockSpec {
	if in == nil {
		return nil
	}
	out := new(IPBlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlockStatus) DeepCopyInto(out *IPBlockStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlockStatus.
func (in *IPBlockStatus) DeepCopy() *IPBlockStatus {
	if in == nil {
		return nil
	}
	out := new(IPBlockStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPFailover) DeepCopyInto(out *IPFailover) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPFailover.
func (in *IPFailover) DeepCopy() *IPFailover {
	if in == nil {
		return nil
	}
	out := new(IPFailover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPFailover) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPFailoverList) DeepCopyInto(out *IPFailoverList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPFailover, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPFailoverList.
func (in *IPFailoverList) DeepCopy() *IPFailoverList {
	if in == nil {
		return nil
	}
	out := new(IPFailoverList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPFailoverList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPFailoverObservation) DeepCopyInto(out *IPFailoverObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPFailoverObservation.
func (in *IPFailoverObservation) DeepCopy() *IPFailoverObservation {
	if in == nil {
		return nil
	}
	out := new(IPFailoverObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPFailoverParameters) DeepCopyInto(out *IPFailoverParameters) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
	in.LanCfg.DeepCopyInto(&out.LanCfg)
	in.NicCfg.DeepCopyInto(&out.NicCfg)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPFailoverParameters.
func (in *IPFailoverParameters) DeepCopy() *IPFailoverParameters {
	if in == nil {
		return nil
	}
	out := new(IPFailoverParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPFailoverSpec) DeepCopyInto(out *IPFailoverSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPFailoverSpec.
func (in *IPFailoverSpec) DeepCopy() *IPFailoverSpec {
	if in == nil {
		return nil
	}
	out := new(IPFailoverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPFailoverStatus) DeepCopyInto(out *IPFailoverStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPFailoverStatus.
func (in *IPFailoverStatus) DeepCopy() *IPFailoverStatus {
	if in == nil {
		return nil
	}
	out := new(IPFailoverStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lan) DeepCopyInto(out *Lan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lan.
func (in *Lan) DeepCopy() *Lan {
	if in == nil {
		return nil
	}
	out := new(Lan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Lan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LanConfig) DeepCopyInto(out *LanConfig) {
	*out = *in
	if in.LanIDRef != nil {
		in, out := &in.LanIDRef, &out.LanIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.LanIDSelector != nil {
		in, out := &in.LanIDSelector, &out.LanIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LanConfig.
func (in *LanConfig) DeepCopy() *LanConfig {
	if in == nil {
		return nil
	}
	out := new(LanConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LanList) DeepCopyInto(out *LanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Lan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LanList.
func (in *LanList) DeepCopy() *LanList {
	if in == nil {
		return nil
	}
	out := new(LanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LanObservation) DeepCopyInto(out *LanObservation) {
	*out = *in
	if in.IPFailovers != nil {
		in, out := &in.IPFailovers, &out.IPFailovers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LanObservation.
func (in *LanObservation) DeepCopy() *LanObservation {
	if in == nil {
		return nil
	}
	out := new(LanObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LanParameters) DeepCopyInto(out *LanParameters) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LanParameters.
func (in *LanParameters) DeepCopy() *LanParameters {
	if in == nil {
		return nil
	}
	out := new(LanParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LanSpec) DeepCopyInto(out *LanSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LanSpec.
func (in *LanSpec) DeepCopy() *LanSpec {
	if in == nil {
		return nil
	}
	out := new(LanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LanStatus) DeepCopyInto(out *LanStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LanStatus.
func (in *LanStatus) DeepCopy() *LanStatus {
	if in == nil {
		return nil
	}
	out := new(LanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nic) DeepCopyInto(out *Nic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nic.
func (in *Nic) DeepCopy() *Nic {
	if in == nil {
		return nil
	}
	out := new(Nic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Nic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicConfig) DeepCopyInto(out *NicConfig) {
	*out = *in
	if in.NicIDRef != nil {
		in, out := &in.NicIDRef, &out.NicIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.NicIDSelector != nil {
		in, out := &in.NicIDSelector, &out.NicIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicConfig.
func (in *NicConfig) DeepCopy() *NicConfig {
	if in == nil {
		return nil
	}
	out := new(NicConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicList) DeepCopyInto(out *NicList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Nic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicList.
func (in *NicList) DeepCopy() *NicList {
	if in == nil {
		return nil
	}
	out := new(NicList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NicList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicObservation) DeepCopyInto(out *NicObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicObservation.
func (in *NicObservation) DeepCopy() *NicObservation {
	if in == nil {
		return nil
	}
	out := new(NicObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicParameters) DeepCopyInto(out *NicParameters) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
	in.ServerCfg.DeepCopyInto(&out.ServerCfg)
	in.LanCfg.DeepCopyInto(&out.LanCfg)
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicParameters.
func (in *NicParameters) DeepCopy() *NicParameters {
	if in == nil {
		return nil
	}
	out := new(NicParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicSpec) DeepCopyInto(out *NicSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicSpec.
func (in *NicSpec) DeepCopy() *NicSpec {
	if in == nil {
		return nil
	}
	out := new(NicSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NicStatus) DeepCopyInto(out *NicStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NicStatus.
func (in *NicStatus) DeepCopy() *NicStatus {
	if in == nil {
		return nil
	}
	out := new(NicStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfig) DeepCopyInto(out *ServerConfig) {
	*out = *in
	if in.ServerIDRef != nil {
		in, out := &in.ServerIDRef, &out.ServerIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ServerIDSelector != nil {
		in, out := &in.ServerIDSelector, &out.ServerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfig.
func (in *ServerConfig) DeepCopy() *ServerConfig {
	if in == nil {
		return nil
	}
	out := new(ServerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerObservation) DeepCopyInto(out *ServerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerObservation.
func (in *ServerObservation) DeepCopy() *ServerObservation {
	if in == nil {
		return nil
	}
	out := new(ServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerParameters) DeepCopyInto(out *ServerParameters) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
	in.VolumeCfg.DeepCopyInto(&out.VolumeCfg)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerParameters.
func (in *ServerParameters) DeepCopy() *ServerParameters {
	if in == nil {
		return nil
	}
	out := new(ServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Volume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConfig) DeepCopyInto(out *VolumeConfig) {
	*out = *in
	if in.VolumeIDRef != nil {
		in, out := &in.VolumeIDRef, &out.VolumeIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.VolumeIDSelector != nil {
		in, out := &in.VolumeIDSelector, &out.VolumeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConfig.
func (in *VolumeConfig) DeepCopy() *VolumeConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeList) DeepCopyInto(out *VolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeList.
func (in *VolumeList) DeepCopy() *VolumeList {
	if in == nil {
		return nil
	}
	out := new(VolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeObservation) DeepCopyInto(out *VolumeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeObservation.
func (in *VolumeObservation) DeepCopy() *VolumeObservation {
	if in == nil {
		return nil
	}
	out := new(VolumeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeParameters) DeepCopyInto(out *VolumeParameters) {
	*out = *in
	in.DatacenterCfg.DeepCopyInto(&out.DatacenterCfg)
	if in.SSHKeys != nil {
		in, out := &in.SSHKeys, &out.SSHKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeParameters.
func (in *VolumeParameters) DeepCopy() *VolumeParameters {
	if in == nil {
		return nil
	}
	out := new(VolumeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeStatus) DeepCopyInto(out *VolumeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeStatus.
func (in *VolumeStatus) DeepCopy() *VolumeStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeStatus)
	in.DeepCopyInto(out)
	return out
}
