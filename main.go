package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	PassWord  string `gorm:"size 255;not null"`
	TelePhone string `gorm:"type:varchar(11);not null"`
}

func main() {
	db := InitDB()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		//  获取参数
		name := c.PostForm("name")
		password := c.PostForm("password")
		telephone := c.PostForm("telephone")
		// 数据验证

		if len(telephone) != 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}

		//如果名称不存在，创建一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, password, telephone)

		// 判断手机号是否存在，如果存在不允许用户注册
		if IsTelePhoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		}

		//  如果用户不存在创建用户
		newUser := User{
			Name:      name,
			PassWord:  password,
			TelePhone: telephone,
		}
		db.Create(&newUser)

		//  返回结果
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run(":8080")

}

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func InitDB() *gorm.DB {
	dsn := "root:123.com@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	return db
}

func IsTelePhoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where(telephone).First(&User{})
	if user.ID != 0 {
		return true
	}
	db.Table("users")
	// 创建数据表
	db.AutoMigrate(&User{})
	return false
}
