package models

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(User{})
	db.AutoMigrate(UserRole{})
	db.AutoMigrate(Role{})
	db.AutoMigrate(Permission{})
	db.AutoMigrate(RolePermission{})

}
