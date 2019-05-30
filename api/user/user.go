package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-rest-gorm-rbac-sample/database/models"
	"github.com/gin-rest-gorm-rbac-sample/lib/common"
	"github.com/gin-rest-gorm-rbac-sample/middleware"
	"github.com/gin-rest-gorm-rbac-sample/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User = models.User

func getUser(c *gin.Context) {
	userID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var user User
	db.First(&user, userID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(200, user.Serialize())
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email    string `json:"email" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(body)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "No todo found!"})
		return
	}
	// check existancy
	var exists User
	if err := db.Where("email = ?", body.Email).First(&exists).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "email existancy!"})
		return
	}

	hash, hashErr := utils.HashPassword(body.Password)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "message": "email existancy!"})
		return
	}

	// create user
	user := User{
		Name:     body.Name,
		Email:    body.Email,
		Password: hash,
	}

	db.NewRecord(user)
	db.Create(&user)

	c.JSON(200, common.JSON{
		"user": user.Serialize(),
	})
}

func generateToken(data common.JSON) (string, error) {

	//  token is valid for i hour
	date := time.Now().Add(time.Hour * 1)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  date.Unix(),
	})

	// get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		return "", readErr
	}
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

func login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "email and password are must!"})
		return
	}

	// check existancy
	var user User
	if err := db.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "email existancy!"})
		return
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "password wrong!"})
		return
	}

	serialized := user.Serialize()
	token, _ := generateToken(serialized)
	c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)

	c.JSON(200, common.JSON{
		"user": user.Serialize(),
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/")
	{
		user.POST("user", middleware.Authorized, create)
		user.GET("user/:id", middleware.Authorized, getUser)
		user.POST("login", login)

	}
}
