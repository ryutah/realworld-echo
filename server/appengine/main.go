package main

import (
	"github.com/ryutah/realworld-echo/config"
	"github.com/ryutah/realworld-echo/di"
)

func main() {
	e := di.InitializeAppEngineRestExecuter(config.GetConfig().ProjectID)
	e.Start()
}
