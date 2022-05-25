package main

import (
	"time"

	"github.com/vitorsalgado/go-location-management/test"
)

func main() {
	test.ConnectDb(30 * time.Second)
}
