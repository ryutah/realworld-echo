package main

import (
	_ "github.com/ryutah/realworld-echo/config"
	"github.com/ryutah/realworld-echo/di"
)

func main() {
	e := di.InitializeLocalRestExecuter()
	e.Start()
}
