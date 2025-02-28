package webserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"strconv"

	"booking/internal/pkg/logger"

	"github.com/gorilla/mux"
)

type OptionFn func(s *WebServer)

func RegisterAppSubRouterOption(baseURLs ...string) OptionFn {
	return func(s *WebServer) {
		for _, url := range baseURLs {
			s.applicationRouter[url] = s.sysRouter.PathPrefix(url).Subrouter()
		}
	}
}

func ControllerForSubRouterOption(baseURL string, controllers ...controller) OptionFn {
	return func(s *WebServer) {
		router, ok := s.applicationRouter[baseURL]
		if !ok {
			panic("no route")
		}

		for _, c := range controllers {
			c.RegisterRoutes(router)
		}
	}
}

func MiddlewareForSubRouterOption(baseURL string, middlewares ...mux.MiddlewareFunc) OptionFn {
	return func(s *WebServer) {
		router, ok := s.applicationRouter[baseURL]
		if !ok {
			panic("no route")
		}

		for _, m := range middlewares {
			router.Use(m)
		}
	}
}

type WebServer struct {
	cfg               Config
	logger            *logger.Logger
	server            *http.Server
	sysRouter         *mux.Router
	applicationRouter map[string]*mux.Router
}

func (s *WebServer) Start(_ context.Context) error {
	addr := net.JoinHostPort(s.cfg.Host, strconv.Itoa(s.cfg.Port))
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	go func() {
		s.logger.LogInfo("starting webserver: %v", s.cfg)
		if err := s.server.Serve(ln); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				s.logger.LogInfo("Server closed")
			}
			s.logger.LogErrorf("Server failed: %s", err)
			os.Exit(1)
		}
	}()
	return nil
}

func (s *WebServer) Stop(ctx context.Context) error {
	c, cancel := context.WithTimeout(ctx, s.cfg.ShutdownTimeout)
	defer cancel()

	return s.server.Shutdown(c)
}

type controller interface {
	RegisterRoutes(router *mux.Router)
}

func NewCustomWebServer(
	cfg Config,
	lg *logger.Logger,
	options ...OptionFn,
) *WebServer {
	sysRouter := mux.NewRouter()

	ws := WebServer{
		cfg:               cfg,
		logger:            lg,
		sysRouter:         sysRouter,
		applicationRouter: make(map[string]*mux.Router),
		server: &http.Server{
			Handler:      NewRecoveryHandler(lg)(sysRouter),
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
		},
	}

	for _, fn := range options {
		fn(&ws)
	}

	return &ws
}

func NewWebServer(
	baseURL string,
	cfg Config,
	lg *logger.Logger,
	options ...OptionFn,
) *WebServer {
	opts := make([]OptionFn, 0, len(options)+1)
	opts = append([]OptionFn{RegisterAppSubRouterOption(baseURL)}, options...)
	ws := NewCustomWebServer(cfg, lg, opts...)

	return ws
}
