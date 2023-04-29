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
	go install github.com/google/wire/cmd/wire@v0.5.0
	go install github.com/golang/mock/mockgen@v1.6.0


.PHONY: generate/oapi
generate/oapi: ## generate oapi code
	oapi-codegen -package gen ./docs/api/openapi.yml > api/rest/gen/realworld.gen.go

.PHONY: generate/wire
generate/wire: ## generate wire di code
	wire ./di

.PHONY: generate/go
generate/go: ## run go generate command
	go generate ./...

.PHONY: generate/domain_diagram
generate/domain_diagram: ## generate domain diagram
	podman container run -it -v ${PWD}:/data --rm plantuml/plantuml:1 -tsvg ./docs/domain/domain.puml
