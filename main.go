package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		//  获取参数
		name := c.PostForm("name")
		password := c.PostForm("password")
		telephone := c.PostForm("telephone")
		// 数据验证

		if len(telephone) != 0 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422, "msg": "手机号必须为11位"})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422, "msg": "密码不能少于6位"})
			return
		}

		//如果名称不存在，创建一个10位的随机字符串
		if len(name) == 0 {
			name = RandomString(10)
		}
		log.Println(name, password, telephone)
		//  创建用户
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
