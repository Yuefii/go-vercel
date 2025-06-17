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

	v1 := app.Group("/v1/api")
	{
		v1.GET("/with-gin/:name", pkg.HelloNameHandler)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
