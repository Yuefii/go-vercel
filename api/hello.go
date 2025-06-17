package api

import (
	pkg "go-vercel/api/_pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	app.GET("/v1/api/hello/:name", pkg.HelloNameHandler)

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
