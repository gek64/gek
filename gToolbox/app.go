package gToolbox

import (
	"fmt"
	"github.com/gek64/gek/gExec"
	"os"
)

// CheckToolbox 检查工具链是否完整,不完整会返回带有缺少的工具链的错误
func CheckToolbox(toolbox []string) error {
	var message = ""

	// 检查工具链,如果有不存在的会写入message
	for _, tool := range toolbox {
		exist, _, _ := gExec.Exist(tool)
		if !exist {
			message = message + " " + tool
		}
	}

	// 将message转换为错误格式
	if message != "" {
		return fmt.Errorf("can not find%s, install before running", message)
	}

	return nil
}

// CheckRoot 检查是否使用root权限
func CheckRoot() bool {
	if os.Geteuid() == 0 {
		return true
	}
	return false
}

// InstallStatus 检查应用与服务的安装状态
func InstallStatus(appLocation string, serviceLocation string) (bool, bool) {
	var appStatus = false
	var serviceStatus = false

	// 检查应用安装状态
	_, err := os.Stat(appLocation)
	if err == nil {
		appStatus = true
	}

	// 检查服务安装状态
	_, err = os.Stat(serviceLocation)
	if err == nil {
		serviceStatus = true
	}
	return appStatus, serviceStatus

}
