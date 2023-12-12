package infrastructure

import (
	"golang_test/adapter/controller"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/product/:id", func(ctx echo.Context) error {
		resp, err := c.Ticket.GetProductById(ctx)
		if err != nil {
			report, ok := err.(*echo.HTTPError)
			if !ok {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			ctx.Logger().Error(report)
			return ctx.JSON(report.Code, report)
		}
		return ctx.JSON(http.StatusOK, resp)
	})
	e.POST("/ticket", func(context echo.Context) error { return c.Ticket.CreateTrx(context) })

	return e
}
