package main

import (
	"github.com/wonderivan/logger"
	"jrpc/lib"
	"jrpc/service"
	"log"
)

func main() {
	// 加载写入日志文件系统
	err := logger.SetLogger("./lib/config/logger.json")
	if err != nil {
		log.Panic("日志服务启动失败")
	}
	//监听端口
	s, _ := lib.NewServer("http", "192.168.2.160", "34712")
	//注册消息服务
	s.Register(new(service.News))
	//todo 其他服务
	//s.Register(new(service.Other))
	//启动
	s.Start()
}
