package setup

import (
	controller "init/project/src/domain/user/controller"
	repository "init/project/src/domain/user/repositories"
	service "init/project/src/domain/user/service"
)

func Initialize() {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	controller.InitUserController(userService)
}
