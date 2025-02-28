package orders

import (
	"net/http"
	"time"

	"booking/internal/api/service/orders/orders_create"
	"booking/internal/pkg/logger"
	"booking/internal/pkg/webserver"

	"github.com/gorilla/mux"
)

type Controller struct {
	create *orders_create.UseCase
	log    *logger.Logger
}

func NewController(
	create *orders_create.UseCase,
	log *logger.Logger,
) *Controller {
	return &Controller{
		create: create,
		log:    log,
	}
}

func (c *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/v1/orders/create", func(writer http.ResponseWriter, request *http.Request) {
		webserver.Request(writer, request, c.log, 15*time.Second, c.create.Do)
	}).Methods(http.MethodPost)
}
