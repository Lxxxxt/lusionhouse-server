package service

import (
	"lusionhouse-server/app/infrastructure/repository"
	"lusionhouse-server/app/service/user"
)

func NewUserService(repo repository.UserRepository) *user.UserServiceIml {
	return user.NewUserService(repo)
}
