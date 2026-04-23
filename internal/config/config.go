package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUrl   string
	SSLmode string
}

func New() *Config {
	return &Config{
		DBUrl:   getString("DB_URL", "postgres://postgres:password@localhost:5432/stockshelf?"),
		SSLmode: getString("sslmode", "disable"),
	}
}

func (c *Config) ConnString() string {
	return fmt.Sprintf(
		c.DBUrl, c.SSLmode,
	)
}

func getString(envName, defaultValue string) string {
	value, ok := os.LookupEnv(envName)
	if !ok || value == "" {
		return defaultValue
	}

	return value
}
