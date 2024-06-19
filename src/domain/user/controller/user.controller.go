package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	service "init/project/src/domain/user/service"
)

var userService service.UserService

func InitUserController(service service.UserService) {
	userService = service
}

func GetUsers(c echo.Context) error {
	users, err := userService.GetAllUsers()

	fmt.Println(
		"get all users",
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not retrieve users"})
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, users)
}
