#!/bin/bash

function clean() {
  docker-compose down --volumes --remove-orphans
  docker-compose rm -f
}

trap clean EXIT

docker-compose up -d --build --remove-orphans

echo "Waiting to ensure everything is online before tests"
go run ./cmd/ping/main.go

echo "Run all integration and e2e tests"
go test -v ./test/...
