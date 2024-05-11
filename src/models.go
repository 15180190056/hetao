package src

import (
	"gorm.io/gorm"
)

// User 是数据库中的用户模型
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
	Phone    string `gorm:"unique"`
}

//
//// Product 对应数据库中的products表
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//	// 添加CategoryID作为外键
//	CategoryID uint `gorm:"column:category_id;ForeignKey:References"`
//}

//// Order 是数据库中的订单模型
//type Order struct {
//	gorm.Model
//	UserID    uint
//	ProductID uint
//	Quantity  int
//	Status    string
//}

// Category 对应数据库中的categories表
//type Category struct {
//	gorm.Model
//	Name     string    `gorm:"column:name"`
//	Products []Product `gorm:"foreignkey:CategoryID"` // 指定外键为Product中的CategoryID
//}

// Profile 是数据库中的用户个人资料模型
//type Profile struct {
//	gorm.Model
//	UserID   uint
//	Address  string
//	Birthday string
//}
