package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloNameHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "name not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": fmt.Sprintf("Hello %s!", name),
	})
}
