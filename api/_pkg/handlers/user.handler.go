package handlers

import (
	m "go-vercel/api/_pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, m.Users)
}
