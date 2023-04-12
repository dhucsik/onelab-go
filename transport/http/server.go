package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"practice/config"
	"practice/transport/http/handler"
	middleware2 "practice/transport/http/middleware"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	cfg     *config.Config
	handler *handler.Manager
	App     *echo.Echo
	m       *middleware2.JWTAuth
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	m := middleware2.NewJWTAuth(cfg)
	return &Server{
		cfg:     cfg,
		handler: handler,
		m:       m,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()
	go func() {
		if err := s.App.Start(fmt.Sprintf(":%s", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%v\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.App.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%v", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	return e
}
