package db

import (
	"database/sql"
	gopin "github.com/vitorsalgado/gopin/internal/config"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/vitorsalgado/gopin/internal/util/panicif"
)

// ConnectToMySQL opens a connection to a MySQL instance and returns a db object
func ConnectToMySQL(config *gopin.Config) *sql.DB {
	db, err := sql.Open("mysql", config.MySQLConnectionString)
	panicif.Err(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(config.MySQLMaxOpenConns)
	db.SetMaxIdleConns(config.MySQLMaxIdleConns)

	return db
}
