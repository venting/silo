package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/infinityworks/go-common/logger"
	"github.com/infinityworks/go-common/router"
	"github.com/sirupsen/logrus"
	"github.com/venting/silo/agent"
	"github.com/venting/silo/config"
	siloHttp "github.com/venting/silo/http"
	"github.com/venting/silo/metrics"
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
	workQueue = agent.StartAgent()
	setupHTTP()
}

// setupHTTP will spin up a golang http server to handle the basic REST interface of the silo agent
func setupHTTP() error {

	log.WithFields(structs.Map(applicationCfg)).Info("Starting Silo Agent")

	binding := fmt.Sprintf(":%s", applicationCfg.ListenPort())

	h := siloHttp.Handler{
		Config: applicationCfg,
		Queue:  workQueue,
		Logger: log,
	}

	metrics.Init()

	r := router.NewRouter(log, h.CreateRoutes())

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("This request panicked. Gracefully killing the server: %s", r)
		}
	}()

	return http.ListenAndServe(binding, r)
}
