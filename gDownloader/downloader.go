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
func Download(fileUrl string, filename string, outputFolder string) (err error) {
	// 使用get方法连接url
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}

	// 文件名
	if filename == "" {
		filename = path.Base(fileUrl)
	}
	// 指定的输出文件夹
	outputFolder, _ = filepath.Abs(outputFolder)
	err = os.MkdirAll(outputFolder, 0755)
	if err != nil {
		return err
	}

	// 文件名中的输出文件夹
	err = os.MkdirAll(filepath.Dir(filepath.Join(outputFolder, filename)), 0755)
	if err != nil {
		return err
	}

	// 创建输出文件
	o, err := os.Create(filepath.Join(outputFolder, filename))
	if err != nil {
		return err
	}

	// 将数据写入到输出文件中
	_, err = io.Copy(o, resp.Body)
	return err
}

// DownloadWithCurl 使用 curl 下载
func DownloadWithCurl(fileUrl string, filename string, outputFolder string) (err error) {
	// 处理输出文件夹
	outputFolder, _ = filepath.Abs(outputFolder)
	err = os.MkdirAll(outputFolder, 0755)
	if err != nil {
		return err
	}

	// 不指定文件名
	if filename == "" {
		cmd := exec.Command("curl", "--create-dirs", "-LOJ", fileUrl)
		cmd.Dir = outputFolder
		return gExec.Run(cmd)
	}

	// 指定文件名
	cmd := exec.Command("curl", "--create-dirs", "-Lo", filepath.Join(outputFolder, filename), fileUrl)
	cmd.Dir = outputFolder
	return gExec.Run(cmd)
}
