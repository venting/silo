package main

import (
	"fmt"
	"net/http"
	"os"

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
	workQueue      chan agent.Work
)

func init() {
	applicationCfg = config.Init()
	// Regisster internal logger
	log = logger.Start(applicationCfg)
}

func main() {
	var configFilePath string

	if len(os.Args) == 2 {
		log.Infof("Starting application with given configuration file in path: %s\n", os.Args[1])
		configFilePath = os.Args[1]
	} else {
		log.Infof("Starting application without any configuration. Listening for config..")
	}

	workQueue = agent.StartAgent(configFilePath, log)

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
