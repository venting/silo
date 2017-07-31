package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/infinityworks/go-common/logger"
	"github.com/infinityworks/go-common/router"
	"github.com/sirupsen/logrus"
	"github.com/venting/silo/config"
	siloHttp "github.com/venting/silo/http"
)

var (
	log            *logrus.Logger
	applicationCfg config.Config
)

func init() {
	applicationCfg = config.Init()
	// Regisster internal logger
	log = logger.Start(applicationCfg)
}

func main() {
	setupHTTP()
}

// setupHTTP will spin up a golang http server to handle the basic REST interface of the silo agent
func setupHTTP() error {
	log.WithFields(structs.Map(applicationCfg)).Info("Starting Silo Agent")

	binding := fmt.Sprintf(":%s", applicationCfg.ListenPort())

	h := siloHttp.Handler{
		Config: applicationCfg,
	}

	r := router.NewRouter(log, h.CreateRoutes())

	return http.ListenAndServe(binding, r)
}
