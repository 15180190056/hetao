package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLogin(c *gin.Context) {
	// 从请求中获取数据，例如从POST请求体中解析JSON数据
	// ...

	user := c.PostForm("username")
	passwd := c.PostForm("password")
	var userloginform = LoginForm{
		Username: user,
		Password: passwd,
	}
	var dbuserform = LoginForm{
		Username: "hetao",
		Password: "123456",
	}
	// 处理逻辑
	// ...
	if userloginform == dbuserform {
		// 发送响应数据
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User logged in successfully",
		})
	}

}
