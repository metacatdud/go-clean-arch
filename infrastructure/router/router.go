package router

import (
	"github.com/labstack/echo"
	"github.com/metacatdud/go-boilerplate/interface/controller"
)

func NewRouter(c controller.Application) *echo.Echo {
	e := echo.New()

	// middleare go here

	// routes go here
	e.GET("/users", func(ctx echo.Context) error {
		return c.Get(ctx)
	})

	return e
}
