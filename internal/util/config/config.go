package config

import (
	goenv "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Config represent application configurations
type Config struct {
	Debug                 bool   `env:"DEBUG,default=false"`
	Port                  string `env:"PORT,default=:8080"`
	SwaggerUiPath         string `env:"SWAGGER_UI_PATH,default=./docs/openapi/swagger-ui"`
	MySQLConnectionString string `env:"MYSQL_CONNECTION_STRING,default=root:mysql123@tcp(127.0.0.1:3306)/go?parseTime=true"`
	MySQLMaxOpenConns     int    `env:"MYSQL_MAX_OPEN_CONNS,default=10"`
	MySQLMaxIdleConns     int    `env:"MYSQL_MAX_IDLE_CONNS,default=10"`
	MaxWorkers            int    `env:"MAX_WORKERS,default=10"`
}

// Load loads configuration using environment variables
func Load() *Config {
	_ = godotenv.Load()

	var config Config
	_, err := goenv.UnmarshalFromEnviron(&config)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to load environment variables into struct")
		return nil
	}

	return &config
}
