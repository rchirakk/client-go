// +build !ignore_autogenerated

/*
Copyright 2016 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	api "github.com/FlorianOtel/client-go/pkg/api"
	v1 "github.com/FlorianOtel/client-go/pkg/api/v1"
	meta_v1 "github.com/FlorianOtel/client-go/pkg/apis/meta/v1"
	policy "github.com/FlorianOtel/client-go/pkg/apis/policy"
	conversion "github.com/FlorianOtel/client-go/pkg/conversion"
	runtime "github.com/FlorianOtel/client-go/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_Eviction_To_policy_Eviction,
		Convert_policy_Eviction_To_v1beta1_Eviction,
		Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget,
		Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget,
		Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList,
		Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList,
		Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec,
		Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec,
		Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus,
		Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus,
	)
}

func autoConvert_v1beta1_Eviction_To_policy_Eviction(in *Eviction, out *policy.Eviction, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.DeleteOptions = (*api.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

func Convert_v1beta1_Eviction_To_policy_Eviction(in *Eviction, out *policy.Eviction, s conversion.Scope) error {
	return autoConvert_v1beta1_Eviction_To_policy_Eviction(in, out, s)
}

func autoConvert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *Eviction, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

func Convert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *Eviction, s conversion.Scope) error {
	return autoConvert_policy_Eviction_To_v1beta1_Eviction(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in, out, s)
}

func autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *PodDisruptionBudget, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.ObjectMeta, &out.ObjectMeta, 0); err != nil {
		return err
	}
	if err := Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *PodDisruptionBudget, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]policy.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PodDisruptionBudget)(unsafe.Pointer(&in.Items))
	return nil
}

func Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = in.MinAvailable
	out.Selector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.Selector))
	return nil
}

func Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = in.MinAvailable
	out.Selector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.Selector))
	return nil
}

func Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]meta_v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	return nil
}

func Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]meta_v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	return nil
}

func Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in, out, s)
}
