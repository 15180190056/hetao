package src

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义处理POST请求的handler函数
//func Userlogin(c *gin.Context) {
//	// 创建一个LoginData实例
//	var loginData LoginData
//
//	// 绑定JSON数据到结构体
//	if err := c.ShouldBindJSON(&loginData); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// 在这里处理登录逻辑，例如验证用户名和密码
//	var userdata = LoginData{
//		Username: "gin_start",
//		Password: "123456",
//	}
//
//	if loginData == userdata {
//		// 返回响应
//		c.JSON(http.StatusOK, gin.H{"status": "success"})
//	}
//}

// 注册新用户
func RegisterUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB, err := InitDB()
	if err != nil {
		// 处理错误
	}

	if err := CreateUser(DB, &user); err != nil {
		// 处理错误
	}

}

// 对密码进行哈希处理
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

// 用户登录
func LoginUser(username, password string) (*User, error) {
	user, exists := users[username]
	if !exists {
		return nil, fmt.Errorf("用户名 %s 不存在", username)
	}

	if !verifyPassword(password, user.Password) {
		return nil, fmt.Errorf("密码错误")
	}

	return user, nil
}

// 更新用户信息
func UpdateUser(user *User, newUsername string, newEmail string) error {
	if newUsername != "" {
		if _, exists := users[newUsername]; exists {
			return fmt.Errorf("新用户名 %s 已被注册", newUsername)
		}
		user.Username = newUsername
	}

	if newEmail != "" {
		user.Email = newEmail
	}

	fmt.Printf("用户 %s 的信息已更新\n", user.Username)
	return nil
}

// 删除用户
func DeleteUser(username string) error {
	delete(users, username)
	fmt.Printf("用户 %s 已删除\n", username)
	return nil
}

// 打印所有用户信息
func PrintAllUsers() {
	for _, user := range users {
		fmt.Printf("用户ID: %d, 用户名: %s, 邮箱: %s\n", user.ID, user.Username, user.Email)
	}
}

//func main() {
//	// 注册新用户
//	err := RegisterUser("alice", "mypassword123", "alice@example.com")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// 登录用户
//	user, err := LoginUser("alice", "mypassword123")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("登录成功，欢迎 %s\n", user.Username)
//	}
//
//	// 更新用户信息
//	err = UpdateUser(user, "alice_new", "alice_new@example.com")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// 打印所有用户信息
//	PrintAllUsers()
//
//	// 删除用户
//	err = DeleteUser("alice_new")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// 再次打印所有用户信息以验证用户是否已被删除
