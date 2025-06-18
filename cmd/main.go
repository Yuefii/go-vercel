package main

import (
	"go-vercel/api/_pkg/handlers"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {
	app = gin.New()
	api := app.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUserByID)
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategoriesByID)
	}

	app.Run(":8080")
}
