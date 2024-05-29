package systemd

import (
	"github.com/gek64/gek/gExec"
	"os"
	"os/exec"
)

var (
	ServiceLocation = "/etc/systemd/system/"
)

// Load 开启服务自启+启动服务
func Load(serviceName string) (err error) {
	// 重载所有服务
	err = gExec.Run(exec.Command("systemctl", "daemon-reload"))
	if err != nil {
		return err
	}

	// 开启服务自启
	err = gExec.Run(exec.Command("systemctl", "enable", serviceName))
	if err != nil {
		return err
	}

	// 启动服务
	return gExec.Run(exec.Command("systemctl", "start", serviceName))
}

// Unload 关闭服务自启+停止服务
func Unload(serviceName string) (err error) {
	// 重载所有服务
	err = gExec.Run(exec.Command("systemctl", "daemon-reload"))
	if err != nil {
		return err
	}

	// 停止服务
	err = gExec.Run(exec.Command("systemctl", "stop", serviceName))
	if err != nil {
		return err
	}

	// 关闭服务自启
	return gExec.Run(exec.Command("systemctl", "disable", serviceName))
}

// Reload 重载服务
func Reload(serviceName string) (err error) {
	// 重载所有服务
	err = gExec.Run(exec.Command("systemctl", "daemon-reload"))
	if err != nil {
		return err
	}

	// 重启服务
	return gExec.Run(exec.Command("systemctl", "restart", serviceName))
}

// Status 查看服务状态,返回错误信息为错误的Code 或者 nil
// Code 代表含义查询 https://www.freedesktop.org/software/systemd/man/systemctl.html#Exit%20status
func Status(serviceName string) (returnCode error) {
	// 查看服务状态
	return gExec.Run(exec.Command("systemctl", "status", serviceName, "--no-pager"))
}

// Uninstall 卸载服务
func Uninstall(serviceName string) (err error) {
	return os.RemoveAll(ServiceLocation + serviceName)
}
