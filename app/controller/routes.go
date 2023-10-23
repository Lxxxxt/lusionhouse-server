package controller

import (
	"fmt"
	"log"
	"lusionhouse-server/app/service/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func RegisterUserHandler(g *gin.RouterGroup) {
	userController := InitUserController("user")
	// 路由
	userGroup := g.Group("/user")
	userGroup.POST("/register", userController.RegisterHandler)
	userGroup.POST("/login", userController.LoginHandler)
	userGroup.Use(middleware.MiddleWareAuth)
	userGroup.POST("/logout", userController.LogoutHandler)

}
func RegisterWsHandler(g *gin.RouterGroup) {
	g.GET("/echo", func(ctx *gin.Context) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
		c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				fmt.Println("write:", err)
				break
			}
		}

	})

}
