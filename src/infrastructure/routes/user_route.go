package router

import (
	"github.com/labstack/echo/v4"

	user "init/project/src/interfaces/user"
)

func UserRoute(e *echo.Echo) {
	e.GET("/users/1234", user.GetUsers)

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
