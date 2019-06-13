package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/database/models"
	"github.com/gin-rest-gorm-rbac-sample/utils"
	"github.com/jinzhu/gorm"
)

func CheckPermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "must login"})
			c.Abort()
		}
		db := c.MustGet("db").(*gorm.DB)
		userPermissions := models.GetUserPermission(db, user.(models.User).ID)
		if utils.Intersection(permissions, userPermissions) == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "You do not have the authority to access this api!"})
			c.Abort()
		}
	}
}
