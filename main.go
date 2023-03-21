package main

import (
	"github.com/ryutah/realworld-echo/di"
)

func main() {
	e := di.InitializeRestExecuter("sample_project_id")
	e.Start()
}
