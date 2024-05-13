package configs_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhikariz/go-commerce/configs"
)

func TestNewConfig(t *testing.T) {
	// Prepare a test .env file content
	envContent := []byte(`
ENV=test
PORT=9090
POSTGRES_HOST=mydbhost
POSTGRES_USER=dbuser
POSTGRES_PASSWORD=dbpass
POSTGRES_DATABASE=mydatabase
JWT_SECRET_KEY=myjwtsecret
REDIS_HOST=myredishost
REDIS_PORT=6380
ENCRYPT_SECRET_KEY=myencryptsecret
ENCRYPT_IV=myencryptiv
`)

	// Create a temporary .env file for testing
	tmpEnvFile := ".env.test"
	err := os.WriteFile(tmpEnvFile, envContent, 0644)
	if err != nil {
		t.Fatalf("failed to create temporary .env file: %v", err)
	}
	defer os.Remove(tmpEnvFile) // Clean up after the test

	// Test NewConfig function with the temporary .env file
	cfg, err := configs.NewConfig(tmpEnvFile)
	if err != nil {
		t.Fatalf("NewConfig failed: %v", err)
	}

	// Verify the parsed Config object
	expected := &configs.Config{
		Env:  "test",
		Port: "9090",
		Postgres: configs.PostgresConfig{
			Host:     "mydbhost",
			User:     "dbuser",
			Password: "dbpass",
			Database: "mydatabase",
			Port:     "5432",
		},
		JWT: configs.JwtConfig{
			SecretKey: "myjwtsecret",
		},
		Redis: configs.RedisConfig{
			Host: "myredishost",
			Port: "6380",
		},
		Encrypt: configs.EncryptConfig{
			SecretKey: "myencryptsecret",
			IV:        "myencryptiv",
		},
	}

	assert.Equal(t, expected, cfg, "parsed configuration does not match expected")

	// Additional assertions can be added to test specific fields or scenarios
	// For example:
	assert.Equal(t, "test", cfg.Env, "environment value mismatch")
	assert.Equal(t, "mydbhost", cfg.Postgres.Host, "Postgres host mismatch")
}

func TestNewConfigError(t *testing.T) {
	_, err := configs.NewConfig("")
	assert.Error(t, err, "expected error when .env file is not found")
}
