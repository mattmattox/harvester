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

package fake

import (
	"context"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScheduleVMBackups implements ScheduleVMBackupInterface
type FakeScheduleVMBackups struct {
	Fake *FakeHarvesterhciV1beta1
	ns   string
}

var schedulevmbackupsResource = v1beta1.SchemeGroupVersion.WithResource("schedulevmbackups")

var schedulevmbackupsKind = v1beta1.SchemeGroupVersion.WithKind("ScheduleVMBackup")

// Get takes name of the scheduleVMBackup, and returns the corresponding scheduleVMBackup object, and an error if there is any.
func (c *FakeScheduleVMBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ScheduleVMBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(schedulevmbackupsResource, c.ns, name), &v1beta1.ScheduleVMBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ScheduleVMBackup), err
}

// List takes label and field selectors, and returns the list of ScheduleVMBackups that match those selectors.
func (c *FakeScheduleVMBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ScheduleVMBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(schedulevmbackupsResource, schedulevmbackupsKind, c.ns, opts), &v1beta1.ScheduleVMBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ScheduleVMBackupList{ListMeta: obj.(*v1beta1.ScheduleVMBackupList).ListMeta}
	for _, item := range obj.(*v1beta1.ScheduleVMBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scheduleVMBackups.
func (c *FakeScheduleVMBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(schedulevmbackupsResource, c.ns, opts))

}

// Create takes the representation of a scheduleVMBackup and creates it.  Returns the server's representation of the scheduleVMBackup, and an error, if there is any.
func (c *FakeScheduleVMBackups) Create(ctx context.Context, scheduleVMBackup *v1beta1.ScheduleVMBackup, opts v1.CreateOptions) (result *v1beta1.ScheduleVMBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(schedulevmbackupsResource, c.ns, scheduleVMBackup), &v1beta1.ScheduleVMBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ScheduleVMBackup), err
}

// Update takes the representation of a scheduleVMBackup and updates it. Returns the server's representation of the scheduleVMBackup, and an error, if there is any.
func (c *FakeScheduleVMBackups) Update(ctx context.Context, scheduleVMBackup *v1beta1.ScheduleVMBackup, opts v1.UpdateOptions) (result *v1beta1.ScheduleVMBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(schedulevmbackupsResource, c.ns, scheduleVMBackup), &v1beta1.ScheduleVMBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ScheduleVMBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScheduleVMBackups) UpdateStatus(ctx context.Context, scheduleVMBackup *v1beta1.ScheduleVMBackup, opts v1.UpdateOptions) (*v1beta1.ScheduleVMBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(schedulevmbackupsResource, "status", c.ns, scheduleVMBackup), &v1beta1.ScheduleVMBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ScheduleVMBackup), err
}

// Delete takes name of the scheduleVMBackup and deletes it. Returns an error if one occurs.
func (c *FakeScheduleVMBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(schedulevmbackupsResource, c.ns, name, opts), &v1beta1.ScheduleVMBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScheduleVMBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(schedulevmbackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ScheduleVMBackupList{})
	return err
}

// Patch applies the patch and returns the patched scheduleVMBackup.
func (c *FakeScheduleVMBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ScheduleVMBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(schedulevmbackupsResource, c.ns, name, pt, data, subresources...), &v1beta1.ScheduleVMBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ScheduleVMBackup), err
}
