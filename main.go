package main

import (
    "fmt"
    "net"
)

func main() {
    addr := net.UDPAddr{
        Port: 10999,
        IP:   net.ParseIP("0.0.0.0"),
    }
    conn, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Println("启动服务器失败:", err)
        return
    }
    defer conn.Close()

    fmt.Println("服务器监听在10999端口")
    buf := make([]byte, 1024)
    for {
        n, remoteAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("读取数据时发生错误:", err)
            continue
        }
        fmt.Printf("接收到 %s 发送的消息: %s\n", remoteAddr, string(buf[:n]))
        response := "服务器已收到消息"
        conn.WriteToUDP([]byte(response), remoteAddr)
    }
}
