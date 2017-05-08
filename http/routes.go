package http

import (
	"github.com/infinityworksltd/go-common/router"
	"github.com/venting/silo/config"
)

// Handler struct defines structure that all routes will hang off
type Handler struct {
	Config config.Config
}

// CreateRoutes will return a set of mux router compatible routes to serve the app with
func (h Handler) CreateRoutes() router.Routes {
	return router.Routes{

		// List contaienrs exposes the containers currenty running on the stack
		router.Route{
			Name:        "ListContainers",
			Method:      "GET",
			Pattern:     "/containers",
			HandlerFunc: h.ListContainers,
		},

		// SetConfig route exposes the basic interface that allows configuration to be updated
		router.Route{
			Name:        "SetConfig",
			Method:      "POST",
			Pattern:     "/config",
			HandlerFunc: h.SetConfig,
		},
	}
}
