package models

import "github.com/gin-rest-gorm-rbac-sample/lib/common"

type User struct {
	ID       float64
	Name     string
	Age      int
	Email    string
	Password string
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
	u.ID = m["id"].(float64)
	u.Name = m["name"].(string)
	u.Email = m["emial"].(string)

}
