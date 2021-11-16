package gek_service

import (
	"fmt"
	"gek_exec"
	"gek_file"
	"os"
)

var (
	ServiceToolbox  = []string{"systemd", "systemctl"}
	ServiceLocation = "/etc/systemd/system/"
)

type Service struct {
	Name    string
	Content string
}

func NewService(name string, content string) (s Service) {
	s.Name = name
	s.Content = content
	return s
}

// Install 安装服务文件 并Load
func (s *Service) Install() (err error) {
	exist, isDir, err := gek_file.Exist(ServiceLocation + s.Name)
	if err == nil || isDir || exist {
		return fmt.Errorf("gek_service %s is already installed", s.Name)
	}

	_, err = gek_file.CreateFile(ServiceLocation+s.Name, s.Content)
	if err != nil {
		return err
	}

	err = s.Load()
	if err != nil {
		return err
	}

	return nil
}

// Uninstall 卸载服务
func (s *Service) Uninstall() (err error) {
	err = Uninstall(s.Name)
	if err != nil {
		return err
	}
	return nil
}

// Load 开启服务自启+启动服务
func (s *Service) Load() (err error) {
	err = Load(s.Name)
	if err != nil {
		return err
	}
	return nil
}

// Unload 关闭服务自启+停止服务
func (s *Service) Unload() (err error) {
	err = Unload(s.Name)
	if err != nil {
		return err
	}
	return nil
}

// Reload 重载服务
func (s *Service) Reload() (err error) {
	err = Reload(s.Name)
	if err != nil {
		return err
	}
	return nil
}

// Status 查看服务状态,返回错误信息为错误的Code 或者 nil
// Code 代表含义查询 https://www.freedesktop.org/software/systemd/man/systemctl.html#Exit%20status
func (s *Service) Status() (returnCode error) {
	// return code https://www.freedesktop.org/software/systemd/man/systemctl.html#Exit%20status
	returnCode = Status(s.Name)
	return returnCode
}

// 通用函数
// 适用于无需使用Install函数的服务

// Uninstall 卸载服务
func Uninstall(serviceName string) (err error) {
	// 检查服务是否已经存在
	exist, _, err := gek_file.Exist(ServiceLocation + serviceName)
	if err != nil {
		return err
	}

	// 不存在则返回错误
	if !exist {
		return fmt.Errorf("gek_service %s is not installed", serviceName)
	}

	// 关闭服务自启+停止服务
	err = Unload(serviceName)
	if err != nil {
		return err
	}

	// 删除服务文件
	err = os.RemoveAll(ServiceLocation + serviceName)
	if err != nil {
		return err
	}

	return nil
}

// Load 开启服务自启+启动服务
func Load(serviceName string) (err error) {
	// 重载所有服务
	err = gek_exec.Run("systemctl daemon-reload")
	if err != nil {
		return err
	}

	// 启动服务
	err = gek_exec.Run("systemctl start " + serviceName)
	if err != nil {
		return err
	}

	// 开启服务自启
	err = gek_exec.Run("systemctl enable " + serviceName)
	if err != nil {
		return err
	}

	return nil
}

// Unload 关闭服务自启+停止服务
func Unload(serviceName string) (err error) {
	// 停止服务
	err = gek_exec.Run("systemctl stop " + serviceName)
	if err != nil {
		return err
	}

	// 关闭服务自启
	err = gek_exec.Run("systemctl disable " + serviceName)
	if err != nil {
		return err
	}

	// 重载所有服务
	err = gek_exec.Run("systemctl daemon-reload")
	if err != nil {
		return err
	}

	return nil
}

// Reload 重载服务
func Reload(serviceName string) (err error) {
	// 重载所有服务
	err = gek_exec.Run("systemctl daemon-reload")
	if err != nil {
		return err
	}

	// 重启服务
	err = gek_exec.Run("systemctl restart " + serviceName)
	if err != nil {
		return err
	}

	return nil
}

// Status 查看服务状态,返回错误信息为错误的Code 或者 nil
// Code 代表含义查询 https://www.freedesktop.org/software/systemd/man/systemctl.html#Exit%20status
func Status(serviceName string) (returnCode error) {
	// 查看服务状态
	returnCode = gek_exec.Run("systemctl status " + serviceName + " --no-pager")
	return returnCode
}
