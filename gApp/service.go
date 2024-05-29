package gApp

import (
	"fmt"
	"github.com/gek64/gek/gApp/service"
	"github.com/gek64/gek/gApp/service/openrc"
	"github.com/gek64/gek/gApp/service/rcd"
	"github.com/gek64/gek/gApp/service/systemd"
	"os"
)

type Service struct {
	service.Service
}

// NewService 新建服务
func NewService(initSystem string, serviceName string, serviceContent []byte) (*Service, error) {
	var sv service.Service
	switch initSystem {
	case "systemd":
		sv = systemd.NewService(serviceName, serviceContent)
	case "openrc":
		sv = openrc.NewService(serviceName, serviceContent)
	case "rc.d":
		sv = rcd.NewService(serviceName, serviceContent)
	default:
		return nil, fmt.Errorf("no supported init system found")
	}
	return &Service{Service: sv}, nil
}

// NewServiceFromFile 新建服务(从文件)
func NewServiceFromFile(initSystem string, serviceName string, serviceFile string) (*Service, error) {
	bytes, err := os.ReadFile(serviceFile)
	if err != nil {
		return nil, err
	}
	return NewService(initSystem, serviceName, bytes)
}

// Install 安装服务,不执行Load
func (s *Service) Install() (err error) {
	return s.Service.Install()
}

// Uninstall 卸载服务,执行Unload
func (s *Service) Uninstall() (err error) {
	return s.Service.Uninstall()
}

// Reload 重载服务
func (s *Service) Reload() (err error) {
	return s.Service.Reload()
}

// Load 开启服务自启+启动服务
func (s *Service) Load() (err error) {
	return s.Service.Load()
}

// Unload 关闭服务自启+停止服务
func (s *Service) Unload() (err error) {
	return s.Service.Unload()
}

// Status 查看服务状态
func (s *Service) Status() (err error) {
	return s.Service.Status()
}
