package v1

import (
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

	//商品管理
	goods := route.Group("goods")
	{
		//添加商品信息
		goods.POST("/AddProduct", goods.AddProduct)
		//商品编辑
		goods.POST("/EditProduct", goods.EditProduct)
		//商品上架
		goods.POST("/ShelveProduct", goods.ShelveProduct)
		//商品下架
		goods.POST("/UnshelveProduct", goods.UnshelveProduct)
	}

	//订单管理
	order := route.Group("order")
	{
		// 打印所有订单信息
		order.GET("list", order.PrintAllOrders)
		//修改订单信息
		order.GET("list", order.UpdateOrderStatus)
		// 查询订单
		order.GET("list", order.GetOrder)
		// 删除订单  DeleteOrder
		order.GET("list", order.DeleteOrder)
		// 添加订单  AddOrder
		order.GET("list", order.AddOrder)
	}

	//用户管理
	user := route.Group("user")
	{
		// 注册新用户
		user.POST("/RegisterUser", user.RegisterUser)
		// 用户登录
		user.POST("/LoginUser", user.LoginUser)
		// 更新用户信息
		user.POST("/UpdateUser", user.UpdateUser)
		// 删除用户
		user.POST("/DeleteUser", user.DeleteUser)

	}
	route.Run(":8083")
}
