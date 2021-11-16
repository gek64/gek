package gek_toolbox

import (
	"fmt"
	"gek_exec"
	"gek_github"
	"os"
	"runtime"
)

var (
	// 系统架构来源 https://golang.org/doc/install/source#environment
	osList   = []string{"android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"}
	archList = []string{"amd64", "386", "arm", "arm64", "ppc64le", "mips64le", "mips64", "mipsle", "mips", "s390x", "wasm"}
)

// GetDownloadLink 获取Github应用对应系统/架构对的下载链接
func GetDownloadLink(repo string, appMap map[string]string) (string, error) {
	var downloadLink string = ""
	var keyPair string = ""

	// 获取Github API
	githubAPI, err := gek_github.NewGithubAPI(repo)
	if err != nil {
		return "", err
	}

	// 从当前 系统/架构 对中拼接 应用集合的 key pair
	for _, oss := range osList {
		for _, arch := range archList {
			if oss == runtime.GOOS && arch == runtime.GOARCH {
				keyPair = oss + "_" + arch
				break
			}
		}
		if keyPair != "" {
			break
		}
	}

	// 从应用集合中找到Github 对应的 part,并通过API获取到下载链接
	part := appMap[keyPair]
	if part != "" {
		downloadLink = githubAPI.SearchRelease(part)
	} else {
		err = fmt.Errorf("can not find the URL that supports your system/architecture %s", keyPair)
	}

	return downloadLink, err
}

// CheckToolbox 检查工具链是否完整,不完整会返回带有缺少的工具链的错误
func CheckToolbox(toolbox []string) error {
	var message string = ""

	// 检查工具链,如果有不存在的会写入message
	for _, tool := range toolbox {
		exist, _, _ := gek_exec.Exist(tool)
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
