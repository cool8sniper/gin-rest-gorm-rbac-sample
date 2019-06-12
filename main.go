package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/api"
	"github.com/gin-rest-gorm-rbac-sample/database"
	"github.com/gin-rest-gorm-rbac-sample/middleware"
)


// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {

	app := gin.Default()


	db, _ := database.Initialize()
	app.Use(database.Inject(db))

	app.Use(middleware.JWTMiddleware())

	api.ApplyRoutes(app)
	app.Run()
}
