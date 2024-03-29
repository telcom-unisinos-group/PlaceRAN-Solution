// +build !ignore_autogenerated

/*


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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Algorithm) DeepCopyInto(out *Algorithm) {
	*out = *in
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.EndTimestamp != nil {
		in, out := &in.EndTimestamp, &out.EndTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Algorithm.
func (in *Algorithm) DeepCopy() *Algorithm {
	if in == nil {
		return nil
	}
	out := new(Algorithm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANDeployer) DeepCopyInto(out *RANDeployer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANDeployer.
func (in *RANDeployer) DeepCopy() *RANDeployer {
	if in == nil {
		return nil
	}
	out := new(RANDeployer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RANDeployer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANDeployerList) DeepCopyInto(out *RANDeployerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RANDeployer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANDeployerList.
func (in *RANDeployerList) DeepCopy() *RANDeployerList {
	if in == nil {
		return nil
	}
	out := new(RANDeployerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RANDeployerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANDeployerSpec) DeepCopyInto(out *RANDeployerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANDeployerSpec.
func (in *RANDeployerSpec) DeepCopy() *RANDeployerSpec {
	if in == nil {
		return nil
	}
	out := new(RANDeployerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANDeployerStatus) DeepCopyInto(out *RANDeployerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANDeployerStatus.
func (in *RANDeployerStatus) DeepCopy() *RANDeployerStatus {
	if in == nil {
		return nil
	}
	out := new(RANDeployerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANPlacer) DeepCopyInto(out *RANPlacer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANPlacer.
func (in *RANPlacer) DeepCopy() *RANPlacer {
	if in == nil {
		return nil
	}
	out := new(RANPlacer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RANPlacer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANPlacerList) DeepCopyInto(out *RANPlacerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RANPlacer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANPlacerList.
func (in *RANPlacerList) DeepCopy() *RANPlacerList {
	if in == nil {
		return nil
	}
	out := new(RANPlacerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RANPlacerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANPlacerSpec) DeepCopyInto(out *RANPlacerSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANPlacerSpec.
func (in *RANPlacerSpec) DeepCopy() *RANPlacerSpec {
	if in == nil {
		return nil
	}
	out := new(RANPlacerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RANPlacerStatus) DeepCopyInto(out *RANPlacerStatus) {
	*out = *in
	in.Algorithm.DeepCopyInto(&out.Algorithm)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RANPlacerStatus.
func (in *RANPlacerStatus) DeepCopy() *RANPlacerStatus {
	if in == nil {
		return nil
	}
	out := new(RANPlacerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesRequest) DeepCopyInto(out *ResourcesRequest) {
	*out = *in
	if in.CU != nil {
		in, out := &in.CU, &out.CU
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.DU != nil {
		in, out := &in.DU, &out.DU
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.RU != nil {
		in, out := &in.RU, &out.RU
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesRequest.
func (in *ResourcesRequest) DeepCopy() *ResourcesRequest {
	if in == nil {
		return nil
	}
	out := new(ResourcesRequest)
	in.DeepCopyInto(out)
	return out
}
