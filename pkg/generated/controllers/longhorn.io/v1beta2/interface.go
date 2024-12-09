/*
Copyright 2025 Rancher Labs, Inc.

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

package v1beta2

import (
	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1beta2.AddToScheme)
}

type Interface interface {
	BackingImage() BackingImageController
	BackingImageDataSource() BackingImageDataSourceController
	Backup() BackupController
	BackupBackingImage() BackupBackingImageController
	Engine() EngineController
	Node() NodeController
	Replica() ReplicaController
	Setting() SettingController
	Snapshot() SnapshotController
	Volume() VolumeController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) BackingImage() BackingImageController {
	return generic.NewController[*v1beta2.BackingImage, *v1beta2.BackingImageList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "BackingImage"}, "backingimages", true, v.controllerFactory)
}

func (v *version) BackingImageDataSource() BackingImageDataSourceController {
	return generic.NewController[*v1beta2.BackingImageDataSource, *v1beta2.BackingImageDataSourceList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "BackingImageDataSource"}, "backingimagedatasources", true, v.controllerFactory)
}

func (v *version) Backup() BackupController {
	return generic.NewController[*v1beta2.Backup, *v1beta2.BackupList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Backup"}, "backups", true, v.controllerFactory)
}

func (v *version) BackupBackingImage() BackupBackingImageController {
	return generic.NewController[*v1beta2.BackupBackingImage, *v1beta2.BackupBackingImageList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "BackupBackingImage"}, "backupbackingimages", true, v.controllerFactory)
}

func (v *version) Engine() EngineController {
	return generic.NewController[*v1beta2.Engine, *v1beta2.EngineList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Engine"}, "engines", true, v.controllerFactory)
}

func (v *version) Node() NodeController {
	return generic.NewController[*v1beta2.Node, *v1beta2.NodeList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Node"}, "nodes", true, v.controllerFactory)
}

func (v *version) Replica() ReplicaController {
	return generic.NewController[*v1beta2.Replica, *v1beta2.ReplicaList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Replica"}, "replicas", true, v.controllerFactory)
}

func (v *version) Setting() SettingController {
	return generic.NewController[*v1beta2.Setting, *v1beta2.SettingList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Setting"}, "settings", true, v.controllerFactory)
}

func (v *version) Snapshot() SnapshotController {
	return generic.NewController[*v1beta2.Snapshot, *v1beta2.SnapshotList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Snapshot"}, "snapshots", true, v.controllerFactory)
}

func (v *version) Volume() VolumeController {
	return generic.NewController[*v1beta2.Volume, *v1beta2.VolumeList](schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "Volume"}, "volumes", true, v.controllerFactory)
}
