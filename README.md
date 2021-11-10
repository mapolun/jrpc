# jrcp
jsonrpc服务端轻量框架，支持tcp和http协议。
## 安装
```
go get -u github.com/mapolun/jrpc
```
## 开始
```go
package main

import (
	"jrpc/lib"
	"jrpc/service"
)

func main() {
	//http协议监听端口
	s, _ := lib.NewServer("http", "34712")
	//tcp协议监听端口
	//s, _ := lib.NewServer("tcp", "34713")
	//注册消息服务
	s.Register(new(service.News))
	//其他功能服务
	//s.Register(new(service.Other))
	//启动
	s.Start()
}
```
## 启动
```
go run main.go
```
