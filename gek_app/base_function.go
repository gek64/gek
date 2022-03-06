package gek_app

import (
	"gek_downloader"
	"gek_exec"
	"os"
	"os/exec"
)

var (
	// SupportedOS 支持的系统
	SupportedOS = []string{"linux", "freebsd", "windows", "darwin", "android", "ios", "openbsd", "netbsd", "dragonfly", "solaris", "aix", "illumos", "js", "plan9"}
	// SupportedArch 支持的架构
	SupportedArch = []string{"amd64", "386", "arm64", "arm", "riscv64", "mips64", "mips64le", "mips", "mipsle", "ppc64", "ppc64le", "s390x", "wasm"}
)

// Download 下载应用到临时目录
func Download(downloadLink, location string) (err error) {
	// 建立下载文件夹
	// 已经存在就删除重建
	_, err = os.Stat(location)
	if os.IsExist(err) {
		err = os.RemoveAll(location)
		if err != nil {
			return err
		}
	}
	err = os.MkdirAll(location, 0755)
	if err != nil {
		return err
	}

	// 下载文件到文件夹
	err = gek_downloader.Downloader(downloadLink, location, "")
	if err != nil {
		return err
	}
	return nil
}

// Extract 从压缩文件中按照给定的的文件列表解压需要的文件到输出路径
func Extract(archiveFile string, fileList []string, location string) (err error) {
	// 如果输出路径不存在则创建
	_, err = os.Stat(location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(location, 0755)
		if err != nil {
			return err
		}
	}
	// 循环解压
	for _, file := range fileList {
		err = gek_exec.Run(exec.Command("unzip", "-o", archiveFile, file, "-d", location))
		if err != nil {
			return err
		}
	}
	return nil
}

// ChmodList 给列表中的路径赋权
func ChmodList(pathList []string, mode os.FileMode) (err error) {
	for _, p := range pathList {
		err = os.Chmod(p, mode)
		if err != nil {
			return err
		}
	}
	return nil
}
