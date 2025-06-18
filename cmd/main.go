package main

import (
	h "go-vercel/api/_pkg/handlers"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {
	app = gin.New()
	api := app.Group("/api")
	{
		api.GET("/users", h.GetUsers)
	}

	app.Run(":8080")
}
