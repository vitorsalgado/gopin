#!/bin/bash

function clean() {
  docker-compose -f "$GOPIN_DOCKER_COMPOSE_ROOT/docker-compose.yml" down --volumes --remove-orphans
  docker-compose -f "$GOPIN_DOCKER_COMPOSE_ROOT/docker-compose.yml" down rm
}

trap clean EXIT

docker-compose -f "$GOPIN_DOCKER_COMPOSE_ROOT/docker-compose.yml" up -d --build --remove-orphans

echo "Waiting to ensure everything is online before tests"
go run ./test/e2e/cmd/ping/main.go

echo "Run all integration and e2e tests"
go test -v ./test/...
