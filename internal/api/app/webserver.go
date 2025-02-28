package app

import (
	"booking/internal/api"
	"booking/internal/api/webserver/orders"
	"booking/internal/pkg/logger"
	"booking/internal/pkg/middleware"
	"booking/internal/pkg/webserver"

	"go.uber.org/fx"
)

const (
	serviceUri = "/api/booking"
)

type webServerDeps struct {
	fx.In

	Config *api.Config
	Logger *logger.Logger

	Orders *orders.Controller
}

func WebserverModule() fx.Option {
	return fx.Options(
		fx.Provide(
			orders.NewController,
		),
		fx.Provide(func(
			cfg *api.Config,
			lg *logger.Logger,
			d webServerDeps,
		) *webserver.WebServer {
			optionFns := make([]webserver.OptionFn, 0)

			optionFns = append(optionFns, webserver.MiddlewareForSubRouterOption(
				serviceUri,
				middleware.GetUnmarshalBodyMiddleware(lg),
			))

			optionFns = append(optionFns,
				webserver.ControllerForSubRouterOption(
					serviceUri,
					d.Orders,
				),
			)

			return webserver.NewWebServer(serviceUri, cfg.WebServer, lg, optionFns...)
		}),
		fx.Invoke(
			func(lf fx.Lifecycle, s *webserver.WebServer) {
				lf.Append(fx.Hook{OnStart: s.Start, OnStop: s.Stop})
			},
		),
	)
}
