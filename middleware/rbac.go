package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/database/models"
	"github.com/gin-rest-gorm-rbac-sample/utils"
)

func CheckPermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if error := recover(); error != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "must login"})
				c.Abort()
			}

		}()
		userObj := c.MustGet("user").(models.User)
		userPermissions := models.GetUserPermission(userObj.ID)
		if utils.Intersection(permissions, userPermissions) == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "You do not have the authority to access this api!"})
			c.Abort()
		}
	}
}
