package gToolbox

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckToolbox 检查工具链是否完整,不完整会返回带有缺少的工具链的错误
func CheckToolbox(toolbox []string) error {
	var miss []string

	// 检查工具链,如果有不存在的会写入message
	for _, tool := range toolbox {
		_, err := exec.LookPath(tool)
		if err != nil {
			miss = append(miss, tool)
		}
	}

	if len(miss) != 0 {
		return fmt.Errorf("can not find %s", miss)
	}

	return nil
}

// CheckRoot 检查是否使用root权限
func CheckRoot() bool {
	return os.Geteuid() == 0
}
