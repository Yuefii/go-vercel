package api

import (
	"go-vercel/api/_pkg/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := gin.New()

	api := app.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUserByID)
		api.GET("/categories", handlers.GetCategories)
		api.GET("/categories/:id", handlers.GetCategoriesByID)
	}

	app.ServeHTTP(w, r)
}
