package main

import (
	"github.com/gin-rest-gorm-rbac-sample/api"
	"github.com/gin-rest-gorm-rbac-sample/database"
	"github.com/gin-rest-gorm-rbac-sample/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	db, _ := database.Initialize()
	app.Use(database.Inject(db))

	app.Use(middleware.JWTMiddleware())

	api.ApplyRoutes(app)
	app.Run()
}
