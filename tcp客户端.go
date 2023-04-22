package main // 声明这个文件属于main包，也就是可执行程序的入口

import ( // 导入需要用到的标准库包
	"fmt"     // 提供格式化输入输出的功能
	"net"     // 提供网络编程的功能
	"os"      // 提供操作系统相关的功能
	"strings" // 提供字符串处理的功能
)

var ( // 声明一个全局变量，用来存储服务器的地址
	serverAddr = "localhost:8080"
)

func main() { // 定义main函数，程序的执行从这里开始
	conn, err := net.Dial("tcp", serverAddr) // 使用net包的Dial函数，尝试建立一个TCP连接到服务器，返回一个连接对象和一个错误值
	if err != nil {                          // 如果发生错误，打印错误信息并返回
		fmt.Println(err)
		return
	}
	defer conn.Close() // 使用defer语句，确保在函数结束时关闭连接

	go func() { // 使用go关键字，启动一个新的协程（轻量级线程），用来接收服务器发来的消息
		for { // 无限循环
			msg := make([]byte, 4096) // 创建一个字节切片，用来存储消息内容，长度为4096字节
			n, err := conn.Read(msg)  // 使用连接对象的Read方法，从连接中读取数据，返回读取到的字节数和错误值
			if err != nil {           // 如果发生错误，打印错误信息并退出程序
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Println(strings.TrimSpace(string(msg[:n]))) // 将读取到的字节切片转换为字符串，并去掉两端的空白字符，然后打印出来
		}
	}()

	for { // 无限循环
		var msg string      // 声明一个字符串变量，用来存储用户输入的消息
		fmt.Scanln(&msg)    // 使用fmt包的Scanln函数，从标准输入读取一行文本，并赋值给msg变量
		if msg == "/quit" { // 如果用户输入/quit，表示要退出程序
			return // 返回，结束main函数
		}

		_, err := conn.Write([]byte(msg)) // 将用户输入的消息转换为字节切片，并使用连接对象的Write方法，将数据写入连接中，返回写入的字节数和错误值
		if err != nil {                   // 如果发生错误，打印错误信息并返回
			fmt.Println(err)
			return
		}
	}
}
