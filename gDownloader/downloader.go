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

	// 处理输出文件夹及文件
	if filename == "" {
		filename = path.Base(resp.Request.URL.Path)
	}
	if outputFolder != "" {
		err = os.MkdirAll(outputFolder, 0755)
		if err != nil {
			return err
		}
	}
	output, err := os.Create(filepath.Join(outputFolder, filename))
	if err != nil {
		return err
	}

	// 将数据写入到文件中
	_, err = io.Copy(output, resp.Body)
	return err
}

// DownloadWithCurl 使用 curl 下载
func DownloadWithCurl(fileUrl string, filename string, outputFolder string) (err error) {
	// 不指定文件名
	if filename == "" {
		return gExec.Run(exec.Command("curl", "--create-dirs", "--output-dir", outputFolder, "-LOJ", fileUrl))
	}

	// 指定文件名
	output, err := os.Create(filepath.Join(outputFolder, filename))
	if err != nil {
		return err
	}
	return gExec.Run(exec.Command("curl", "--create-dirs", "--output-dir", outputFolder, "-Lo", output.Name(), fileUrl))
}
