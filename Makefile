.PHONY: help
help: ## Prints help for targets with comments
	@grep -E '^[/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: db/up
db/up: ## start databse server
	docker compose up -d db migration

.PHONY: server/start
server/start: ## start local server
	docker compose up -d

.PHONY: server/stop
server/stop: ## stop local server
	docker compose down
