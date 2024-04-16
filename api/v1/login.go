package v1

import (
	"gin_start/server"
	"github.com/gin-gonic/gin"
)

func Route() {
	route := gin.Default()
	v1 := route.Group("/v1")
	{
		v1.POST("/login", server.UserLogin)

	}

	route.Run("8083")
}
