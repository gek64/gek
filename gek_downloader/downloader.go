package gek_downloader

import (
	"fmt"
	"gek_exec"
	"gek_file"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// GetFileName 从url中获取文件名
func GetFileName(url string) (fileName string, err error) {
	fileName = ""

	// 使用get方法连接url
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
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

func Downloader(url string, outputFile ...interface{}) error {
	var fileName = ""

	// 使用get方法连接url
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(response.Body)

	// outputFile 有则从中提取文件路径
	if len(outputFile) != 0 && outputFile[0] != "" {
		fileName = outputFile[0].(string)
	}
	// outputFile为空则从header中提取文件路径
	if fileName == "" {
		// 从header中提取文件路径
		contentDisposition := response.Header.Get("Content-Disposition")
		if contentDisposition != "" {
			fileName = strings.Split(contentDisposition, "filename=")[1]
		}
	}
	// header中无文件路径则使用默认路径
	if fileName == "" {
		fileName = "default_file"
	}

	// 新建输出文件
	output, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(output)

	// 将数据写入到文件中
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func ExternalDownloader(url string, outputFile ...interface{}) error {
	var downloader string = ""
	// 循环找到是否存在外部下载器
	externalDownloaders := []string{"aria2c", "wget", "curl"}
	for _, d := range externalDownloaders {
		exist, _, _ := gek_exec.Exist(d)
		if exist {
			downloader = "aria2c"
			break
		}
	}
	// 外部下载器下载
	switch downloader {
	case "aria2c":
		err := gek_exec.Run(exec.Command("aria2c", "-c", "-s", "16", "-x", "16", url))
		if err != nil {
			return err
		}
	case "wget":
		err := gek_exec.Run(exec.Command("wget", url))
		if err != nil {
			return err
		}
	case "curl":
		err := gek_exec.Run(exec.Command("curl", "-LOJ", url))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("can not find aria2c, wget and curl")
	}

	// outputFile存在则从url中提取文件名,并改名已下载文件为outputFile提供的文件名
	if len(outputFile) != 0 && outputFile[0] != "" {
		fileName, err := GetFileName(url)
		if err != nil {
			log.Println(err)
		}

		exist, _, _ := gek_file.Exist(fileName)

		if exist {
			err := os.Rename(fileName, outputFile[0].(string))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
