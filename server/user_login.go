package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 定义处理POST请求的handler函数
func Handlelogin(c *gin.Context) {
	// 创建一个LoginData实例
	var loginData LoginData

	// 绑定JSON数据到结构体
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在这里处理登录逻辑，例如验证用户名和密码
	var userdata = LoginData{
		Username: "hetao",
		Password: "123456",
	}
	if loginData == userdata {
		// 返回响应
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
