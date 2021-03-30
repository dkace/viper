package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"viper/common"
	"viper/model"
	"viper/util"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//  获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	telephone := c.PostForm("telephone")
	// 数据验证

	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//如果名称不存在，创建一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)

	// 判断手机号是否存在，如果存在不允许用户注册
	if IsTelePhoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	//  如果用户不存在创建用户
	newUser := model.User{
		Name:      name,
		PassWord:  password,
		TelePhone: telephone,
	}
	DB.Create(&newUser)

	//  返回结果
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func IsTelePhoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	db.Table("users")
	// 创建数据表
	db.AutoMigrate(&user)
	return false
}
