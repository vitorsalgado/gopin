package main

import (
	"time"

	"github.com/vitorsalgado/gopin/test"
)

func main() {
	test.ConnectDb(30 * time.Second)
}
