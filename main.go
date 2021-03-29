package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/api/auth/register", func(c *gin.Context) {
		//  获取参数
		//name := c.PostForm("name")
		password := c.PostForm("password")
		telephone := c.PostForm("telephone")
		// 数据验证
		if len(telephone) != 0 { // 判断手机号码是否存在
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "手机号必须为11位",
			})
		}

		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg":  "",
			})
		}

		//  创建用户
		//  返回结果

		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run(":8080")

}
