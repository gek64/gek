package gek_github

import (
	"fmt"
	"runtime"
)

var (
	// 系统架构来源 https://golang.org/doc/install/source#environment
	osList   = []string{"android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"}
	archList = []string{"amd64", "386", "arm", "arm64", "ppc64le", "mips64le", "mips64", "mipsle", "mips", "s390x", "wasm"}
)

// GetDownloadLink 获取Github应用对应系统/架构对的下载链接
func GetDownloadLink(repo string, appMap map[string]string) (downloadLink string, err error) {
	downloadLink = ""
	keyPair := ""

	// 获取Github API
	githubAPI, err := NewGithubAPI(repo)
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
