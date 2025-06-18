package handlers

import (
	"go-vercel/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Categories)
}
