package src

import (
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// 创建一个新用户实例进行测试
	testUser := &User{
		Username: "hetao",
		Password: "03220410Az",
		Email:    "15180190056@example.com",
		Phone:    "456asda",
	}
	// 调用CreateUser函数并检查错误
	CreateUser(testUser)

}

// TestGetUserByID 测试GetUserByID函数
func TestGetUserByID(t *testing.T) {
	// 使用db进行测试
	user, err := GetUserByName("tom")
	if err != nil {
		t.Errorf("Failed to get user by ID: %v", err)
	}
	fmt.Println(user.Username)
}

func TestUpdateUser(t *testing.T) {
	// 创建一个新用户实例进行测试
	testUser := &User{
		Username: "tom",
		Password: "03220410Az",
		Email:    "15180190056@example.com",
		Phone:    "456asda",
	}
	// 调用CreateUser函数并检查错误
	UpdateUser(testUser)
}

func TestDeleteByName(t *testing.T) {
	db, err := InitDB() // 假设InitDB是您的数据库初始化函数
	if err != nil {
		t.Fatalf("初始化数据库以供测试失败: %v", err)
	}

	// 假设我们有一个函数可以检查用户名是否存在
	exists, err := UserExists(db, "hetao")
	if err != nil {
		t.Fatalf("检查用户是否存在时出错: %v", err)
	}
	if !exists {
		t.Skip("用户'hetao'不存在，跳过删除测试")
	}
	c
	// 尝试删除用户名为"hetao"的用户
	err = DeleteByName(db, "hetao")
	if err != nil {
		t.Errorf("删除用户名为'hetao'的用户失败: %v", err)
	}

	// 验证用户是否已被删除
	exists, err = UserExists(db, "hetao")
	if err != nil {
		t.Errorf("验证用户是否被删除时出错: %v", err)
	}
	if exists {
		t.Error("用户'hetao'应该已被删除，但仍然存在")
	}
}
