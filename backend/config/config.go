package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PortNumber string
	CertFile   string
	KeyFile    string
	Database   DatabaseConfig
}

type DatabaseConfig struct {
	Path           string
	MigrationsPath string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	AppConfig = Config{
		PortNumber: mustGetEnv("PORT_NUMBER"),
		CertFile:   mustGetEnv("CERT_FILE"),
		KeyFile:    mustGetEnv("KEY_FILE"),
		Database:   loadDatabaseConfig("DB_PATH", "DB_MIGRATIONS_PATH"),
	}
}

func loadDatabaseConfig(pathKey, migrationsPathKey string) DatabaseConfig {
	return DatabaseConfig{
		Path:           mustGetEnv(pathKey),
		MigrationsPath: getEnv(migrationsPathKey, ""),
	}
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
