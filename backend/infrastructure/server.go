package infrastructure

import (
	"log"
	"strconv"
	"time"

	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/h4shu/lounge-books/infrastructure/database"
	"github.com/h4shu/lounge-books/infrastructure/router"
)

type config struct {
	dbSQL         repositories.SQL
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) DbSQL(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		log.Fatalln(err, "Could not make a connection to the database")
	}

	log.Println("Successfully connected to the SQL database")

	c.dbSQL = db
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.dbSQL,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
