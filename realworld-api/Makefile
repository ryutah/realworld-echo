export
BIN := $(CURDIR)/.bin
PATH := $(abspath $(BIN)):$(PATH)
GOBIN := $(abspath $(BIN))

.PHONY: help
help: ## Prints help for targets with comments
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## initialize projects
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4
	go install github.com/vektra/mockery/v2@v2.32.0
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.16.2
	pip install sqlfluff==2.1.2

.PHONY: generate
generate: generate/oapi generate/mock generate/go ## run all generate task

.PHONY: generate/oapi
generate/oapi: ## generate oapi code
	oapi-codegen -config ./server.cfg.yaml ../docs/api/openapi.yml
	oapi-codegen -config ./types.cfg.yaml ../docs/api/openapi.yml

.PHONY: generate/go
generate/go: ## run go generate command
	go generate ./...

.PHONY: generate/mock
generate/mock: ## run mock generate
	mockery

.PHONY: generate/domain_diagram
generate/domain_diagram: ## generate domain diagram
	podman container run -it -v ${PWD}:/data --rm plantuml/plantuml:1 -tsvg ./docs/domain/domain.puml

.PHONY: build/cloudrun
build/cloudrun: ## build cloudrun container image
	gcloud builds submit --config ./server/cloudrun/cloudbuild.yaml .

.PHONY: deploy/cloudrun
deploy/cloudrun: ## deploy cloudrun container image
	gcloud run deploy \
	  --image gcr.io/${PROJECT_ID}/realworld-api:latest \
	  --platform managed \
	  --region asia-northeast1 \
	  --no-allow-unauthenticated \
	  realworld-api

.PHONY: db/up
db/up: ## start database server
	docker container run -it --rm -p 5432:5432 \
	  -e POSTGRES_PASSWORD=psql \
	  -e POSTGRES_USER=psql \
	  -e POSTGRES_DB=realworld \
	  postgres:15.3-bullseye

.PHONY: migrates/create
migrates/create: ## create new migrate file
	migrate create -dir ./resources/migrations -ext sql -seq -tz Asia/Tokyo ${MIGRATE_NAME}

.PHONY: migrates/up
migrates/up: ## run migrate up
	migrate -source \
	  "file://${PWD}/resources/migrations" \
	  -database 'postgres://psql:psql@localhost:5432/realworld?sslmode=disable' \
	  up

.PHONY: migrates/down
migrates/down: ## run migrate down
	migrate -source \
	  "file://${PWD}/resources/migrations" \
	  -database 'postgres://psql:psql@localhost:5432/realworld?sslmode=disable' \
	  down

.PHONY: format
format: ## format codes
	sqlfluff format  --dialect postgres ./resources/migrations/*.sql

.PHONY: test
test: ## run unit test
	go test -shuffle on -cover -coverprofile cover.out ./...

.PHONY: test/view_coverage
test/view_coverage: ## view coverage report
	go tool cover -html=cover.out -o cover.html
