package handlers

import (
	"go-vercel/api/_pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": models.Categories,
	})
}

func GetCategoriesByID(c *gin.Context) {
	id := c.Param("id")

	for _, category := range models.Categories {
		if category.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"data": category,
			})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Category not found"})
}
