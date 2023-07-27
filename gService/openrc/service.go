package openrc

import (
    "fmt"
    "github.com/gek64/gek/gFile"
    "os"
    "path/filepath"
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
func (s Service) Install() (err error) {
    // 检查服务文件夹是否存在
    _, err = os.Stat(ServiceLocation)
    if os.IsNotExist(err) {
        err = os.MkdirAll(ServiceLocation, 0755)
        if err != nil {
            return err
        }
    }

    // 检查服务文件是否存在
    _, err = os.Stat(filepath.Join(ServiceLocation, s.Name))
    if os.IsExist(err) {
        return fmt.Errorf("gek_service %s is already installed", s.Name)
    }

    // 创建服务文件
    _, err = gFile.CreateFile(filepath.Join(ServiceLocation, s.Name), s.Content)
    if err != nil {
        return err
    }

    // 服务文件赋权755
    err = os.Chmod(filepath.Join(ServiceLocation, s.Name), 0755)
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
func (s Service) Uninstall() (err error) {
    return Uninstall(s.Name)
}

// Load 开启服务自启+启动服务
func (s Service) Load() (err error) {
    return Load(s.Name)
}

// Unload 关闭服务自启+停止服务
func (s Service) Unload() (err error) {
    return Unload(s.Name)
}

// Reload 重载服务
func (s Service) Reload() (err error) {
    return Reload(s.Name)
}

// Status 查看服务状态
func (s Service) Status() (err error) {
    return Status(s.Name)
}
