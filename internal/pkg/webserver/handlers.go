package webserver

import (
	"net/http"

	"booking/internal/pkg/logger"
)

func NewRecoveryHandler(l *logger.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &RecoveryHandler{logger: l, handler: h}
	}
}

type RecoveryHandler struct {
	logger  *logger.Logger
	handler http.Handler
}

func (h RecoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			h.logger.LogErrorf("webserver recover error: %s", err)
		}
	}()

	h.handler.ServeHTTP(w, req)
}
