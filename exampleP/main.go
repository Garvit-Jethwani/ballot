package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v3/pkg/bindings"
	"github.com/containers/podman/v3/pkg/bindings/containers"
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
	// Container list
	containerLatestList, err := containers.List(connText, &containers.ListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Latest container is %+v\n", containerLatestList)
}
