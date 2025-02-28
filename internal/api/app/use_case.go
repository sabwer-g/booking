package app

import (
	"booking/internal/api/service/orders/orders_create"
	"booking/internal/pkg/logger"

	"go.uber.org/fx"
)

func UseCaseModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(log *logger.Logger) *orders_create.UseCase {
				return orders_create.NewUseCase(log)
			},
		),
	)
}
