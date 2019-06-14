package mysql

import (
	"fmt"

	"github.com/gin-rest-gorm-rbac-sample/lib/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
)

var db *gorm.DB
var err error

// Initialize initializes the database
func Initialize() *gorm.DB {
	mysqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	db, err = gorm.Open(setting.DatabaseSetting.Type, mysqlConnectionString)
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	return db
}

func GetMysql() *gorm.DB {
	return db
}
