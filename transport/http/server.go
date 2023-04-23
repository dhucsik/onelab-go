package http

import (
	"context"
	"fmt"
	"net/http"
	"practice/config"
	"practice/logging"
	"practice/transport/http/handler"
	middleware2 "practice/transport/http/middleware"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
	m       *middleware2.JWTAuth
	logger  *logging.Logger
}

func NewServer(cfg *config.Config, handler *handler.Manager, logger *logging.Logger) *Server {
	m := middleware2.NewJWTAuth(cfg)
	return &Server{
		cfg:     cfg,
		handler: handler,
		m:       m,
		logger:  logger,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()
	s.logger.Infof("set routes\n")

	go func() {
		if err := s.App.Start(fmt.Sprintf(":%s", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("listen:%v\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		s.logger.Fatalf("server Shutdown Failed:%v", err)
	}
	s.logger.Infof("server exited properly\n")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	s.logger.Infof("Set cors configuration\n")

	return e
}
