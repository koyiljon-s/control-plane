package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET / and GET /health
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "user-service",
		"endpoints": []string{
			"/api/register",
			"/api/login",
			"/api/me",
			"/api/users/:id",
			"/api/oauth/google",
		},
	})
}
