package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/vitorsalgado/gopin/internal/utils/config"
	"github.com/vitorsalgado/gopin/internal/utils/database"
)

func ConnectDb(d time.Duration) *sql.DB {
	conf := config.Load()
	db := database.Connect(conf)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ticker := time.NewTicker(1 * time.Second)
	timeout := time.After(d)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("Something bad happened ... %v", r))
		}
	}()

	defer cancel()
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Trying to connect ...")

			if db == nil {
				db = database.Connect(conf)
			} else {
				if err := db.PingContext(ctx); err == nil {
					fmt.Println("Successfully connected with MySQL.")
					return db
				}
			}

		case <-timeout:
			log.Fatal("Unable to establish connection with MySQL instance. Exiting ...")
			return nil
		}

		time.Sleep(1 * time.Second)
	}
}
