package main

import (
	"log"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	_ "github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"
)

func main() {
	di.InjectLocal(func(e *rest.Extcuter) {
		log.Println("start app")
	}).Run()
}
