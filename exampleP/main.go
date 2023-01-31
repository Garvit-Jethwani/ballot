package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/libpod/v2/pkg/bindings"
	"github.com/containers/libpod/v2/pkg/bindings/images"
	"github.com/containers/libpod/v2/pkg/domain/entities"
)

func main() {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Get Podman socket location
	sock_dir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sock_dir + "/podman/podman.sock"
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Pulling Busybox image...")
	_, err = images.Pull(connText, "docker.io/busybox", entities.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Pull Fedora image (Sample 2)
	rawImage := "registry.fedoraproject.org/fedora:latest"
	fmt.Println("Pulling Fedora image...")
	_, err = images.Pull(connText, rawImage, entities.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
