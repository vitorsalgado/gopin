PROJECT := gopin
REGISTRY := localhost:5000
IMAGE := $(REGISTRY)/$(PROJECT)
MAIN := cmd/app/main.go
GOPIN_DOCKER_COMPOSE_ROOT := ./deployments/local

.ONESHELL:
.DEFAULT_GOAL := help

# allow user specific optional overrides
-include Makefile.overrides

export

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up: ## run local environment with all service dependencies using with docker compose
	@docker-compose -f $(GOPIN_DOCKER_COMPOSE_ROOT)/docker-compose.yml -p $(PROJECT) up --build --force-recreate

down: ## tear down local docker compose environment
	@docker-compose -f $(GOPIN_DOCKER_COMPOSE_ROOT)/docker-compose.yml down

run: ## run application
	@go run $(MAIN)

.PHONY: test
test: ## run unit tests
	@go test -v ./internal/... ./cmd/...

test-e2e: ## run end-to-end tests
	@chmod +x ./test/e2e/run.sh
	./test/e2e/run.sh

test-all: test test-e2e ## run all tests

.PHONY: bench
bench: ## run benchmarks
	@go test -v ./... -bench=. -count 2 -run=^#

.PHONY: coverage
coverage: ## run tests and generate coverage report
	@mkdir -p coverage
	@go test -v ./... -race -coverprofile=coverage/coverage.out -covermode=atomic
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html

vet: ## check go code
	@go vet ./...

fmt: ## run gofmt in all project files
	@go fmt ./...

check: vet ## check source code
	@staticcheck ./...

.PHONY: build
build: ## build application
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/gopin $(MAIN)

deps: ## check dependencies
	@go mod verify

download: ## download dependencies
	@go mod download

build-docker-compose: ## build docker compose
	@docker-compose -f $(GOPIN_DOCKER_COMPOSE_ROOT)/docker-compose.yml build

build-docker: ## build docker image
	@docker build -t $(IMAGE) .

dev: ## run local development environment with hot reload using docker compose
	@docker-compose -f $(GOPIN_DOCKER_COMPOSE_ROOT)/docker-compose-dev.yml -p $(PROJECT).dev up --build

clean-dev: ## cleanup local development environment
	@docker-compose -f $(GOPIN_DOCKER_COMPOSE_ROOT)/docker-compose-dev.yml down --remove-orphans --rmi=all

.PHONY: swagger
swagger:
	@echo "preparing swagger-ui"
	@tar -xf docs/openapi/swagger-ui.tar.gz
	@cp ./docs/openapi/swagger-initializer.js docs/openapi/swagger-ui/swagger-initializer.js

swagger-docker: ## run swagger documentation with docker
	@docker run -p 8081:8080 -e URL=/doc/swagger.yml -v $$PWD/docs/openapi:/usr/share/nginx/html/doc swaggerapi/swagger-ui

install-staticcheck: ## download and install staticcheck tool locally
	@echo "installing staticcheck locally"
	@go install honnef.co/go/tools/cmd/staticcheck@latest

prep: swagger install-staticcheck ## prepare local development environment
	@echo "preparing local tools"
	@npm i
