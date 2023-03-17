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

.PHONY: generate/oapi
generate/oapi: ## generate oapi code
	oapi-codegen -package gen ./api/openapi.yml > adapter/rest/gen/realworld.gen.go
