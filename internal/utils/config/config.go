package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/vitorsalgado/gopin/internal/utils/panicif"
)

// Config represent application configurations
type Config struct {
	Port                  string
	SwaggerUiPath         string
	MySQLConnectionString string
	MySQLMaxOpenConns     int
	MySQLMaxIdleConns     int
	MaxWorkers            int
}

// List of environment variables
const (
	EnvPort                  = "PORT"
	EnvSwaggerUIPath         = "SWAGGER_UI_PATH"
	EnvMySQLConnectionString = "MYSQL_CONNECTION_STRING"
	EnvMySQLMaxOpenConns     = "MYSQL_MAX_OPEN_CONNS"
	EnvMySQLMaxIdleConns     = "MySQL_MAX_IDLE_CONNS"
	EnvMaxWorkers            = "MAX_WORKERS"
)

// Default values for when environment variable is no set
const (
	DefPort                  = ":8080"
	DefSwaggerUiPath         = "./api/swagger-ui"
	DefMySQLConnectionString = "root:mysql123@tcp(127.0.0.1:3306)/go?parseTime=true"
	DefMySQLMaxOpenConns     = 10
	DefMySQLMaxIdleConns     = 10
	DefMaxWorker             = 10
)

// Load loads configuration using environment variables
func Load() *Config {
	// For now, we ignore the error if there's no .strEnv file
	_ = godotenv.Load()

	return &Config{
		Port:                  strEnv(EnvPort, DefPort),
		SwaggerUiPath:         strEnv(EnvSwaggerUIPath, DefSwaggerUiPath),
		MySQLConnectionString: strEnv(EnvMySQLConnectionString, DefMySQLConnectionString),
		MySQLMaxOpenConns:     intEnv(EnvMySQLMaxOpenConns, DefMySQLMaxOpenConns),
		MySQLMaxIdleConns:     intEnv(EnvMySQLMaxIdleConns, DefMySQLMaxIdleConns),
		MaxWorkers:            intEnv(EnvMaxWorkers, DefMaxWorker),
	}
}

func strEnv(k string, def string) string {
	if v := os.Getenv(k); v == "" {
		return def
	} else {
		return v
	}
}

func intEnv(k string, def int) int {
	if v := os.Getenv(k); v == "" {
		return def
	} else {
		c, err := strconv.Atoi(v)
		panicif.Err(err)
		return c
	}
}
