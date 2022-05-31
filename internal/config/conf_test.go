package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := Load()

	t.Run("it should load with default values when there's no .env file", func(t *testing.T) {
		assert.Equal(t, ":8080", config.Port)
		assert.Equal(t, "./docs/openapi/swagger-ui", config.SwaggerUiPath)
		assert.NotNil(t, config.MySQLConnectionString)
		assert.Equal(t, 10, config.MySQLMaxOpenConns)
		assert.Equal(t, 10, config.MySQLMaxIdleConns)
		assert.Equal(t, 10, config.MaxWorkers)
	})
}
