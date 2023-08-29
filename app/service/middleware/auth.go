package middleware

import (
	"log"
	"lusionhouse-server/app/infrastructure/repository"
	"lusionhouse-server/app/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddleWareAuth(g *gin.Context) {
	// 从cookie中获取session_id
	sessionID, err := g.Cookie("session_id")
	if err != nil {
		log.Printf("get session_id from cookie failed:%s\n", err)
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return

	}

	userSvc := user.NewUserService(repository.NewUserRepository())
	user, err := userSvc.FindUserBySessionId(g, sessionID)
	if err != nil {
		log.Printf("get user from session failed:%s\n", err)
	}
	if user == nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	g.Set("user", user)
	g.Set("session_id", sessionID)
	g.Next()

}
