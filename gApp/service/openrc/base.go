package openrc

import (
	"github.com/gek64/gek/gExec"
	"os"
	"os/exec"
)

// https://wiki.gentoo.org/wiki/OpenRC_to_systemd_Cheatsheet

var (
	ServiceLocation = "/etc/init.d/"
)

// Load 开启服务自启+启动服务
func Load(serviceName string) (err error) {
	// 开启服务自启
	err = gExec.Run(exec.Command("rc-update", "add", serviceName))
	if err != nil {
		return err
	}

	// 启动服务
	return gExec.Run(exec.Command("rc-service", serviceName, "start"))
}

// Unload 关闭服务自启+停止服务
func Unload(serviceName string) (err error) {
	// 停止服务
	err = gExec.Run(exec.Command("rc-service", serviceName, "stop"))
	if err != nil {
		return err
	}

	// 关闭服务自启
	return gExec.Run(exec.Command("rc-update", "del", serviceName))
}

// Reload 重载服务
func Reload(serviceName string) (err error) {
	// 重启服务
	return gExec.Run(exec.Command("rc-service", serviceName, "restart"))
}

// Status 查看服务状态
func Status(serviceName string) (err error) {
	// 查看服务状态
	return gExec.Run(exec.Command("rc-service", serviceName, "status"))
}

// Uninstall 卸载服务
func Uninstall(serviceName string) (err error) {
	return os.RemoveAll(ServiceLocation + serviceName)
}
