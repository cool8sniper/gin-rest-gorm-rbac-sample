package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "must login"})
		c.Abort()
	}
}
