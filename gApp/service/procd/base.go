package procd

import (
	"github.com/gek64/gek/gExec"
	"os"
	"os/exec"
)

// https://openwrt.org/docs/guide-user/base-system/managing_services

var (
	ServiceLocation = "/etc/init.d/"
)

// Load 开启服务自启+启动服务
func Load(serviceName string) (err error) {
	// 开启服务自启(openwrt 下使用 command.run() 会报错,使用 command.start() 代替)
	err = exec.Command("service", serviceName, "enable").Start()
	if err != nil {
		return err
	}

	// 启动服务
	return gExec.Run(exec.Command("service", serviceName, "start"))
}

// Unload 关闭服务自启+停止服务
func Unload(serviceName string) (err error) {
	// 停止服务
	err = gExec.Run(exec.Command("service", serviceName, "stop"))
	if err != nil {
		return err
	}

	// 关闭服务自启
	return gExec.Run(exec.Command("service", serviceName, "disable"))
}

// Reload 重载服务
func Reload(serviceName string) (err error) {
	// 重启服务
	return gExec.Run(exec.Command("service", serviceName, "restart"))
}

// Status 查看服务状态
func Status(serviceName string) (err error) {
	// 查看服务状态
	return gExec.Run(exec.Command("service", serviceName, "status"))
}

// Uninstall 卸载服务
func Uninstall(serviceName string) (err error) {
	return os.RemoveAll(ServiceLocation + serviceName)
}
