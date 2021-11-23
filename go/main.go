package main

import (
	"base"
	"config"
	"database_"
	"elasticsearch"
	"kafka"
	"mongo"
	"redis"
	"resource"
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
	zookeeper.Init()
	resource.Init()
	web.Init()
	worker.Init()
}
func main() {
	web.StartServer()
}
