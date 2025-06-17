package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	app.GET("/v1/api/hello/:name", func(ctx *gin.Context) {

		name := ctx.Param("name")
		if name == "" {
			ctx.JSON(400, gin.H{
				"message": "name not found",
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})

	app.GET("/v1/api/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"id": ctx.Query("id"),
			},
		})
	})
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
