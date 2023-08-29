package main

import (
	"lusionhouse-server/app"

	"github.com/gin-gonic/gin"
)

func main() {
	err := app.Startup()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	app.RegisterHandlers(r)
	r.Run(":8080")
}
