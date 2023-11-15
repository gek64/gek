package gNet

import (
	"net/http"
	"strings"
)

func GetIPFromRequest(req *http.Request) (ip string, err error) {
	// 查找 req.Header 中 X-Real-Ip
	ip = req.Header.Get("X-Real-Ip")
	// 查找 req.Header 中 X-Forwarded-For
	if ip == "" {
		ip = strings.Split(req.Header.Get("X-Forwarded-For"), ",")[0]
	}
	// 查找 req.RemoteAddr
	if ip == "" {
		ip = strings.Split(req.RemoteAddr, ":")[0]
	}
	return ip, err
}
