package main

import (
	"log"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
)

func main() {
	di.InjectLocal(
		di.InjectParam{
			DBConfig: sqlc.DBConfig{
				ConnectionName: config.GetConfig().DBConnection,
			},
		},
		func(e *rest.Extcuter) {
			log.Println("start app")
		},
	).Run()
}
