package controller

import (
	"lusionhouse-server/app/service/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(g *gin.RouterGroup) {
	// repo := repository.NewUserRepository()
	// userSvc := svcUser.NewUserService(repo)
	// userController := controllerUser.NewUserController(userSvc)
	userController := InitUserController("user")
	// 路由
	userGroup := g.Group("/user")
	userGroup.POST("/register", userController.RegisterHandler)
	userGroup.POST("/login", userController.LoginHandler)
	userGroup.Use(middleware.MiddleWareAuth)
	userGroup.POST("/logout", userController.LogoutHandler)

}
