package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/middleware"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"go.uber.org/multierr"
)

type Extcuter struct {
	traceInitializer xtrace.Initializer
	server           *Server
}

func NewExecuter(server *Server, traceInitializer xtrace.Initializer) *Extcuter {
	return &Extcuter{
		server:           server,
		traceInitializer: traceInitializer,
	}
}

func (e *Extcuter) Start() {
	ec := echo.New()
	ec.Use(middleware.WithLogger)

	gen.RegisterHandlersWithBaseURL(ec, e.server, "/api")

	if err := e.startServerwithGracefulShutdown(fmt.Sprintf(":%s", config.GetConfig().Port), ec, func() {
		log.Println("Server shutdown")
	}); err != nil {
		log.Fatal(err)
	}
}

func (e *Extcuter) startServerwithGracefulShutdown(addr string, h http.Handler, cleanFunc func()) error {
	traceHandler, traceFinish, err := e.traceInitializer.HandlerWithTracing(h)
	if err != nil {
		return fmt.Errorf("failed to generate trace: %w", err)
	}

	srv := &http.Server{
		Handler:           traceHandler,
		Addr:              addr,
		ReadHeaderTimeout: 60 * time.Second,
	}

	errChan := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- fmt.Errorf("failed to listen server: %w ", err)
		}
	}()

	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
		defer stop()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if cleanFunc != nil {
			cleanFunc()
		}

		if e := srv.Shutdown(ctx); e != nil {
			errChan <- fmt.Errorf("failed to shutdown server: %w", err)
		}
		if err := traceFinish(ctx); err != nil {
			errChan <- fmt.Errorf("failed to finish tracing: %w", err)
		}
		close(errChan)
	}()

	for e := range errChan {
		err = multierr.Append(err, e)
	}
	return err
}
