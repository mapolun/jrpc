package common

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

// FnGetDate 获取日期
func FnGetDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// FnGetClientIp 获取客户端IP
func FnGetClientIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("XRealIP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("XForwardedFor"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

// FnBtox 二进制转十六进制
func Fn2to16(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 16)
}

// FnXtob 十六进制转二进制
func Fn16to2(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}
