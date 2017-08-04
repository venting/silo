package http

import (
	"github.com/infinityworks/go-common/router"
	"github.com/venting/silo/config"
)

// Handler struct defines structure that all routes will hang off
type Handler struct {
	Config config.Config
}

// CreateRoutes will return a set of mux router compatible routes to serve the app with
func (h Handler) CreateRoutes() router.Routes {
	return router.Routes{

		// List containers exposes the containers currenty running on the stack
		router.Route{
			Name:        "ListContainers",
			Method:      "GET",
			Pattern:     "/containers",
			HandlerFunc: h.ListContainers,
		},

		// Restarts the container with the specified ID
		router.Route{
			Name:        "RestartContainer",
			Method:      "POST",
			Pattern:     "/container/{id}/restart",
			HandlerFunc: h.RestartContainer,
		},

		// SetConfig route exposes the basic interface that allows configuration to be updated
		router.Route{
			Name:        "SetConfig",
			Method:      "POST",
			Pattern:     "/config",
			HandlerFunc: h.SetConfig,
		},

		// GetHealth Exposes basic health state of the containers and silo-agent
		router.Route{
			Name:        "GetHealth",
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: h.GetHealth,
		},
	}
}
