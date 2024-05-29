package gDownloader

import (
	"fmt"
	"github.com/gek64/gek/gExec"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

// Download 下载
func Download(fileUrl string, outputFile string, outputFolder string) (err error) {
	// 使用get方法连接url
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}

	// 输出文件
	if outputFile == "" {
		outputFile = path.Base(fileUrl)
	}

	// 输出文件夹
	if outputFolder != "" {
		outputFile = filepath.Join(outputFolder, outputFile)
	}

	// 创建最终输出文件夹
	err = os.MkdirAll(filepath.Dir(outputFile), 0755)
	if err != nil {
		return err
	}

	o, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	// 函数结束时关闭文件,不然会造成其他对此文件的操作失败
	defer o.Close()

	// 将数据写入到输出文件中
	_, err = io.Copy(o, resp.Body)
	return err
}

// DownloadWithCurl 使用 curl 下载
func DownloadWithCurl(fileUrl string, outputFile string, outputFolder string) (err error) {
	var cmd *exec.Cmd

	if outputFolder == "" {
		if outputFile == "" {
			cmd = exec.Command("curl", "-LOJ", fileUrl)
		} else {
			err = os.MkdirAll(filepath.Dir(outputFile), 0755)
			if err != nil {
				return err
			}

			cmd = exec.Command("curl", "-Lo", outputFile, fileUrl)
		}
	} else {
		if outputFile == "" {
			err = os.MkdirAll(outputFolder, 0755)
			if err != nil {
				return err
			}

			cmd = exec.Command("curl", "-LOJ", fileUrl)
			cmd.Dir = outputFolder
		} else {
			outputFile = filepath.Join(outputFolder, outputFile)
			err = os.MkdirAll(filepath.Dir(outputFile), 0755)
			if err != nil {
				return err
			}

			cmd = exec.Command("curl", "-Lo", outputFile, fileUrl)
			fmt.Println(cmd.String())
		}
	}

	return gExec.Run(cmd)
}
