package main

import (
	"jrpc/lib"
	"jrpc/service"
)

func main() {
	//监听端口
	s, _ := lib.NewServer("http", "34712")
	//注册消息服务
	s.Register(new(service.News))
	//todo 其他服务
	//s.Register(new(service.Other))
	//启动
	s.Start()
}
