package infrastructure

import (
	"golang_test/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/product/:id", func(context echo.Context) error { return c.Ticket.GetProductById(context) })
	e.POST("/ticket", func(context echo.Context) error { return c.Ticket.CreateTrx(context) })

	return e
}
