package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := Load()

	t.Run("it should load with default values when there's no .strEnv file", func(t *testing.T) {
		assert.NotNil(t, config.Port)
		assert.Equal(t, DefPort, config.Port)
		assert.NotNil(t, config.SwaggerUiPath)
		assert.Equal(t, DefSwaggerUiPath, config.SwaggerUiPath)
		assert.NotNil(t, config.MySQLConnectionString)
		assert.Equal(t, DefMySQLConnectionString, config.MySQLConnectionString)
		assert.NotNil(t, config.MySQLMaxOpenConns)
		assert.Equal(t, DefMySQLMaxOpenConns, config.MySQLMaxOpenConns)
		assert.NotNil(t, config.MySQLMaxIdleConns)
		assert.Equal(t, DefMySQLMaxIdleConns, config.MySQLMaxIdleConns)
		assert.NotNil(t, config.MaxWorkers)
		assert.Equal(t, DefMaxWorker, config.MaxWorkers)
	})
}
