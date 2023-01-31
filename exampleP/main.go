package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/libpod/v2/pkg/bindings"
	"github.com/containers/libpod/v2/pkg/bindings/containers"
)

type ListContainer struct {
	// Container command
	Command []string
	// Container creation time
	// Human readable container creation time.
	CreatedAt string
	// If container has exited/stopped
	Exited bool
	// Time container exited
	ExitedAt int64
	// If container has exited, the return code from the command
	ExitCode int32
	// The unique identifier for the container
	ID string `json:"Id"`
	// Container image
	Image string
	// Container image ID
	ImageID string
	// If this container is a Pod infra container
	IsInfra bool
	// Labels for container
	Labels map[string]string
	// User volume mounts
	Mounts []string
	// The names assigned to the container
	Names []string
	// Namespaces the container belongs to.  Requires the
	// namespace boolean to be true
	// The process id of the container
	Pid int
	// If the container is part of Pod, the Pod ID. Requires the pod
	// boolean to be set
	Pod string
	// If the container is part of Pod, the Pod name. Requires the pod
	// boolean to be set
	PodName string
	// Port mappings
	Ports []PortMapping
	// Size of the container rootfs.  Requires the size boolean to be true
	// Time when container started
	StartedAt int64
	// State of container
	State string
}
type PortMapping struct {
	// HostPort is the port number on the host.
	HostPort int32 `json:"hostPort"`
	// ContainerPort is the port number inside the sandbox.
	ContainerPort int32 `json:"containerPort"`
	// Protocol is the protocol of the port mapping.
	Protocol string `json:"protocol"`
	// HostIP is the host ip to use.
	HostIP string `json:"hostIP"`
}

func main() {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Get Podman socket location
	socket := "unix:///run/user/1000/podman/podman.sock"
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Container list
	var latestContainers = 1
	var containerLatestList []ListContainer
	containerLatestList, err = containers.List(connText, nil, nil, &latestContainers, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Latest container is %+v\n", containerLatestList)
}
