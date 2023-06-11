package main

import (
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"

	"cloud.google.com/go/profiler"
)

func main() {
	if err := profiler.Start(profiler.Config{
		Service:        "my_service",
		ServiceVersion: "v1",
	}); err != nil {
		panic(err)
	}

	e := di.InitializeAppEngineRestExecuter(
		xtrace.ProjectID(config.GetConfig().ProjectID),
		xerrorreport.Service("my_service"),
		xerrorreport.Version("v1"),
	)
	e.Start()
}
