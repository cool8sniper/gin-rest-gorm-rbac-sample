package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/api/user"
	"github.com/gin-rest-gorm-rbac-sample/lib/common"
	"github.com/gin-rest-gorm-rbac-sample/log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/gin-rest-gorm-rbac-sample/docs"
)

func version(c *gin.Context) {
	log.Logger.Info("access version api")

	c.JSON(200, common.JSON{
		"version": "1.2.2",
	})

}

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("version", version)
		user.ApplyRoutes(api)

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
