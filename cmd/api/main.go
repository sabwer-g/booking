package main

import (
	"net/http"
	"os"

	"booking/internal/api"
	"booking/internal/api/app"
	"booking/internal/pkg/logger"

	"go.uber.org/fx"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		_ = api.Usage()
		return
	}

	conf, err := api.NewConfigFromEnv()
	if err != nil {
		panic(err)
	}

	conf.Logger.AppName = app.Name

	lg := initLogger(conf)

	if conf.ProfilerEnable {
		runProfiler(lg)
	}

	defer app.Recover(lg)
	fx.New(
		app.Provide(conf, lg),
	).Run()
}

func initLogger(conf *api.Config) *logger.Logger {
	lg, err := logger.New(conf.Logger)
	if err != nil {
		panic(err)
	}

	return lg
}

func runProfiler(logger *logger.Logger) {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			logger.LogErrorf("pprof http server can can't start: %s", err)
		}
	}()
}
