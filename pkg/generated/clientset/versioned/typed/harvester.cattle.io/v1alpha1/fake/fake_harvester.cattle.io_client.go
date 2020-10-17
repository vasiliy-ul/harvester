/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/rancher/harvester/pkg/generated/clientset/versioned/typed/harvester.cattle.io/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHarvesterV1alpha1 struct {
	*testing.Fake
}

func (c *FakeHarvesterV1alpha1) KeyPairs(namespace string) v1alpha1.KeyPairInterface {
	return &FakeKeyPairs{c, namespace}
}

func (c *FakeHarvesterV1alpha1) Settings() v1alpha1.SettingInterface {
	return &FakeSettings{c}
}

func (c *FakeHarvesterV1alpha1) Users() v1alpha1.UserInterface {
	return &FakeUsers{c}
}

func (c *FakeHarvesterV1alpha1) VirtualMachineImages(namespace string) v1alpha1.VirtualMachineImageInterface {
	return &FakeVirtualMachineImages{c, namespace}
}

func (c *FakeHarvesterV1alpha1) VirtualMachineTemplates(namespace string) v1alpha1.VirtualMachineTemplateInterface {
	return &FakeVirtualMachineTemplates{c, namespace}
}

func (c *FakeHarvesterV1alpha1) VirtualMachineTemplateVersions(namespace string) v1alpha1.VirtualMachineTemplateVersionInterface {
	return &FakeVirtualMachineTemplateVersions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHarvesterV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
