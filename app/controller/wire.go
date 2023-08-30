//go:build wireinject
// +build wireinject

package controller

import (
	controllerUser "lusionhouse-server/app/controller/user"
	"lusionhouse-server/app/infrastructure/repository"
	svcUser "lusionhouse-server/app/service/user"

	"github.com/google/wire"
)

func InitUserController(name string) *controllerUser.LoginController {
	wire.Build(repository.NewUserRepository, svcUser.NewUserService, controllerUser.NewUserController)
	return &controllerUser.LoginController{}
}
