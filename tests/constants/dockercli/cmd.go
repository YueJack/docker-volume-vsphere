// Copyright 2017 VMware, Inc. All Rights Reserved.
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

// A home to hold test constants related with docker cli.

package dockercli

const (
	docker    = "docker "
	dockerVol = docker + "volume "

	// ListVolumes to list down docker volumes
	ListVolumes = dockerVol + "ls "

	// InspectVolume to grab volume properties
	InspectVolume = dockerVol + "inspect "

	// CreateVolume create a volume with vsphere driver
	CreateVolume = dockerVol + " create --driver=vsphere "

	// RemoveVolume constant refers delete volume command
	RemoveVolume = dockerVol + "rm "

	// KillDocker kill docker
	KillDocker = "pkill -9 dockerd "

	// RunContainer create and run a container
	RunContainer = docker + "run "

	// RemoveContainer remove the container
	RemoveContainer = docker + "rm -f "
)
