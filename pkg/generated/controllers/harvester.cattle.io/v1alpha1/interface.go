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

package v1alpha1

import (
	v1alpha1 "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1alpha1.AddToScheme)
}

type Interface interface {
	KeyPair() KeyPairController
	Setting() SettingController
	User() UserController
	VirtualMachineImage() VirtualMachineImageController
	VirtualMachineTemplate() VirtualMachineTemplateController
	VirtualMachineTemplateVersion() VirtualMachineTemplateVersionController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) KeyPair() KeyPairController {
	return NewKeyPairController(schema.GroupVersionKind{Group: "harvester.cattle.io", Version: "v1alpha1", Kind: "KeyPair"}, "keypairs", true, c.controllerFactory)
}
func (c *version) Setting() SettingController {
	return NewSettingController(schema.GroupVersionKind{Group: "harvester.cattle.io", Version: "v1alpha1", Kind: "Setting"}, "settings", false, c.controllerFactory)
}
func (c *version) User() UserController {
	return NewUserController(schema.GroupVersionKind{Group: "harvester.cattle.io", Version: "v1alpha1", Kind: "User"}, "users", false, c.controllerFactory)
}
func (c *version) VirtualMachineImage() VirtualMachineImageController {
	return NewVirtualMachineImageController(schema.GroupVersionKind{Group: "harvester.cattle.io", Version: "v1alpha1", Kind: "VirtualMachineImage"}, "virtualmachineimages", true, c.controllerFactory)
}
func (c *version) VirtualMachineTemplate() VirtualMachineTemplateController {
	return NewVirtualMachineTemplateController(schema.GroupVersionKind{Group: "harvester.cattle.io", Version: "v1alpha1", Kind: "VirtualMachineTemplate"}, "virtualmachinetemplates", true, c.controllerFactory)
}
func (c *version) VirtualMachineTemplateVersion() VirtualMachineTemplateVersionController {
	return NewVirtualMachineTemplateVersionController(schema.GroupVersionKind{Group: "harvester.cattle.io", Version: "v1alpha1", Kind: "VirtualMachineTemplateVersion"}, "virtualmachinetemplateversions", true, c.controllerFactory)
}
