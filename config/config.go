package config

import (
	cfg "github.com/infinityworks/go-common/config"
)

// Config - Defines application configuration
type Config struct {
	*cfg.BaseConfig
	Socket string
}

// Init - Initialises Config struct with safe defaults if not present
func Init() Config {
	ac := cfg.Init()

	appConfig := Config{
		&ac,
		// DOCKER_HOST is the default environment variable used by docker for specifying the socket
		cfg.GetEnv("DOCKER_HOST", "unix:///var/run/docker.sock"),
	}

	return appConfig
}
