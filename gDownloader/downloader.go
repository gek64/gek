package gDownloader

import (
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
		outputFile = path.Base(resp.Request.URL.Path)
	}

	// 仅有这一种情况需要进行文件与输出路径的拼接
	if filepath.Dir(outputFile) == "." && outputFolder != "" {
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
	// 参数中未指定输出文件名时,使用参数中下载url的最后路径作为文件名
	if outputFile == "" {
		resp, err := http.Get(fileUrl)
		if err != nil {
			return err
		}
		outputFile = path.Base(resp.Request.URL.Path)
	}

	// 仅有这一种情况需要进行文件与输出路径的拼接
	if filepath.Dir(outputFile) == "." && outputFolder != "" {
		outputFile = filepath.Join(outputFolder, outputFile)
	}

	// 创建最终输出文件夹
	err = os.MkdirAll(filepath.Dir(outputFile), 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("curl", "-Lo", outputFile, fileUrl)
	return gExec.Run(cmd)
}
