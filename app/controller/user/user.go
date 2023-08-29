package controller

import (
	"log"
	"lusionhouse-server/app/domain"
	"lusionhouse-server/app/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type loginController struct {
	userService *user.UserServiceIml
}

func NewUserController(userService *user.UserServiceIml) *loginController {
	return &loginController{userService: userService}
}

func (l *loginController) LoginHandler(c *gin.Context) {
	var requestData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBind(&requestData)

	if err != nil {
		log.Printf("bind request data failed:%s\n", err)
		c.AbortWithError(http.StatusBadRequest, errors.Wrap(err, "invalid request data"))
		return
	}

	u, err := l.userService.Authenticate(c, requestData.Username, requestData.Password)
	if err == user.ErrAuthFailed {
		c.AbortWithError(http.StatusUnauthorized, errors.Wrap(err, "authenticate failed"))
		return
	}
	newSess := uuid.NewV4().String()
	err = l.userService.SetUserSession(c, newSess, u)
	if err != nil {
		log.Printf("set user session failed:%s\n", err)
		c.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "set user session failed"))
		return
	}
	c.SetCookie("session_id", newSess, 3600*24*365, "/", ".toddliu.com", false, true)
	var responseData = struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}{
		ID:       u.ID,
		Username: u.Name,
	}
	c.JSON(http.StatusOK, responseData)
}

func (c *loginController) LogoutHandler(g *gin.Context) {
	sessionId, ok := g.Get("session_id")
	if !ok {
		log.Printf("get user sessionId from context failed\n")
		g.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	err := c.userService.ClearUserSession(g, sessionId.(string))
	if err != nil {
		log.Printf("clear user session failed:%s\n", err)
		g.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "clear user session failed"))
		return
	}

}
func (l *loginController) RegisterHandler(g *gin.Context) {
	var requestData struct {
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		Emial       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}
	err := g.ShouldBind(&requestData)
	if err != nil {
		log.Printf("bind request data failed:%s\n", err)
		g.AbortWithError(http.StatusBadRequest, errors.Wrap(err, "invalid request data"))
		return
	}
	user := &domain.User{
		Name:        requestData.Username,
		Password:    requestData.Password,
		Email:       requestData.Emial,
		PhoneNumber: requestData.PhoneNumber,
	}
	err = l.userService.CreateUser(g, user)
	if err != nil {
		log.Printf("create user failed:%s\n", err)
		g.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "create user failed"))
		return
	}
	newSess := uuid.NewV4().String()
	err = l.userService.SetUserSession(g, newSess, user)
	if err != nil {
		log.Printf("set user session failed:%s\n", err)
		g.AbortWithError(http.StatusInternalServerError, errors.Wrap(err, "set user session failed"))
		return
	}
	g.SetCookie("session_id", newSess, 3600*24*365, "/", ".toddliu.com", false, true)
	g.JSON(http.StatusCreated, gin.H{
		"message": "created",
	})

}
