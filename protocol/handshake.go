package protocol

import (
	"fmt"
	"net"
)

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取数据时发生错误:", err)
		return
	}

	fmt.Printf("接收到客户端的数据: %x\n", buf[:n])

	if n > 0 {
		response := "服务器已收到请求，连接成功！"

		responseBytes := []byte{
			0x01, 0x00, 0x00, 0x00,
			0x02, 0x01, 0x00, 0x00,
		}

		_, err = conn.Write(responseBytes)
		if err != nil {
			fmt.Println("发送响应时发生错误:", err)
		}

		fmt.Printf("已向客户端发送响应: %x\n", responseBytes)
	}
}
