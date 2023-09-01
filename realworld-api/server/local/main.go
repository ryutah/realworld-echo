package main

import (
	"log"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	_ "github.com/ryutah/realworld-echo/realworld-api/config"
	"github.com/ryutah/realworld-echo/realworld-api/di"
	"github.com/ryutah/realworld-echo/realworld-api/infrastructure/psql/sqlc"
)

func main() {
	di.InjectLocal(
		di.InjectParam{
			DBConfig: sqlc.DBConfig{
				ConnectionName: "postgresql://psql:psql@127.0.0.1:5432/realworld?sslmode=disable",
			},
		},
		func(e *rest.Extcuter) {
			log.Println("start app")
		},
	).Run()
}
