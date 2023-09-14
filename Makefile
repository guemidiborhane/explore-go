PROJECT_NAME := explore-go
user := $(shell id -u)
group := $(shell id -g)

go_version := $(shell cat go.mod | grep '^go' | sed -E "s/go\s//")

PORT?=3000
dc = USER_ID=$(user) GROUP_ID=$(group) GO_VERSION="$(go_version)" PORT=$(PORT) COMPOSE_PROJECT_NAME=$(PROJECT_NAME) docker-compose -p $(PROJECT_NAME) -f docker-compose.yaml $(1)$(2)
dr = @$(call dc, run --rm $(1)$(2))
COMPOSE_OPTIONS ?=
de = @$(call dc, exec $(COMPOSE_OPTIONS) $(1)$(2))

args = $(shell arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}})

include .env

.DEFAULT_GOAL := help

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=25
## Show this help message
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			sub(/:/, "", helpCommand); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: dc
## Run docker-compose
dc:
	@$(call dc, $(args))

.PHONY: pre-dev
pre-dev:
	@$(call dc, up -d db redis)

dev_services = backend frontend
.PHONY: dev
## Start development containers in foreground
dev: pre-dev history/go/fish history/node/fish
	EXTERNAL_PORT=$(PORT) $(call dc, up $(if $(args),$(args),$(dev_services)))

.PHONY: up
## Start containers in background
up:
	EXTERNAL_PORT=$(PORT) $(call dc, up -d $(args))

.PHONY: stop
## Stop containers
stop:
	EXTERNAL_PORT=$(PORT) $(call dc, stop $(args))

.PHONY: down
## Stop and remove all containers (including runners)
down:
	$(call dc, down --remove-orphans)

### Building and installing ###

.PHONY: setup
## Setup the project (build images, install deps, etc)
setup: down .env _bundless history/go/fish history/node/fish
	make docker-build
	$(call dr, backend go mod tidy)
	$(call dr, frontend yarn install)
	@echo "✅ ${GREEN}Build complete!${RESET}"
	@echo "❗ ${YELLOW}You should run specs with 'make test' to make sure everything is working ${RESET}"

.env: .env.example
	@test -f .env || cp .env.example .env

history/%:
	@if [ -d .docker/$@ ]; then rm -rfv .docker/$@; fi
	@if [ ! -f .docker/$@ ]; then touch .docker/$@; fi

.PHONY: _bundless
_bundless:
	-docker volume rm $(PROJECT_NAME)_go_pkg

.PHONY: docker-build
## Build docker images
docker-build:
	$(call dc, build --pull $(args))

.PHONY: build
## Build production image
build:
	$(call dc, -f docker-compose.prod.yaml, build production)
