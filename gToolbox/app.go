package gToolbox

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckToolbox 检查工具链是否完整,不完整会返回带有缺少的工具链的错误
func CheckToolbox(toolbox []string) error {
	var message = ""

	// 检查工具链,如果有不存在的会写入message
	for _, tool := range toolbox {
		_, err := exec.LookPath(tool)
		if err != nil {
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
