package api

import (
	"github.com/gin-rest-gorm-rbac-sample/api/user"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		user.ApplyRoutes(api)
	}
}
