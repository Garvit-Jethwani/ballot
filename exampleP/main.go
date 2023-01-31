package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/libpod/v2/pkg/bindings"
	"github.com/containers/libpod/v2/pkg/bindings/containers"
	"github.com/containers/libpod/v2/pkg/bindings/images"
	"github.com/containers/libpod/v2/pkg/domain/entities"
)

func main() {
	fmt.Println("Welcome to the Podman Go bindings tutorial")

	// Get Podman socket location
	socket := "unix:///run/user/1000/podman/podman.sock"
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Pulling Busybox image...")
	_, err = images.Pull(connText, "docker.io/zbio/ballot", entities.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Container list
	var latestContainers = 1
	containerLatestList, err := containers.List(connText, nil, nil, &latestContainers, nil, nil, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Latest container is %s\n", containerLatestList[0].Names[0])
}
