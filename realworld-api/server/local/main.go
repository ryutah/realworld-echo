package main

import (
	_ "github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"
)

func main() {
	e := di.InitializeLocalRestExecuter()
	e.Start()
}
