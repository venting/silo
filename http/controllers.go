package http

import "net/http"

// ListContainers will query the local docker socket, and return a list of containers that are currently running
// it will also attempt to represent other information (like uptime and history) in an ordered manner
func (h Handler) ListContainers(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	return http.StatusOK, []byte(""), nil
}

// SetConfig method will accept an incoming configuration change, it will take the body of the request, and use that
// as a new source for container configuration
func (h Handler) SetConfig(w http.ResponseWriter, r *http.Request) (status int, body []byte, err error) {

	return http.StatusNoContent, []byte(""), nil
}
