package main

import (
	"github.com/vitorsalgado/gopin/test/e2e"
	"time"
)

func main() {
	e2e.ConnectDb(30 * time.Second)
}
