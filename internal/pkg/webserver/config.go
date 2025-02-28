package webserver

import (
	"time"
)

type Config struct {
	Host            string        `envconfig:"host"`
	ReadTimeout     time.Duration `envconfig:"read_timeout" default:"30s"`
	WriteTimeout    time.Duration `envconfig:"write_timeout" default:"30s"`
	ShutdownTimeout time.Duration `envconfig:"shutdown_timeout" default:"10s"`
	Port            int           `envconfig:"port" default:"8080"`
	IsAuthEnabled   bool          `envconfig:"auth_enabled" default:"false"`
}
