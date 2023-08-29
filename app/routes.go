package app

import (
	controller "lusionhouse-server/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	g := r.Group("/api/v1")
	controller.RegisterUserHandler(g)

}
