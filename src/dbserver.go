package src

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	var dsn string = "root:123456@tcp(121.40.117.237:3306)/SpotProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	// 自动迁移模式，根据模型创建表
	// 注意：在生产环境中，通常使用迁移工具来管理数据库结构
	db.AutoMigrate(&User{})

	return db, nil
}

func CreateUser(usertable *User) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	DB.Create(&usertable)
	return nil
}

func GetUserByName(username string) (*User, error) {
	var user User
	DB, err := InitDB()
	if err != nil {
		return nil, err
	}
	DB.Find(&user, "username=?", username)
	return &user, nil // 用户存在，返回用户指针和nil错误
}

func UpdateUser(user *User) error {
	DB, err := InitDB()
	if err != nil {
		return err
	}
	DB.Table("users").Where("username = ?", user.Username).Updates(user)
	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
	return nil
}

func DeleteByName(db *gorm.DB, username string) error {
	var user User
	result := db.Delete(user, "username = ?", username)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "record not found") {
			return errors.New("用户不存在") // 用户不存在，返回自定义错误
		}
		return result.Error
	}
	return nil
}
