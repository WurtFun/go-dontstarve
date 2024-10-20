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
        fmt.Println("服务器启动失败:", err)
        return
    }
    defer conn.Close()

    fmt.Println("服务器已在10999端口开放")

    buf := make([]byte, 1024)

    for {
        n, remoteAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("读取数据时发生错误:", err)
            continue
        }
        //返回值疑似加密
        fmt.Printf("接收到客户端 (%s) 的消息: %x\n", remoteAddr, buf[:n])
        joinRequest := string(buf[:n])
        fmt.Printf("客户端申请加入: %s\n", joinRequest)

        response := null
        _, err = conn.WriteToUDP([]byte(response), remoteAddr)
        if err != nil {
            fmt.Println("发送响应时发生错误:", err)
            continue
        }

        fmt.Printf("已向客户端 (%s) 发送响应: %s\n", remoteAddr, response)
    }
}
