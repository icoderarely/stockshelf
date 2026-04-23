package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DbName  string
	SSLMode string
}

func New() *Config {
	return &Config{
		Host:    getString("POSTGRES_HOST", "localhost"),
		Port:    getString("POSTGRES_PORT", "5432"),
		User:    getString("POSTGRES_USER", "postgres"),
		Pass:    getString("POSTGRES_PASSWORD", "password"),
		DbName:  getString("POSTGRES_DB", "stockshelf"),
		SSLMode: getString("sslmode", "disable"),
	}
}

func (c *Config) ConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Pass, c.Host, c.Port, c.DbName, c.SSLMode,
	)
}

func getString(envName, defaultValue string) string {
	value, ok := os.LookupEnv(envName)
	if !ok || value == "" {
		return defaultValue
	}

	return value
}
