package router

import (
	"fmt"
	"time"

	"github.com/h4shu/lounge-books/adapter/controllers"
	"github.com/h4shu/lounge-books/adapter/gateways"
	"github.com/h4shu/lounge-books/adapter/presenters"
	"github.com/h4shu/lounge-books/application/interactors"
	"github.com/h4shu/lounge-books/application/repositories"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoEngine struct {
	router     *echo.Echo
	db         repositories.SQL
	port       Port
	ctxTimeout time.Duration
}

func newEchoServer(db repositories.SQL, port Port, t time.Duration) *echoEngine {
	return &echoEngine{
		router:     echo.New(),
		db:         db,
		port:       port,
		ctxTimeout: t,
	}
}

func (e *echoEngine) Listen() {
	e.setMiddlewares(e.router)
	e.setAppHandlers(e.router)
	e.router.Logger.Fatal(e.router.Start(fmt.Sprintf(":%d", e.port)))
}

func (e *echoEngine) setMiddlewares(router *echo.Echo) {
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
}

func (e *echoEngine) setAppHandlers(router *echo.Echo) {
	router.POST("/books", e.buildCreateBookHandler())

}

func (e *echoEngine) buildCreateBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		g, err := gateways.NewBookPostgres(e.db)
		if err != nil {
			return err
		}
		u := interactors.NewCreateBookInteractor(g, e.ctxTimeout)
		p := presenters.NewCreateBookPresenter()
		ctl := controllers.NewCreateBookController(u, p)
		ctl.Handle(c.Response(), c.Request())
		return nil
	}
}
