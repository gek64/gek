package gek_github

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
)

var (
	// 系统架构来源 https://golang.org/doc/install/source#environment
	osList   = []string{"aix", "android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"}
	archList = []string{"amd64", "386", "arm64", "arm", "mips64", "mips64le", "mips", "mipsle", "ppc64", "ppc64le", "riscv64", "s390x", "wasm"}
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
		downloadLink, err = githubAPI.SearchRelease(part)
		if err != nil {
			return "", err
		}
	} else {
		err = fmt.Errorf("can not find the URL that supports your system/architecture %s", keyPair)
	}

	return downloadLink, err
}

// GetFileName 从Github应用对应系统/架构对的下载链接中获取文件名
func GetFileName(repo string, appMap map[string]string) (fileName string, err error) {
	// 获取url
	url, err := GetDownloadLink(repo, appMap)
	if err != nil {
		return "", err
	}

	// 使用get方法连接url
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(response.Body)

	// 从header中提取文件路径
	contentDisposition := response.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		fileName = strings.Split(contentDisposition, "filename=")[1]
	}

	if fileName == "" {
		return "", fmt.Errorf("can not get file name from %s", url)
	}

	return fileName, nil
}
