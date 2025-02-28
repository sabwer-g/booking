package api

import (
	"time"

	"booking/internal/pkg/logger"
	"booking/internal/pkg/webserver"
	"github.com/kelseyhightower/envconfig"
)

const CoreEnvironmentPrefix = "BOOKING"

type Config struct {
	Env            string           `envconfig:"env"`
	Debug          bool             `envconfig:"debug"`
	ProfilerEnable bool             `envconfig:"pprof"`
	StartTimeout   time.Duration    `envconfig:"start_timeout" default:"20s"`
	StopTimeout    time.Duration    `envconfig:"stop_timeout" default:"60s"`
	Logger         logger.Config    `envconfig:"logger"`
	WebServer      webserver.Config `envconfig:"webserver"`
}

func Usage() error {
	return envconfig.Usage(CoreEnvironmentPrefix, &Config{})
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process(CoreEnvironmentPrefix, cfg); err != nil {
		return nil, err
	}

	if cfg.Debug {
		cfg.Logger.Level = "debug"
		cfg.Logger.Debug = true
	}

	return cfg, nil
}
