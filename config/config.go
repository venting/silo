package config

import (
	"os"
	"strconv"
	"time"

	cfg "github.com/infinityworks/go-common/config"
)

// Config - Defines application configuration
type Config struct {
	*cfg.BaseConfig
	DockerTimeout *time.Duration
}

// Init - Initialises Config struct with safe defaults if not present
func Init() Config {
	ac := cfg.Init()

	var timeout time.Duration = 5

	if os.Getenv("DOCKER_TIMEOUT") != "" {
		envTimeout, err := strconv.Atoi(os.Getenv("DOCKER_TIMEOUT"))
		if err != nil {
			println("Error converting env variable DOCKER_TIMEOUT to int")
		}

		timeout = time.Duration(envTimeout)
	}

	appConfig := Config{
		&ac,
		&timeout,
	}

	return appConfig
}
