package models

import (
	"fmt"

	"github.com/gin-rest-gorm-rbac-sample/database/mysql"
	"github.com/gin-rest-gorm-rbac-sample/lib/common"
)

type User struct {
	ID       uint
	Name     string
	Age      uint8
	Email    string
	Password string
}

type UserRole struct {
	UserId uint `gorm:"primary_key"`
	RoleId uint `gorm:"primary_key"`
}

func (UserRole) TableName() string {
	return "user_role"
}

func (User) TableName() string {
	return "user"
}

func (u User) Serialize() common.JSON {
	return common.JSON{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email}
}

func (u *User) Read(m common.JSON) {
	u.ID = m["id"].(uint)
	u.Name = m["name"].(string)
	u.Email = m["emial"].(string)

}

func getRolePermissionByRoleIds(roleIds []uint) []string {
	type NameResult struct {
		Name string
	}
	db := mysql.GetMysql()
	var nameResult []NameResult
	db.Table("permission").Select("name").
		Joins("left join role_permission on role_permission.permission_id = permission.id").
		Where("role_permission.role_id in (?)", roleIds).Scan(&nameResult)
	var result []string
	for _, v := range nameResult {
		result = append(result, v.Name)
	}
	return result
}

func GetUserPermission(userId uint) []string {
	db := mysql.GetMysql()
	fmt.Println(db, "=============")
	var userRoles []UserRole
	db.Where("user_id = ?", userId).Find(&userRoles)
	var roleIds []uint
	for _, v := range userRoles {
		roleIds = append(roleIds, v.RoleId)
	}
	return getRolePermissionByRoleIds(roleIds)

}
