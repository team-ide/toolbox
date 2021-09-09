package main

import (
	"zookeeper"
)

func main() {
	server := ""
	zkService, err := zookeeper.CreateZKService(server)
	println("连接Zookeeper：", server)
	if err != nil {
		panic(err)
	}
	println(zkService)
}
