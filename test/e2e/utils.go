package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/vitorsalgado/gopin/internal/util/config"
	"log"
	"time"

	"github.com/vitorsalgado/gopin/internal/util/db"
)

func ConnectDb(d time.Duration) *sql.DB {
	conf := config.Load()
	database, err := db.ConnectToMySQL(conf)
	if err != nil {
		log.Printf("error connecting to the database. will try to reconnect. reason %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ticker := time.NewTicker(1 * time.Second)
	timeout := time.After(d)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Sprintf("something bad happened ... %v", r))
		}
	}()

	defer cancel()
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("trying to connect ...")

			if database == nil {
				database, err = db.ConnectToMySQL(conf)
				log.Println(err)
			} else {
				if err := database.PingContext(ctx); err == nil {
					fmt.Println("successfully connected with MySQL.")
					return database
				}
			}

		case <-timeout:
			log.Fatal("unable to establish connection with MySQL instance. Exiting ...")
			return nil
		}

		time.Sleep(1 * time.Second)
	}
}
