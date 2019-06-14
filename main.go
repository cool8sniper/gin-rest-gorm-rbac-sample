package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/gin-rest-gorm-rbac-sample/log"

	"github.com/gin-gonic/gin"
	"github.com/gin-rest-gorm-rbac-sample/api"
	"github.com/gin-rest-gorm-rbac-sample/database"
	"github.com/gin-rest-gorm-rbac-sample/database/models"
	"github.com/gin-rest-gorm-rbac-sample/database/mysql"
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

	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

	log.InitLog()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Logger.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Logger.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	app := gin.Default()

	db := mysql.Initialize()
	app.Use(database.Inject(db))
	models.Migrate(db)

	defer db.Close()

	app.Use(middleware.JWTMiddleware())

	api.ApplyRoutes(app)
	app.Run()

	log.Logger.Info("Ready to provide services.")

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Logger.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Logger.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}
