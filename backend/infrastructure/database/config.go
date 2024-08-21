package database

import "os"

type config struct {
	host     string
	port     string
	driver   string
	user     string
	password string
	database string
}

func newConfigPostgres() *config {
	return &config{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     os.Getenv("POSTGRES_PORT"),
		driver:   os.Getenv("POSTGRES_DRIVER"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		database: os.Getenv("POSTGRES_DB"),
	}
}
