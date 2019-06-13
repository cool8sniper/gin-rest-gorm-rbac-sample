package database

import (
	"fmt"

	"github.com/gin-rest-gorm-rbac-sample/database/models"
	"github.com/gin-rest-gorm-rbac-sample/lib/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
)

var db *gorm.DB

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	db, err := gorm.Open(setting.DatabaseSetting.Type, dbString)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}

func GetMysql() *gorm.DB {
	return db
}
