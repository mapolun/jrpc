package http

import (
	"fmt"
	"io/ioutil"
	"jrpc/lib/common"
	"log"
	"net/http"
	"sync"
)

type moduleServer struct {
	Server common.Server
	Port   string
}

// NewServer 监听
func NewServer(port string) *moduleServer {
	//开始监听
	return &moduleServer{
		common.Server{
			sync.Map{},
			common.Hooks{},
			nil,
		},
		port,
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", m.handleFunc)
	var url = fmt.Sprintf("%v:%v", "127.0.0.1", m.Port)
	log.Printf("Listening http://%v:%v", "127.0.0.1", m.Port)
	err := http.ListenAndServe(url, mux)
	if err != nil {
		log.Println("server \t 监听错误")
	}
}

func (m *moduleServer) handleFunc(w http.ResponseWriter, r *http.Request) {
	var (
		err  error
		data []byte
	)
	w.Header().Set("Content-Type", "application/json")
	if data, err = ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("请求时间：" + common.FnGetDate())
	fmt.Println("来源IP：" + common.FnGetClientIp(r))
	fmt.Println(string(data))
	res := m.Server.Handler(data)
	w.Write(res)
}
