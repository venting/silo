package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/infinityworks/go-common/router"

	"github.com/pkg/errors"
	"github.com/venting/silo/agent"
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

	// Get the configType and config form input variables
	r.ParseForm()

	configType, typeExists := r.Form["configType"]
	configData, configExists := r.Form["config"]

	if !typeExists || !configExists {
		return http.StatusNotAcceptable, []byte(""), fmt.Errorf("You must pass a configType and config elements")
	}

	config := agent.ConfigUpdate{
		ConfigType: configType[0],
		Config:     []byte(configData[0]),
	}

	configJSON, err := json.Marshal(config)

	if err != nil {
		return http.StatusInternalServerError, []byte("TITS"), errors.Wrap(err, "Could not convert the config to JSON")
	}

	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
	work := agent.Work{Name: "config", Data: configJSON}

	h.Queue <- work

	return http.StatusOK, []byte(""), nil
}

// KillAgent stops the current docker-compose project and then kills the worker
func (h Handler) KillAgent(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	agent.KillAgent(h.Logger)

	return http.StatusOK, []byte(""), nil
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
