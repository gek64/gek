package systemd

import (
	"fmt"
	"github.com/gek64/gek/gFile"
	"os"
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
	// 检查服务文件夹是否存在
	_, err = os.Stat(ServiceLocation)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ServiceLocation, 0755)
		if err != nil {
			return err
		}
	}

	// 检查服务文件是否存在
	_, err = os.Stat(ServiceLocation + s.Name)
	if os.IsExist(err) {
		return fmt.Errorf("gek_service %s is already installed", s.Name)
	}

	_, err = gFile.CreateFile(ServiceLocation+s.Name, s.Content)
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
	return Uninstall(s.Name)
}

// Load 开启服务自启+启动服务
func (s *Service) Load() (err error) {
	return Load(s.Name)
}

// Unload 关闭服务自启+停止服务
func (s *Service) Unload() (err error) {
	return Unload(s.Name)
}

// Reload 重载服务
func (s *Service) Reload() (err error) {
	return Reload(s.Name)
}

// Status 查看服务状态,返回错误信息为错误的Code 或者 nil
// Code 代表含义查询 https://www.freedesktop.org/software/systemd/man/systemctl.html#Exit%20status
func (s *Service) Status() (returnCode error) {
	return Status(s.Name)
}
