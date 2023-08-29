package controller

import (
	controller "lusionhouse-server/app/controller/user"
	"lusionhouse-server/app/infrastructure/repository"
	"lusionhouse-server/app/service/middleware"
	"lusionhouse-server/app/service/user"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(g *gin.RouterGroup) {
	repo := repository.NewUserRepository()
	userSvc := user.NewUserService(repo)
	userController := controller.NewUserController(userSvc)

	// 路由
	userGroup := g.Group("/user")
	userGroup.POST("/register", userController.RegisterHandler)
	userGroup.POST("/login", userController.LoginHandler)
	userGroup.Use(middleware.MiddleWareAuth)
	userGroup.POST("/logout", userController.LogoutHandler)
	
}
