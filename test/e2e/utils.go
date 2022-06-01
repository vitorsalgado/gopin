package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vitorsalgado/gopin/internal/config"
	"log"
	"time"

	"github.com/vitorsalgado/gopin/internal/util/db"
)

func ConnectDb(d time.Duration) *sql.DB {
	conf := config.Load()
	database := db.ConnectToMySQL(conf)
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

			if database == nil {
				database = db.ConnectToMySQL(conf)
			} else {
				if err := database.PingContext(ctx); err == nil {
					fmt.Println("Successfully connected with MySQL.")
					return database
				}
			}

		case <-timeout:
			log.Fatal("Unable to establish connection with MySQL instance. Exiting ...")
			return nil
		}

		time.Sleep(1 * time.Second)
	}
}
