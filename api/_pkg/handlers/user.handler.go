package handlers

import (
	m "go-vercel/api/_pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, m.Users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, user := range m.Users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
