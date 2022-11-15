package gApp

import (
	"gek/gService/rc"
	"gek/gService/systemd"
	"os"
	"runtime"
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

// Install 安装服务
func (s Service) Install() (err error) {
	// 分系统运行不同的命令
	switch runtime.GOOS {
	case SupportedOS[0]:
		service := systemd.NewService(s.Name, s.Content)
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
	case SupportedOS[1]:
		service := rc.NewService(s.Name, s.Content)
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
	}
	return nil
}

// Uninstall 卸载服务
func (s Service) Uninstall() (err error) {
	// 分系统运行不同的命令
	switch runtime.GOOS {
	case SupportedOS[0]:
		service := systemd.NewService(s.Name, s.Content)
		// 卸载服务
		err = service.Uninstall()
		if err != nil {
			return err
		}
	case SupportedOS[1]:
		service := rc.NewService(s.Name, s.Content)
		// 卸载服务
		err = service.Uninstall()
		if err != nil {
			return err
		}
	}
	return nil
}

// Restart 重启服务
func (s Service) Restart() (err error) {
	// 分系统运行不同的命令
	switch runtime.GOOS {
	case SupportedOS[0]:
		service := systemd.NewService(s.Name, s.Content)
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
	case SupportedOS[1]:
		service := rc.NewService(s.Name, s.Content)
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
	}
	return nil
}
