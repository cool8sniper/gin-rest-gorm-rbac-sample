package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/api"
	"github.com/gin-rest-gorm-rbac-sample/database"
	"github.com/gin-rest-gorm-rbac-sample/lib/setting"
	"github.com/gin-rest-gorm-rbac-sample/middleware"
)

func init() {
	setting.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/cool8sniper/gin-rest-gorm-rbac-sample
// @license.name MIT
// @license.url https://github.com/cool8sniper/gin-rest-gorm-rbac-sample/blob/master/LICENSE
func main() {

	app := gin.Default()

	db, _ := database.Initialize()
	defer db.Close()
	app.Use(database.Inject(db))

	app.Use(middleware.JWTMiddleware())

	api.ApplyRoutes(app)
	app.Run()
}
