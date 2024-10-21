package main

import (
	"go-dontstarve/server"
)

func main() {
	go server.StartTCPServer()
	go server.StartUDPServer()
	select {}
}
 
