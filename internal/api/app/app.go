package app

import (
	"booking/internal/api"
	"booking/internal/pkg/logger"

	"go.uber.org/fx"
)

const Name = "booking_api"

func Provide(conf *api.Config, lg *logger.Logger) fx.Option {
	return fx.Options(
		fx.StartTimeout(conf.StartTimeout),
		fx.StopTimeout(conf.StopTimeout),
		fx.Provide(
			func() *api.Config { return conf },
			func() *logger.Logger { return lg },
		),
		UseCaseModule(),
		WebserverModule(),
		fx.Invoke(func(lg *logger.Logger) {
			lg.LogInfo("Booking service started!")
		}),
	)
}

func Recover(lg *logger.Logger) {
	if err := recover(); err != nil {
		lg.LogFatalf("app recover error: %s", err)
	}
}
