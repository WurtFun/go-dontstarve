// server/udp.go
package server

import (
	"fmt"
	"net"
)

// 启动 UDP 服务器
func StartUDPServer() {
	addr := net.UDPAddr{
		Port: 10998, // RakNet 的 UDP 端口
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("无法启动 UDP 服务器:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP 服务器正在监听 10998 端口...")

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("读取 UDP 数据时发生错误:", err)
			continue
		}

		fmt.Printf("接收到客户端 (%s) 的 UDP 消息: %x\n", remoteAddr, buf[:n])

		// 响应客户端的 UDP 消息
		response := "UDP 响应成功"
		_, err = conn.WriteToUDP([]byte(response), remoteAddr)
		if err != nil {
			fmt.Println("发送 UDP 响应时发生错误:", err)
		}
	}
}
