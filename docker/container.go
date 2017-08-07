package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/venting/silo/metrics"
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

	// clear the existing gauge of running containers
	metrics.ContainersTotal.Reset()

	// increment the gauge metric so we capture the container.state dimension in the simplest way
	for _, container := range containers {
		metrics.ContainersTotal.WithLabelValues(container.State).Inc()
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
		metrics.ContainersTotal.WithLabelValues("container_restart", "failed").Inc()
		return errors.Wrapf(err, "Error restarting container %s", id)
	}

	metrics.ContainersTotal.WithLabelValues("container_restart", "success").Inc()

	return nil
}
