package src

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var dsn string = "root:123456@tcp(121.40.117.237:3306)/SpotProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 自动迁移模式，根据模型创建表
	// 注意：在生产环境中，通常使用迁移工具来管理数据库结构
	db.AutoMigrate(&User{}, Product{}, Order{}, Category{}, Profile{})

	return db, nil
}

// CreateUser 创建一个新用户
//func CreateUser(db *gorm.DB, user *User) error {
//	result := db.Create(user)
//	return result.Error
//}

// UpdateUser 更新一个已存在的用户
//func UpdateUser(db *gorm.DB, user *User) error {
//	result := db.Model(&user).Updates(user)
//	return result.Error
//}
