package docker

import (
	"fmt"

	gdc "github.com/fsouza/go-dockerclient"
)

// ListContainers returns the effective output of a `docker ps -a`
func ListContainers(socket string) ([]gdc.APIContainers, error) {

	client, err := gdc.NewClient(socket)
	if err != nil {
		return nil, fmt.Errorf("Error creating Docker client: %v", err)

	}

	cs, err := client.ListContainers(gdc.ListContainersOptions{All: true})
	if err != nil {
		return cs, fmt.Errorf("Error listing containers: %v", err)
	}

	return cs, nil

}
