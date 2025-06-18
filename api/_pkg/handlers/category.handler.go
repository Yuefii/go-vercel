package handlers

import (
	"go-vercel/api/_pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Categories)
}
