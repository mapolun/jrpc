package tcp

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type moduleServer struct {
	Listener net.Listener
	Server   rpc.Server
}

// NewServer 监听
func NewServer(port string) *moduleServer {
	//开始监听
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	log.Println(fmt.Sprintf("listen at :%v", port))
	if err != nil {
		log.Panic("server \t listen error:", err.Error())
	}
	return &moduleServer{
		Listener: listener,
	}
}

// Register 注册
func (m *moduleServer) Register(i interface{}) {
	err := m.Server.Register(i)
	if err != nil {
		log.Panic("server \t listen register:", err.Error())
	}
}

// Start 启动
func (m *moduleServer) Start() {
	for {
		conn, err := m.Listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		// 在goroutine中处理请求
		// 绑定rpc的编码器，使用http connection新建一个jsonrpc编码器，并将该编码器绑定给http处理器
		go m.Server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
