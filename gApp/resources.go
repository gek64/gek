package gApp

import (
	"fmt"
	"github.com/gek64/gek/gDownloader"
	"os"
	"path/filepath"
	"strings"
)

type Resources struct {
	// 资源文件
	Files []string
	// 资源链接
	Urls []string
	// 资源安装文件夹
	Location string
	// 删除时是否删除安装路径
	UninstallDeleteLocationFolder bool
}

// NewResources 新建资源
func NewResources(files []string, urls []string, location string, uninstallDeleteLocationFolder bool) (r Resources) {
	return Resources{Files: files, Urls: urls, Location: location, UninstallDeleteLocationFolder: uninstallDeleteLocationFolder}
}

// InstallFromInternet 从网络安装资源
func (r Resources) InstallFromInternet() (err error) {
	// 检测资源安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(r.Location, 0755)
		if err != nil {
			return err
		}
	}
	// 下载资源文件到资源安装路径
	for _, url := range r.Urls {
		err = gDownloader.Downloader(url, r.Location, "")
		if err != nil {
			return err
		}
	}

	// 赋权644
	for _, file := range r.Files {
		err = os.Chmod(filepath.Join(r.Location, file), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// InstallFromLocation 从本地目录安装资源
func (r Resources) InstallFromLocation(localLocation string) (err error) {
	// 检查本地路径是否存在
	_, err = os.Stat(localLocation)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localLocation)
	}

	// 检测资源安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(r.Location, 0755)
		if err != nil {
			return err
		}
	}

	// 拷贝文件夹中资源文件到安装文件夹
	fileInfos, err := os.ReadDir(localLocation)
	if err != nil {
		return err
	}

	for _, rFile := range r.Files {
		for _, file := range fileInfos {
			if strings.Contains(file.Name(), rFile) {
				err = CopyFiles(filepath.Join(localLocation, file.Name()), filepath.Join(r.Location, rFile))
				if err != nil {
					return err
				}
				break
			}
		}
	}

	// 赋权0644
	for _, rFile := range r.Files {
		err = os.Chmod(filepath.Join(r.Location, rFile), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Uninstall 卸载资源,并删除资源安装文件夹
func (r Resources) Uninstall() (err error) {
	// 检测资源安装路径是否存在
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find resources location %s", r.Location)
	}

	if r.UninstallDeleteLocationFolder {
		err = os.RemoveAll(r.Location)
	} else {
		for _, file := range r.Files {
			err = os.RemoveAll(filepath.Join(r.Location, file))
		}
	}
	return err
}
