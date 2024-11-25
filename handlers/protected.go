package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProtectedHandler handles requests to protected routes
func ProtectedHandler(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + username + "!"})
}
