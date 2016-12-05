package main

import (
	"fmt"
	"os"

	"github.com/fgrehm/go-dockerpty"
	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:        "ubuntu",
			Cmd:          []string{"/bin/sh"},
			OpenStdin:    true,
			StdinOnce:    true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		},
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    container.ID,
			Force: true,
		})
	}()

	if err = dockerpty.Start(client, container, &docker.HostConfig{}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
