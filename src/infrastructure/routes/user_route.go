package router

import (
	user "init/project/src/domain/user/controller"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo, prefix string, middleware echo.MiddlewareFunc) {
	g := e.Group(prefix, middleware)
	g.GET("/", user.GetUsers)

	// e.GET("/users/:id", func(c echo.Context) error {
	// 	id := c.Param("id")
	// 	var user struct {
	// 		ID    int    `db:"id" json:"id"`
	// 		Name  string `db:"name" json:"name"`
	// 		Email string `db:"email" json:"email"`
	// 	}
	// 	err := db.Get(&user, "SELECT id, name, email FROM users WHERE id=$1", id)
	// 	if err != nil {
	// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "User not found"})
	// 	}
	// 	return c.JSON(http.StatusOK, user)
	// })
}
