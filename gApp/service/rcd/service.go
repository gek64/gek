package rcd

import (
	"os"
	"path/filepath"
)

type Service struct {
	Name    string
	Content []byte
}

func NewService(name string, content []byte) (s *Service) {
	return &Service{
		Name:    name,
		Content: content,
	}
}

// Install 安装服务文件
func (s *Service) Install() (err error) {
	// 检查服务文件夹是否存在
	_, err = os.Stat(ServiceLocation)
	if os.IsNotExist(err) {
		err = os.MkdirAll(ServiceLocation, 0755)
		if err != nil {
			return err
		}
	}

	// 创建服务文件
	err = os.WriteFile(filepath.Join(ServiceLocation, s.Name), s.Content, 0755)
	if err != nil {
		return err
	}

	// 服务文件赋权755
	return os.Chmod(filepath.Join(ServiceLocation, s.Name), 0755)
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

// Status 查看服务状态
func (s *Service) Status() (err error) {
	return Status(s.Name)
}
