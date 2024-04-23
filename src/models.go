package src

import "gorm.io/gorm"

type (
	// User 对应 users 表
	User struct {
		gorm.Model
		Username string
		Password string
		Email    string
		Phone    string
	}
	// Product 对应 products 表
	Product struct {
		gorm.Model
		Name        string
		Description string
		Price       float64
		Stock       int
	}
	// Order 是数据库中的订单模型
	Order struct {
		gorm.Model
		UserID    uint
		ProductID uint
		Quantity  int
		Status    string
	}
	// Category 是数据库中的产品分类模型
	Category struct {
		gorm.Model
		Name     string
		Products []Product `gorm:"foreignkey:CategoryID"`
	}
	// Profile 是数据库中的用户个人资料模型
	Profile struct {
		gorm.Model
		UserID   uint
		Address  string
		Birthday string
	}
)
