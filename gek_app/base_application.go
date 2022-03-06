package gek_app

import (
	"fmt"
	"gek_downloader"
	"gek_github"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	// TEMP 应用下载临时目录
	TEMP = "/tmp/gek_app_installer/"
)

type Application struct {
	// app文件
	AppFiles []string
	// 应用URL
	Url string
	// 是否需要解压
	NeedExtract bool
	// 应用安装文件夹
	Location string
}

// NewApplication 新建应用
func NewApplication(appFiles []string, url string, needExtract bool, location string) (a Application) {
	return Application{AppFiles: appFiles, Url: url, NeedExtract: needExtract, Location: location}
}

// NewApplicationFromGithub 新建应用(从Github)
func NewApplicationFromGithub(appFiles []string, repo string, appMap map[string]string, needExtract bool, location string) (a Application, err error) {
	// 获取应用链接
	downloadLink, err := gek_github.GetDownloadLink(repo, appMap)
	if err != nil {
		return Application{}, err
	}
	return NewApplication(appFiles, downloadLink, needExtract, location), nil
}

// Install 安装应用
func (a Application) Install() (err error) {
	// 检查安装文件夹情况
	// 不存在则新建
	_, err = os.Stat(a.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(a.Location, 0755)
		if err != nil {
			return err
		}
	}

	// 应用下载+解压安装
	if a.NeedExtract {
		// 检查临时文件夹情况
		_, err = os.Stat(TEMP)
		if os.IsNotExist(err) {
			err = os.MkdirAll(TEMP, 0755)
			if err != nil {
				return err
			}
		}
		// 结束后删除临时文件夹
		defer func(path string) {
			err = os.RemoveAll(path)
			if err != nil {
				log.Panicln(err)
			}
		}(TEMP)

		// 下载压缩文件到临时文件夹
		err = gek_downloader.Downloader(a.Url, TEMP, "")
		if err != nil {
			return err
		}

		// 读取下载的压缩文件名
		var zipFile string
		fileInfos, err := ioutil.ReadDir(TEMP)
		if err != nil {
			fmt.Println(err)
		}
		for _, f := range fileInfos {
			if strings.Contains(f.Name(), ".zip") {
				zipFile = f.Name()
				break
			}
		}
		if zipFile == "" {
			return fmt.Errorf("can't find the download application archive file")
		}
		// 解压文件到安装文件夹
		err = Extract(filepath.Join(TEMP, zipFile), a.AppFiles, a.Location)
		if err != nil {
			return err
		}
	} else {
		err = gek_downloader.Downloader(a.Url, "", filepath.Join(a.Location, a.AppFiles[0]))
		if err != nil {
			return err
		}
	}

	// 可执行文件赋权0755
	for _, appFile := range a.AppFiles {
		err = os.Chmod(filepath.Join(a.Location, appFile), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// InstallFromLocal 安装应用(从本地文件,无需网络下载)
func (a Application) InstallFromLocal(localFile string) (err error) {
	// 检查本地文件是否存在
	_, err = os.Stat(localFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localFile)
	}

	// 检查安装文件夹情况
	// 不存在则新建
	_, err = os.Stat(a.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(a.Location, 0755)
		if err != nil {
			return err
		}
	}

	// 应用解压安装
	if a.NeedExtract {
		// 解压文件到安装文件夹
		err = Extract(localFile, a.AppFiles, a.Location)
		if err != nil {
			return err
		}
	} else {
		// 打开本地文件
		fs, err := os.OpenFile(localFile, os.O_RDWR, 0755)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Panicln(err)
			}
		}(fs)
		// 创建复制目标文件
		fd, err := os.OpenFile(filepath.Join(a.Location, a.AppFiles[0]), os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Panicln(err)
			}
		}(fd)
		// 复制文件
		_, err = io.Copy(fd, fs)
		if err != nil {
			return err
		}
	}

	// 可执行文件赋权0755
	for _, appFile := range a.AppFiles {
		err = os.Chmod(filepath.Join(a.Location, appFile), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// Uninstall 卸载应用
func (a Application) Uninstall() (err error) {
	// 检测应用安装情况
	_, err = os.Stat(filepath.Join(a.Location, a.AppFiles[0]))
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find app location %s", filepath.Join(a.Location, a.AppFiles[0]))
	}

	// 删除应用文件
	for _, app := range a.AppFiles {
		err = os.RemoveAll(filepath.Join(a.Location, app))
		if err != nil {
			return err
		}
	}

	return nil
}
