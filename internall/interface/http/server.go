package http

import (
	"context"
	"github.com/amirhosseinmoayedi/Project-template/internall/config"
	"github.com/amirhosseinmoayedi/Project-template/internall/interface/http/custom_middleware"
	"github.com/amirhosseinmoayedi/Project-template/internall/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true

	return &Server{
		e: e,
	}
}

func (s *Server) Serve() {
	s.initiate()

	address := config.Configs.Server.Address()

	go func() {
		if err := s.e.Start(address); err != nil && err != http.ErrServerClosed {
			s.e.Logger.Fatal("shutting down the server")
		}
	}()
	log.Logger.Infof("server started at: %s", address)
}

func (s *Server) initiate() {
	s.e.Use(middleware.Recover())
	s.e.Use(custom_middleware.CorsMiddleware())
	s.e.Use(custom_middleware.LoggerMiddleware())
	s.e.GET("/health-check/", HeartBeat)

	log.Logger.Info("server initiated")
}

func (s *Server) WaitForShutDownSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Fatal(err)
	}

	log.Logger.Info("server shutdown")
}
