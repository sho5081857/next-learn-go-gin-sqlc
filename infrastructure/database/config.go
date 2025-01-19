package database

import "next-learn-go-gin-sqlc/pkg"

type Config struct {
	Host     string
	Database string
	Port     string
	User     string
	Password string
}

func NewConfigPostgres() *Config {
	return &Config{
		Host:     pkg.GetEnvDefault("DB_HOST", "localhost"),
		Database: pkg.GetEnvDefault("DB_DATABASE", "postgres"),
		Port:     pkg.GetEnvDefault("DB_PUBLISHED_PORT", "5432"),
		User:     pkg.GetEnvDefault("DB_USERNAME", "postgres"),
		Password: pkg.GetEnvDefault("DB_PASSWORD", "password"),
	}
}
