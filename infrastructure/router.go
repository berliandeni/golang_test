package infrastructure

import (
	"golang_test/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/konsumen/:id", func(context echo.Context) error { return c.Konsumen.GetById(context) })
	e.GET("/konsumen", func(context echo.Context) error { return c.Konsumen.GetAll(context) })
	e.POST("/konsumen", func(context echo.Context) error { return c.Konsumen.Create(context) })

	e.GET("/pinjaman/:id", func(context echo.Context) error { return c.Pinjaman.GetById(context) })
	e.POST("/pinjaman", func(context echo.Context) error { return c.Pinjaman.CreateTrx(context) })

	return e
}
