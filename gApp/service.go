package gApp

import (
	"fmt"
	"github.com/gek64/gek/gService"
	"github.com/gek64/gek/gService/openrc"
	"github.com/gek64/gek/gService/rcd"
	"github.com/gek64/gek/gService/systemd"
	"os"
	"os/exec"
)

type Service struct {
	Name    string
	Content string
}

// NewService 新建服务
func NewService(name string, content string) (s Service) {
	return Service{Name: name, Content: content}
}

// NewServiceFromFile 新建服务(从文件)
func NewServiceFromFile(name string, file string) (s Service, err error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return Service{}, err
	}
	return NewService(name, string(bytes)), nil
}

// CheckInitSystem 检查系统中的init system
func CheckInitSystem() (initSystemName string, initSystemBin string) {
	for key, value := range InitSystem {
		_, err := exec.LookPath(value)
		if err != nil {
			continue
		} else {
			return key, value
		}
	}
	return "", ""
}

// Install 安装服务
func (s Service) Install() (err error) {
	var service gService.Service

	// 检查系统中的init system
	_, initSystemBin := CheckInitSystem()

	// 分init system创建服务
	switch initSystemBin {
	case InitSystem["systemd"]:
		service = systemd.NewService(s.Name, s.Content)
	case InitSystem["openrc"]:
		service = openrc.NewService(s.Name, s.Content)
	case InitSystem["rc.d"]:
		service = rcd.NewService(s.Name, s.Content)
	default:
		var supportInitSystemListString string
		for key := range InitSystem {
			supportInitSystemListString = supportInitSystemListString + ", " + key
		}
		return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
	}

	// 安装服务
	err = service.Install()
	if err != nil {
		return err
	}

	// 查看服务状态
	err = service.Status()
	if err != nil {
		return err
	}

	return nil
}

// Uninstall 卸载服务
func (s Service) Uninstall() (err error) {
	var service gService.Service

	// 检查系统中的init system
	_, initSystemBin := CheckInitSystem()

	// 分init system创建服务
	switch initSystemBin {
	case InitSystem["systemd"]:
		service = systemd.NewService(s.Name, s.Content)
	case InitSystem["openrc"]:
		service = openrc.NewService(s.Name, s.Content)
	case InitSystem["rc.d"]:
		service = rcd.NewService(s.Name, s.Content)
	default:
		var supportInitSystemListString string
		for key := range InitSystem {
			supportInitSystemListString = supportInitSystemListString + ", " + key
		}
		return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
	}

	// 卸载服务
	err = service.Uninstall()
	if err != nil {
		return err
	}

	return nil
}

// Restart 重启服务
func (s Service) Restart() (err error) {
	var service gService.Service

	// 检查系统中的init system
	_, initSystemBin := CheckInitSystem()

	// 分init system创建服务
	switch initSystemBin {
	case InitSystem["systemd"]:
		service = systemd.NewService(s.Name, s.Content)
	case InitSystem["openrc"]:
		service = openrc.NewService(s.Name, s.Content)
	case InitSystem["rc.d"]:
		service = rcd.NewService(s.Name, s.Content)
	default:
		var supportInitSystemListString string
		for key := range InitSystem {
			supportInitSystemListString = supportInitSystemListString + ", " + key
		}
		return fmt.Errorf("no supported init system found, currently only %s are supported", supportInitSystemListString)
	}

	// 重载服务
	err = service.Reload()
	if err != nil {
		return err
	}
	// 查看服务状态
	err = service.Status()
	if err != nil {
		return err
	}

	return nil
}
