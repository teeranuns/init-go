package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	e.GET("/product/1234", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, product! 1234")
		//addBook(c)
		//addProduct(c)
	})

}
