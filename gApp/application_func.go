package gApp

import (
	"fmt"
	"github.com/gek64/gek/gDownloader"
	"log"
	"os"
	"path/filepath"
)

// Install 安装应用
func (a Application) Install(tempLocation string, needDownload bool, needExtract bool) (err error) {
	// 检查安装文件夹情况
	// 不存在则新建
	_, err = os.Stat(a.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(a.Location, 0755)
		if err != nil {
			return err
		}
	}

	// 建立临时文件夹
	appTemp := NewTemp(tempLocation)
	err = appTemp.Create()
	if err != nil {
		return err
	}
	defer func(appTemp Temp) {
		err := appTemp.Delete()
		if err != nil {
			log.Panicln(err)
		}
	}(appTemp)

	// 下载文件
	if needDownload {
		err = gDownloader.Downloader(a.Url, appTemp.Location, "")
		if err != nil {
			return err
		}
	}
	// 解压文件
	if needExtract {
		err = a.extract(appTemp.Location)
		if err != nil {
			return err
		}
	}
	// 复制文件到安装文件夹
	err = a.copy(appTemp.Location)
	if err != nil {
		return err
	}
	// 赋权0755
	err = a.chmod(0755)

	return err
}

// InstallFromLocal 安装应用(从本地文件,无需网络下载)
func (a Application) InstallFromLocal(tempLocation string, localFile string, needExtract bool) (err error) {
	// 检查本地文件是否存在
	_, err = os.Stat(localFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localFile)
	}

	// 建立临时文件夹
	appTemp := NewTemp(tempLocation)
	err = appTemp.Create()

	// 复制本地文件到临时文件夹
	err = CopyFiles(localFile, appTemp.Location)
	if err != nil {
		return err
	}

	err = a.Install(appTemp.Location, false, needExtract)

	return err
}

// Uninstall 卸载应用
func (a Application) Uninstall() (err error) {
	// 检测应用安装情况
	_, err = os.Stat(filepath.Join(a.Location, a.AppFiles[0]))
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find app location %s", filepath.Join(a.Location, a.AppFiles[0]))
	}
	// 删除应用文件或者应用文件夹
	if a.UninstallDeleteLocationFolder {
		err = os.RemoveAll(a.Location)
		if err != nil {
			return err
		}
	} else {
		for _, app := range a.AppFiles {
			err = os.RemoveAll(filepath.Join(a.Location, app))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
