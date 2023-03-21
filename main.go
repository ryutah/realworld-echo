package main

import (
	"github.com/ryutah/realworld-echo/api/rest"
	"github.com/ryutah/realworld-echo/di"
)

func main() {
	s := di.InitializeServer()
	rest.Start(s)
}

