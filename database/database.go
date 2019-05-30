package database

import (
	"fmt"

	"github.com/gin-rest-gorm-rbac-sample/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:admin@/test?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true) // logs SQL
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
