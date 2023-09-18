package rest

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/middleware"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"go.uber.org/fx"
	"go.uber.org/multierr"
)

type Extcuter struct {
	traceInitializer xtrace.Initializer
	server           *Server
}

func NewExecuter(lc fx.Lifecycle, server *Server, traceInitializer xtrace.Initializer) (*Extcuter, error) {
	e := &Extcuter{
		server:           server,
		traceInitializer: traceInitializer,
	}

	ec := echo.New()
	ec.Use(middleware.WithLogger)
	si := gen.NewStrictHandler(e.server, []gen.StrictMiddlewareFunc{})
	gen.RegisterHandlersWithBaseURL(ec, si, "/api")

	traceHandler, traceFinish, err := e.traceInitializer.HandlerWithTracing(ec)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize trace handler: %w", err)
	}

	ls, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GetConfig().Port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	srv := &http.Server{
		Handler:           traceHandler,
		ReadHeaderTimeout: config.GetConfig().RequestTimeOut(),
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Printf("start server: %s", ls.Addr().String())
				if err := srv.Serve(ls); err != nil && err != http.ErrServerClosed {
					log.Fatalf("failed to serve: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("shutdown server ...")

			return multierr.Combine(
				traceFinish(ctx),
				srv.Shutdown(ctx),
			)
		},
	})

	return e, nil
}
