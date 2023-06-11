package rest

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest/middleware"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xhttp"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
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

	traceHandler, traceFinish, err := e.traceInitializer.HandlerWithTracing(ec)
	if err != nil {
		panic(err)
	}
	if err := xhttp.RunWithGraceful(fmt.Sprintf(":%s", config.GetConfig().Port), traceHandler, func(ctx context.Context) error {
		log.Println("shutdown server ...")
		return traceFinish(ctx)
	}); err != nil {
		log.Fatal(err)
	}
}
