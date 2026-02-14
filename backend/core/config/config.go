package config

import (
	"os"
)

type Config struct {
	ApiHost      string
	TestBotToken string
	TestChatID   string
	PGHost       string
	PGPort       string
	PGDatabase   string
	PGUser       string
	PGPassword   string
	PGSSLMode    string
}

func Load() *Config {
	return &Config{
		ApiHost:      getEnv("API_HOST", "localhost:8080"),
		TestBotToken: getEnv("TEST_BOT_TOKEN", ""),
		TestChatID:   getEnv("TEST_CHAT_ID", ""),
		PGHost:       getEnv("POSTGRES_HOST", "localhost"),
		PGPort:       getEnv("POSTGRES_PORT", "5432"),
		PGDatabase:   getEnv("POSTGRES_DB", "postgres"),
		PGUser:       getEnv("POSTGRES_USER", "user"),
		PGPassword:   getEnv("POSTGRES_PASSWORD", "password"),
		PGSSLMode:    getEnv("POSTGRES_SSLMODE", "disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
