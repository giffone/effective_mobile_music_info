package server

import (
	"context"
	"fmt"
	"log"
	"music_info/internal/api"
	"music_info/internal/config"
	"music_info/internal/repo"
	"music_info/internal/service"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	echoLog "github.com/labstack/gommon/log"
)

type Server interface {
	Run(ctx context.Context, cfg *config.Config)
	Stop(ctx context.Context)
}

type server struct {
	router *echo.Echo
	env    *Env
}

func NewServer(env *Env, cfg *config.Config) Server {
	s := server{
		router: echo.New(),
		env:    env,
	}

	// router logging level
	if cfg.Debug {
		s.router.Logger.SetLevel(echoLog.DEBUG)
	}

	// storage
	storage := repo.New(env.pool)

	// service
	svc := service.New(storage)

	// handlers
	hndl := api.New(svc)

	// set middlewares
	s.router.Use(middleware.Logger(), middleware.Recover())

	// register handlers
	g1 := s.router.Group("/api/v1/")
	g1.GET("/info", hndl.GetInfoBy)

	return &s
}

func (s *server) Run(ctx context.Context, cfg *config.Config) {
	ctxSignal, cancelSignal := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// start rest api server
	go func() {
		defer cancelSignal()

		portStr := fmt.Sprintf(":%d", cfg.AppPort)

		if err := s.router.Start(portStr); err != nil && err != http.ErrServerClosed {
			log.Printf("server start error: %s\n", err.Error())
		}
	}()

	// wait system notifiers or cancel func
	<-ctxSignal.Done()
}

func (s *server) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// stop server
	err := s.router.Shutdown(ctx)
	if err != nil {
		log.Printf("rest api server stop error: %s\n", err.Error())
	}

	if err == nil {
		log.Println("server stopped successfully with no error")
	} else {
		log.Println("server stop done")
	}
}
