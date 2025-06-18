package handlers

import (
	"go-vercel/api/_pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range models.Users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
