package main

import (
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"

	"cloud.google.com/go/profiler"
)

func main() {
	if err := profiler.Start(profiler.Config{
		Service:        "my_service",
		ServiceVersion: "v1",
	}); err != nil {
		panic(err)
	}

	e := di.InitializeAppEngineRestExecuter(config.GetConfig().ProjectID)
	e.Start()
}
