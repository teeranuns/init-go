package user

import (
	user "init/project/src/domain/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// เปลี่ยนจากการใช้ http.ResponseWriter และ *http.Request เป็น echo.Context
	var users = []user.UserResponse{
		{ID: 1, Name: "12345", Email: "Book 1"},
		{ID: 2, Name: "23456", Email: "Book 2"},
	}

	users = append(users, user.UserResponse{ID: 3, Name: "34567", Email: "Book 3"})

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, users)
}
