export SHELL := /bin/bash

PROJECT := gopin
API_MAIN := cmd/app/main.go

.ONESHELL:

.DEFAULT_GOAL := help
.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

up:
	docker-compose -p $(PROJECT) up --build --force-recreate

down:
	docker-compose down

run:
	go run ${API_MAIN}

.PHONY: test
test:
	go test -v ./internal/... ./cmd/...

.PHONY: coverage
coverage:
	mkdir -p coverage
	go test -v ./... -race -coverprofile=coverage/coverage.out -covermode=atomic
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html

test-e2e:
	chmod +x ./test/run.sh
	./test/run.sh

lint:
	go vet ./...

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/app ${API_MAIN}

deps:
	go mod download
	go get -v -t -d ./...

build-docker-compose:
	docker-compose build

dev:
	docker-compose -f ./docker-compose-dev.yml -p $(PROJECT).dev up --build
dev-cleanup:
	docker-compose -f ./docker-compose-dev.yml down --remove-orphans --rmi=all

install-nodemon:
	npm i nodemon -g
nodemon:
	nodemon --exec go run cmd/app/main.go --signal SIGTERM

.PHONY: docs
docs:
	@echo Navigate to: http://localhost:6060/
	godoc -http=:6060

.PHONY: swagger
swagger:
	docker run -p 8081:8080 -e URL=/doc/swagger.yml -v $$PWD/api/openapi:/usr/share/nginx/html/doc swaggerapi/swagger-ui
