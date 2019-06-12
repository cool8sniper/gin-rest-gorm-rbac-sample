package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/api/user"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/gin-rest-gorm-rbac-sample/docs"

)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		user.ApplyRoutes(api)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
