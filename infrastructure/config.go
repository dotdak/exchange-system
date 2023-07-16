package infrastructure

import "os"

type DbConfig struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         os.Getenv("DB_HOST"),
		Port:         os.Getenv("DB_PORT"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}
