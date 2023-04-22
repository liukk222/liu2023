package main // 声明包名为main，表示这是一个可执行程序

import (
	"fmt" // 导入fmt包，用于格式化输出
	"net" // 导入net包，用于网络编程
)

var (
	clients    = make(map[net.Addr]net.Conn) // 定义一个全局变量clients，用于存储所有连接到服务器的客户端，键是客户端的地址，值是客户端的连接对象
	addCh      = make(chan net.Conn)         // 定义一个全局变量addCh，用于向广播器发送新连接的客户端
	delCh      = make(chan net.Addr)         // 定义一个全局变量delCh，用于向广播器发送断开连接的客户端的地址
	messageCh  = make(chan []byte)           // 定义一个全局变量messageCh，用于向广播器发送客户端发送的消息
	listenAddr = "localhost:8080"            // 定义一个全局变量listenAddr，用于指定服务器监听的地址和端口
)

func main() { // 定义main函数，程序的入口点
	fmt.Println("Server started on", listenAddr)   // 打印一条日志，表示服务器启动了
	listener, err := net.Listen("tcp", listenAddr) // 调用net.Listen函数，创建一个TCP监听器，返回一个listener对象和一个错误对象
	if err != nil {                                // 如果发生了错误
		fmt.Println(err) // 打印错误信息
		return           // 退出程序
	}
	defer listener.Close() // 延迟执行listener.Close()函数，关闭监听器

	go broadcaster() // 启动一个新的goroutine（轻量级线程），运行broadcaster函数，用于处理所有客户端之间的通信

	for { // 无限循环
		conn, err := listener.Accept() // 调用listener.Accept()函数，等待并接受新的客户端连接，返回一个conn对象和一个错误对象
		if err != nil {                // 如果发生了错误
			fmt.Println(err) // 打印错误信息
			continue         // 跳过本次循环，继续下一次循环
		}

		addCh <- conn // 将新连接的客户端发送到addCh通道

		go handleConn(conn) // 启动一个新的goroutine，运行handleConn函数，用于处理该客户端的消息
	}
}

func broadcaster() { // 定义broadcaster函数，用于处理所有客户端之间的通信
	for { // 无限循环
		select { // 多路选择语句，根据不同的通道事件执行不同的分支
		case conn := <-addCh: // 如果从addCh通道接收到了新连接的客户端
			clients[conn.RemoteAddr()] = conn             // 将该客户端添加到clients映射中，键是客户端的地址，值是客户端的连接对象
			fmt.Println("New client:", conn.RemoteAddr()) // 打印一条日志，表示有新客户端连接了
		case addr := <-delCh: // 如果从delCh通道接收到了断开连接的客户端的地址
			delete(clients, addr)                     // 将该客户端从clients映射中删除
			fmt.Println("Client disconnected:", addr) // 打印一条日志，表示有客户端断开了连接
		case msg := <-messageCh: // 如果从messageCh通道接收到了客户端发送的消息
			for _, conn := range clients { // 遍历所有的客户端
				_, err := conn.Write(msg) // 将消息写入到客户端的连接中
				if err != nil {           // 如果发生了错误
					fmt.Println(err) // 打印错误信息
				}
			}
		}
	}
}

func handleConn(conn net.Conn) { // 定义handleConn函数，用于处理客户端的消息
	defer func() { // 延迟执行一个匿名函数
		delCh <- conn.RemoteAddr() // 将断开连接的客户端的地址发送到delCh通道
		conn.Close()               // 关闭客户端的连接
	}()

	for { // 无限循环
		msg := make([]byte, 4096) // 创建一个字节切片，用于存储客户端发送的消息
		n, err := conn.Read(msg)  // 从客户端的连接中读取数据，返回读取到的字节数和一个错误对象
		if err != nil {           // 如果发生了错误
			return // 退出函数
		}

		messageCh <- msg[:n] // 将读取到的消息发送到messageCh通道
	}
}
