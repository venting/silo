package docker

import (
	"fmt"

	gdc "github.com/fsouza/go-dockerclient"
	"github.com/venting/silo/config"
)

// ListRunningContainers returns the effective output of a `docker ps` for running containers only.
func ListRunningContainers(cfg config.Config) ([]gdc.APIContainers, error) {

	cl := []gdc.APIContainers{}

	client, err := gdc.NewClient(cfg.Socket)
	if err != nil {
		return nil, fmt.Errorf("Error creating Docker client: %v", err)

	}

	cs, err := client.ListContainers(gdc.ListContainersOptions{All: false})
	if err != nil {
		return cs, fmt.Errorf("Error listing containers: %v", err)
	}

	for _, c := range cs {
		cl = append(cl, c)
	}

	return cl, nil

}
