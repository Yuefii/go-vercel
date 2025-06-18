package api

import (
	h "go-vercel/api/_pkg/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()

	api := app.Group("/api")
	{
		api.GET("/users", h.GetUsers)
	}
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
