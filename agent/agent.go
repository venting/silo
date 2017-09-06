package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/venting/silo/retrieveconfig"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/project/options"
	"github.com/sirupsen/logrus"
)

// Project that this silo-agent is working with
var Project project.APIProject

// ServicesHaveBeenStarted Declares if any services have been started on this silo agent
var ServicesHaveBeenStarted = false

// SiloProjectName is the const name of the "project" we will run through libcompose
const SiloProjectName string = "silo-client-services"

// WorkQueue is a buffered channel that we can send work requests on
var WorkQueue = make(chan Work, 100)

// WorkerQueue defines the interface for workers to read from, a channel of work channels
var WorkerQueue chan chan Work

// Work struct defines the basic format for all work processed by the agent
// Data is an encoded JSON raw message and is decoded based upon the given name
type Work struct {
	Name  string
	Data  json.RawMessage
	Delay time.Duration
}

// Worker struct defines and holds the state of a given worker
type Worker struct {
	ID          int
	Work        chan Work
	WorkerQueue chan chan Work
	QuitChan    chan bool
}

// ConfigUpdate struct contains data to update the running contianers on the system
type ConfigUpdate struct {
	ConfigType string
	Config     []byte
}

// StartAgent will spin up the docker stack specified by the given compose file
func StartAgent(dockerComposePath string, logger *logrus.Logger) (queue chan Work) {
	// First, initialize the channel we are going to but the workers' work channels into.
	WorkerQueue = make(chan chan Work, 1)

	worker := NewWorker(1, WorkerQueue)
	worker.Start(logger)

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				logger.Debug("Received work requeust")
				go func() {
					worker := <-WorkerQueue

					logger.Debugf("Dispatching work request with type: %s", work.Name)
					worker <- work
				}()
			}
		}
	}()

	if len(dockerComposePath) != 0 {
		logger.Infof("Starting initial worker with given docker-compose")
		cfg := ConfigUpdate{
			Config: []byte(dockerComposePath),
		}

		configJSON, err := json.Marshal(cfg)

		if err != nil {
			logger.Fatalf("Could not convert config into a JSON object, so could not start services. Error: %s", err)
			return WorkQueue
		}
		// Setup the initial silo server
		work := Work{Name: "config", Data: configJSON}

		WorkQueue <- work
	} else {
		logger.Infof("No docker-compose config given. Starting silo agent with no other services..")
	}

	return WorkQueue
}

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan Work) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan Work),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

// Start begins the goroutine infinite for
func (w *Worker) Start(logger *logrus.Logger) {
	logger.Infof("Starting worker..")
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request.

				logger.Infof("Worker %d has recieved a %s job", w.ID, work.Name)

				switch work.Name {
				case "config":
					config := ConfigUpdate{}
					err := json.Unmarshal(work.Data, &config)
					if err != nil {
						logger.Errorf("Could not parse updated config. Error: %s", err)
						break
					}

					cBytes, err := retrieveconfig.RetrieveConfig(fmt.Sprintf("%s", config.Config))

					if len(cBytes) == 0 && err == nil {
						// Not a URL or a file path, maybe it's just straight docker compose..
					}

					if err != nil {
						logger.Errorf("Could not deconstruct configuration, error given: %s", err)
						break
					}

					// Write the configuration bytes back to the config struct
					config.Config = cBytes

					err = KillAgent(logger)

					if err != nil {
						err = errors.Wrap(err, "Could not kill old project")
						logger.Error(err)
					}

					composeConfig := make([][]byte, 1)
					composeConfig[0] = config.Config
					Project, err = docker.NewProject(
						&ctx.Context{
							Context: project.Context{
								ComposeBytes: composeConfig,
								ProjectName:  SiloProjectName,
							},
						},
						nil,
					)

					if err != nil {
						logger.Errorf("Could not create the new set of services. Got the following error back from libcompose: %s. Silo is now unhealthy", err)
						break
					}

					err = Project.Up(context.Background(), options.Up{})

					ServicesHaveBeenStarted = true

					if err != nil {
						logger.Errorf("Could not start services. Got the following error back from libcompose: %s. Silo is unhealthy", err)
					}

					logger.Infof("Services started")
				}
			}
		}
	}()
}

// KillAgent kills the current project
func KillAgent(logger *logrus.Logger) error {
	if ServicesHaveBeenStarted {
		logger.Info("Taking down services..")

		err := Project.Down(
			context.Background(),
			options.Down{
				// We want to remove all local images from the machine, to stop it getting full of old images
				RemoveImages: options.ImageType("local"),
				// We want to remove orphaned containers when we kill the services
				RemoveOrphans: true,
				// We want to rm the old volumes to get back the disk space, as the old project is dead to us
				RemoveVolume: true,
			},
		)

		if err != nil {
			errors.Wrap(err, "Could not take down project")
		}

		return err
	}

	return nil
}
