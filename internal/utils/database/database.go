package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/vitorsalgado/gopin/internal/utils/config"
	"github.com/vitorsalgado/gopin/internal/utils/panicif"
)

var db *sql.DB

const driver = "mysql"

// Connect opens a connection to a MySQL instance and returns a database object
func Connect(config *config.Config) *sql.DB {
	var err error

	db, err = sql.Open(driver, config.MySQLConnectionString)
	panicif.Err(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(config.MySQLMaxOpenConns)
	db.SetMaxIdleConns(config.MySQLMaxIdleConns)

	return db
}
