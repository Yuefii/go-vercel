package api

import (
	h "go-vercel/api/_pkg/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	app := gin.New()

	app.GET("/api/users", h.GetUsers)
	app.GET("/api/users/:id", h.GetUserByID)
	app.GET("/api/categories", h.GetCategories)

	app.ServeHTTP(w, r)
}
