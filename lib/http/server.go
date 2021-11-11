package http

import (
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"jrpc/lib/common"
	"log"
	"net/http"
	"sync"
)

type moduleServer struct {
	Server common.Server
	Port   string
	Ip     string
}

// NewServer 监听
func NewServer(ip string, port string) *moduleServer {
	//开始监听
	return &moduleServer{
		common.Server{
			sync.Map{},
			common.Hooks{},
			nil,
		},
		port,
		ip,
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
	var url = fmt.Sprintf("%v:%v", m.Ip, m.Port)
	log.Printf("Listening http://%v:%v", m.Ip, m.Port)
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

	logger.Info("来源IP", common.FnGetClientIp(r))
	logger.Info("请求头部", r.Header)
	logger.Info("请求参数", string(data))
	res := m.Server.Handler(data)
	logger.Info("响应参数", string(res))
	w.Write(res)
}
