package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 假设以下是用户输入的密码明文
	password := "password123"

	// 生成盐值
	salt := generateSalt()

	// 将密码明文和盐值组合，进行哈希加密
	hashedPassword1, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		// 哈希密码生成失败
		fmt.Println("哈希密码生成失败:", err)
		return
	}
	hashedPassword2, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 哈希密码生成失败
		fmt.Println("哈希密码生成失败:", err)
		return
	}
	fmt.Println("byte", hashedPassword1)
	fmt.Println("哈希密码加盐后:", string(hashedPassword1))
	fmt.Println("哈希密码不加盐:", string(hashedPassword2))
	fmt.Println("byte", hashedPassword2)

}

// 生成盐值
func generateSalt() string {
	// 在实际应用中，盐值应该是随机生成的，且每个用户应该有唯一的盐值
	// 这里为了演示目的，使用固定的盐值
	return "somesalt"
}
