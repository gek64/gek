package gek_app

import (
	"gek_github"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Application struct {
	// app文件
	AppFiles []string
	// 应用URL
	Url string
	// 应用安装文件夹
	Location string
	// 删除时是否删除安装路径
	UninstallDeleteLocationFolder bool
}

// NewApplication 新建应用
func NewApplication(appFiles []string, url string, location string, uninstallDeleteLocationFolder bool) (a Application) {
	return Application{AppFiles: appFiles, Url: url, Location: location, UninstallDeleteLocationFolder: uninstallDeleteLocationFolder}
}

// NewApplicationFromGithub 新建应用(从Github)
func NewApplicationFromGithub(appFiles []string, repo string, appMap map[string]string, location string, uninstallDeleteLocationFolder bool, tempLocation string) (a Application, err error) {
	// 获取应用链接
	downloadLink, err := gek_github.GetDownloadLink(repo, appMap)
	if err != nil {
		return Application{}, err
	}
	return NewApplication(appFiles, downloadLink, location, uninstallDeleteLocationFolder), nil
}

// 应用解压
func (a Application) extract(tempLocation string) (err error) {
	fileInfo, err := ioutil.ReadDir(tempLocation)
	if err != nil {
		return err
	}
	for _, f := range fileInfo {
		if strings.Contains(f.Name(), ".zip") {
			err = ExtractZip(filepath.Join(tempLocation, f.Name()), tempLocation)
			if err != nil {
				return err
			}
			break
		}
		if strings.Contains(f.Name(), ".tar") {
			err = ExtractTar(filepath.Join(tempLocation, f.Name()), tempLocation)
			if err != nil {
				return err
			}
			break
		}
	}
	return nil
}

// 复制文件
func (a Application) copy(tempLocation string) (err error) {
	for _, appFile := range a.AppFiles {
		err = CopyFiles(filepath.Join(tempLocation, appFile), filepath.Join(a.Location, appFile))
		if err != nil {
			return err
		}
	}
	return nil
}

// 应用文件赋权
func (a Application) chmod(mode int) (err error) {
	for _, appFile := range a.AppFiles {
		err = os.Chmod(filepath.Join(a.Location, appFile), os.FileMode(mode))
		if err != nil {
			return err
		}
	}
	return nil
}
