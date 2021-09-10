package main

import (
	"base"
	"config"
	"database_"
	"elasticsearch"
	"kafka"
	"mongo"
	"redis"
	"web"
	"worker"
	"zookeeper"
)

func init() {
	base.Init()
	config.Init()
	database_.Init()
	elasticsearch.Init()
	kafka.Init()
	mongo.Init()
	redis.Init()
	web.Init()
	worker.Init()
	zookeeper.Init()
}
func main() {
	web.StartServer()
	// server := ""
	// zkService, err := zookeeper.CreateZKService(server)
	// println("连接Zookeeper：", server)
	// if err != nil {
	// 	panic(err)
	// }
	// println(zkService)
}
