package v1

import (
	"gin_start/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route() {
	route := gin.Default()
	// 设置CORS中间件
	route.Use(func(c *gin.Context) {
		// 允许所有的域进行跨域请求
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 允许的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 预检间隔时间
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// 允许发送带凭证的请求
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理OPTIONS请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	})
	v1 := route.Group("/v1")
	{
		v1.POST("/login", server.Handlelogin)
		v1.GET("/ping", server.Ping)

	}

	route.Run(":8083")
}
