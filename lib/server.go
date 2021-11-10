package lib

import (
	"errors"
	"jrpc/lib/http"
	"jrpc/lib/tcp"
)

type ServerInterface interface {
	Start()
	Register(i interface{})
}

func NewServer(network string, port string) (ServerInterface, error) {
	var err error
	//简单工厂
	switch network {
	case "tcp":
		return tcp.NewServer(port), err
	case "http":
		return http.NewServer(port), err
	}
	return nil, errors.New("未匹配到可用网关")
}
