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

package fake

import (
	v1alpha1 "github.com/FlorianOtel/client-go/kubernetes/typed/policy/v1alpha1"
	rest "github.com/FlorianOtel/client-go/rest"
	testing "github.com/FlorianOtel/client-go/testing"
)

type FakePolicyV1alpha1 struct {
	*testing.Fake
}

func (c *FakePolicyV1alpha1) PodDisruptionBudgets(namespace string) v1alpha1.PodDisruptionBudgetInterface {
	return &FakePodDisruptionBudgets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePolicyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
