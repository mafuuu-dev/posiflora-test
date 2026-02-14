# ----------------------------------------------------------------------
# Configuration:
# ----------------------------------------------------------------------

.SILENT:
.ONESHELL:

THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY: help
help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

args := --env-file .docker/.env \
    	--env-file .docker/pgsql/.env

d := COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker
dc := $(d) compose $(args)

## ----------------------------------------------------------------------
## Environment:
## ----------------------------------------------------------------------

.PHONY: setup
setup: ## Environment setup
	make frontend.env
	make down
	make build
	make frontend.install
	make up
	make stop

.PHONY: up
up: ## Environment up
	$(dc) up -d --force-recreate --remove-orphans

.PHONY: stop
stop: ## Environment stop
	$(dc) stop

.PHONY: restart
restart: ## Environment restart
	make stop
	make up

.PHONY: down
down: ## Environment down
	make stop
	$(dc) down --remove-orphans --volumes

.PHONY: build
build: ## No cache building containers
	$(dc) build --no-cache

.PHONY: logs
logs: ## Show containers logs
	$(dc) logs -f

## ----------------------------------------------------------------------
## Application:
## ----------------------------------------------------------------------

.PHONY: app.test
app.test: ## Run app test
	$(dc) exec api go test ./test/e2e -v

## ----------------------------------------------------------------------
## Frontend:
## ----------------------------------------------------------------------

.PHONY: frontend.env
frontend.env: ## Set frontend .env
	cp -f ./frontend/.env.example ./frontend/.env

.PHONY: frontend.install
frontend.install: ## Run frontend npm install
	$(dc) run --remove-orphans ui npm install