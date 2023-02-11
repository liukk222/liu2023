package main

import (
	"bufio"
	"fmt"
	"os"
)

func timingSafeEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result int
	for i := 0; i < len(a); i++ {
		result |= int(a[i] ^ b[i])
	}
	return result == 0
}

func main() {
	// 设置云端的密码
	cloudPassword := []byte("secret")

	// 读取用户输入的密码
	fmt.Print("请输入密码：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadBytes('\n')

	// 比对密码
	if timingSafeEqual(cloudPassword, input) {
		fmt.Println("密码正确")
	} else {
		fmt.Println("密码错误")
	}
}
