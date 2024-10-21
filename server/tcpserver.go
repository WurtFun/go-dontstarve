// server/tcp.go
package server

import (
	"fmt"
	"net"
)

// 启动 TCP 服务器
func StartTCPServer() {
	listener, err := net.Listen("tcp", ":10999")
	if err != nil {
		fmt.Println("无法启动 TCP 服务器:", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCP 服务器正在监听 10999 端口...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("无法接受连接:", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取数据时发生错误:", err)
		return
	}

	fmt.Printf("接收到客户端的数据: %x\n", buf[:n])
	response := "连接确认成功！"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("发送响应时发生错误:", err)
	}

	fmt.Printf("已向客户端发送响应: %s\n", response)
}
