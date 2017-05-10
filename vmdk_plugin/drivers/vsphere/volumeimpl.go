// Copyright 2016 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vsphere

import (
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/vmware/docker-volume-vsphere/vmdk_plugin/utils/plugin_utils"
)

// VolumeDriver interface used by the refcountedVolume module to handle
// recovery mounts/unmounts.
type VolumeImpl interface {
	Create(volume.Request) volume.Response
	Mount(volume.MountRequest, *plugin_utils.VolumeInfo) volume.Response
	Unmount(volume.UnmountRequest) volume.Response
	Get(volume.Request) volume.Response
	Remove(volume.Request) volume.Response
	Path(volume.Request) volume.Response
	List(volume.Request) ([]*volume.Volume, error)
	GetMountPoint(string) string
	IsMounted(string) bool
	MountVolume(string, string, string, bool, bool) (string, error)
	UnmountVolume(string) error
	GetVolume(string) (map[string]interface{}, error)
}
