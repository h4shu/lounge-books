package main

import (
	"os"
	"time"

	"github.com/h4shu/lounge-books/infrastructure"
	"github.com/h4shu/lounge-books/infrastructure/database"
	"github.com/h4shu/lounge-books/infrastructure/router"
)

func main() {
	var app = infrastructure.NewConfig().
		ContextTimeout(10 * time.Second).
		DbSQL(database.InstancePostgres)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceEcho).
		Start()
}
