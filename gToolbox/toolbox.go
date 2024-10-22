package gToolbox

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckToolbox 检查工具链是否完整,不完整会返回带有缺少的工具链的错误
func CheckToolbox(toolbox []string) (missItems []string, err error) {
	// 检查工具链,如果有不存在的会写入message
	for _, tool := range toolbox {
		_, err = exec.LookPath(tool)
		if err != nil {
			missItems = append(missItems, tool)
		}
	}
	if len(missItems) != 0 {
		err = fmt.Errorf("can not find %s", missItems)
	}
	return missItems, err
}

// CheckRoot 检查是否使用root权限
func CheckRoot() bool {
	return os.Geteuid() == 0
}
