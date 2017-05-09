package docker

import (
	"fmt"

	gdc "github.com/fsouza/go-dockerclient"
)

// ListRunningContainers returns the effective output of a `docker ps` for running containers only.
func ListRunningContainers(socket string) ([]gdc.APIContainers, error) {

	client, err := gdc.NewClient(socket)
	if err != nil {
		return nil, fmt.Errorf("Error creating Docker client: %v", err)

	}

	cs, err := client.ListContainers(gdc.ListContainersOptions{All: false})
	if err != nil {
		return cs, fmt.Errorf("Error listing containers: %v", err)
	}

	return cs, nil

}
