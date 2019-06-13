package models

type Role struct {
	ID     uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name   string `json:"name" gorm:"Type:varchar(100);NOT NULL"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

type Permission struct {
	ID          uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"Type:varchar(100);NOT NULL"`
	DisplayName string `json:"displayName"`
	Desc        string `json:"desc"`
	Status      string `json:"status"`
}

type RolePermission struct {
	PermissionId uint `gorm:"primary_key"`
	RoleId       uint `gorm:"primary_key"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}

func (Permission) TableName() string {
	return "permission"
}

func (Role) TableName() string {
	return "role"
}
