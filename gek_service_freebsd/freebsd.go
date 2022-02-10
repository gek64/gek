package gek_service_freebsd

import (
	"fmt"
	"gek_exec"
	"os"
	"os/exec"
)

var (
	ServiceLocation = "/usr/local/etc/rc.d/"
)

// Load 开启服务自启+启动服务
func Load(serviceName string) (err error) {
	// 开启服务自启
	err = gek_exec.Run(exec.Command("service", serviceName, "enable"))
	if err != nil {
		return err
	}

	// 启动服务
	err = gek_exec.Run(exec.Command("service", serviceName, "start"))
	if err != nil {
		return err
	}

	return nil
}

// Unload 关闭服务自启+停止服务
func Unload(serviceName string) (err error) {
	// 停止服务
	err = gek_exec.Run(exec.Command("service", serviceName, "stop"))
	if err != nil {
		return err
	}

	// 关闭服务自启
	err = gek_exec.Run(exec.Command("service", serviceName, "delete"))
	if err != nil {
		return err
	}

	return nil
}

// Reload 重载服务
func Reload(serviceName string) (err error) {
	// 重启服务
	err = gek_exec.Run(exec.Command("service", serviceName, "restart"))
	if err != nil {
		return err
	}

	return nil
}

// Status 查看服务状态
func Status(serviceName string) (err error) {
	// 查看服务状态
	err = gek_exec.Run(exec.Command("service", serviceName, "status"))
	if err != nil {
		return err
	}

	return nil
}

// Uninstall 卸载服务
func Uninstall(serviceName string) (err error) {
	// 检查服务是否已经存在,不存在则返回错误
	_, err = os.Stat(ServiceLocation + serviceName)
	if os.IsNotExist(err) {
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
