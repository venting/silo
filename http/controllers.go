package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/infinityworks/go-common/router"
	"github.com/venting/silo/docker"
)

// ListContainers will query the local docker socket, and return a list of containers that are currently running
// it will also attempt to represent other information (like uptime and history) in an ordered manner
func (h Handler) ListContainers(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	cs, err := docker.ListContainers()

	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}

	return router.MarshalBody(cs)
}

// RestartContainer will query the local docker socket, and return a list of containers that are currently running
// it will also attempt to represent other information (like uptime and history) in an ordered manner
func (h Handler) RestartContainer(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	vars := mux.Vars(r)
	id := vars["id"]

	err = docker.RestartContainer(id, h.Config.DockerTimeout)

	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}

	return http.StatusOK, []byte(""), nil
}

// SetConfig method will accept an incoming configuration change, it will take the body of the request, and use that
// as a new source for container configuration
func (h Handler) SetConfig(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	return http.StatusNoContent, []byte(""), nil
}

// GetHealth returns the overall state of the node, this comprises of the status of the containers running
// and the state of the silo-agent.
func (h Handler) GetHealth(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	cs, err := docker.ListContainers()

	if err != nil {
		return http.StatusInternalServerError, []byte(""), err
	}

	// nrc is used to capture any containers that aren't in a running state
	var nrc = []string{}

	for _, c := range cs {
		if c.State != "running" {
			nrc = append(nrc, c.ID)
		}

	}

	if len(nrc) > 0 {
		return http.StatusFailedDependency, []byte(""), fmt.Errorf("Containers not running: %s", nrc)
	}

	return http.StatusOK, []byte(""), nil
}
