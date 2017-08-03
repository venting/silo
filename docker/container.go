package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListContainers returns a list of containers on the local system
func ListContainers() ([]types.Container, error) {

	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("Error creating Docker client: %v", err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return containers, fmt.Errorf("Error listing containers: %v", err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return containers, nil
}

// RestartContainer issues a restart command to docker for the specified container
func RestartContainer(id string, timeout *time.Duration) error {

	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("Error creating Docker client: %v", err)
	}

	err = cli.ContainerRestart(context.Background(), id, timeout)

	if err != nil {
		return fmt.Errorf("Error restarting container %s, error: %v", id, err)
	}

	return nil
}
