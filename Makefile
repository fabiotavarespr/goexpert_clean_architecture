PWD = $(shell pwd -L)
GO_CMD=go
DOCKER_CMD=docker
DOCKER_COMPOSE_CMD=docker-compose
GO_TEST=$(GO_CMD) test
PATH_DOCKER_COMPOSE_FILE=docker-compose.yaml

.PHONY: docker-compose-up docker-compose-down docker-compose-restart

all: help

help: ## Display help screen
	@echo "Usage:"
	@echo "	make [COMMAND]"
	@echo "	make help \n"
	@echo "Commands: \n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

tidy: ## Go mod tidy
	$(GO_CMD) mod tidy

fmt: go-tidy ## Go mod tidy
	$(GO_CMD) fmt ./...

test: go-fmt go-test-clean  ## Go test all project
	$(GO_TEST) -cover -p=1 ./...

go-test-clean: go-fmt ## Run the clean cache tests of the project
	$(GO_CMD) clean -testcache

docker-compose-up: ## Run docker-compose services of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) up -d

docker-compose-down: ## Stop docker-compose services of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) down --remove-orphans

docker-compose-restart: docker-compose-down docker-compose-up ## Restart docker-compose services of project

docker-compose-logs: ## Logs docker-compose containers of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) logs -f

docker-compose-ps: ## List docker-compose containers of project
	$(DOCKER_COMPOSE_CMD) -f $(PATH_DOCKER_COMPOSE_FILE) ps

graphql-generate: ## Generate graphql
	$(GO_CMD) run github.com/99designs/gqlgen generate --verbose