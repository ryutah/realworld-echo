package main

import (
	_ "github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
)

func main() {
	e := di.InitializeLocalRestExecuter(
		xerrorreport.Service("my_service"),
		xerrorreport.Version("v1"),
	)
	e.Start()
}
