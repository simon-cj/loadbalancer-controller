/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/caicloud/loadbalancer-controller/pkg/kong/client/clientset/versioned/typed/configuration/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConfigurationV1 struct {
	*testing.Fake
}

func (c *FakeConfigurationV1) KongClusterPlugins() v1.KongClusterPluginInterface {
	return &FakeKongClusterPlugins{c}
}

func (c *FakeConfigurationV1) KongConsumers(namespace string) v1.KongConsumerInterface {
	return &FakeKongConsumers{c, namespace}
}

func (c *FakeConfigurationV1) KongCredentials(namespace string) v1.KongCredentialInterface {
	return &FakeKongCredentials{c, namespace}
}

func (c *FakeConfigurationV1) KongIngresses(namespace string) v1.KongIngressInterface {
	return &FakeKongIngresses{c, namespace}
}

func (c *FakeConfigurationV1) KongPlugins(namespace string) v1.KongPluginInterface {
	return &FakeKongPlugins{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConfigurationV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
