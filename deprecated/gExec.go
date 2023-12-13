package deprecated

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Exist 检查指定程序是否存在于环境变量或者可执行文件的同一目录中
func Exist(program string) (path string, err error) {
	// 获取当前运行的可执行文件的目录
	exd, err := FindLocation()
	if err != nil {
		return "", err
	}
	// 当前运行的可执行文件的目录是否存在program
	exp := filepath.Join(exd, program)
	path, err = exec.LookPath(exp)
	if err != nil {
		// 当拼接路径不存在时,则从环境路径中查找program
		path, err = exec.LookPath(program)
		//当拼接路径及环境路径中均找不到program时返回false 空字符串 错误
		if err != nil {
			return "", err
		}
	}
	return path, nil
}

// FindLocation 返回当前运行的可执行文件的地址
// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
func FindLocation() (string, error) {
	// 当前可执行文件地址
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 截取可执行文件地址的文件夹地址部分
	ex = filepath.Dir(ex)
	return ex, err
}

// StringToCmd 字符串按空格分词转换为 *exec.Cmd
func StringToCmd(cmdString string) (cmd *exec.Cmd) {
	cmdArgs := strings.Fields(cmdString)
	return exec.Command(cmdArgs[0], cmdArgs[1:]...)
}
