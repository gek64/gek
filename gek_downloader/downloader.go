package gek_downloader

import (
	"fmt"
	"gek_exec"
	"gek_math"
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
			log.Panicln(err)
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

// InternalDownloaderWithFolder 内部下载器,指定输出文件夹
func InternalDownloaderWithFolder(url string, outputFolder string) (err error) {
	// 使用get方法连接url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(response.Body)

	var fileName string = ""
	// 从header中提取文件路径
	contentDisposition := response.Header.Get("Content-Disposition")
	if contentDisposition != "" {
		fileName = strings.Split(contentDisposition, "filename=")[1]
	}
	// header中无文件路径则使用默认路径
	if fileName == "" {
		fileName = gek_math.RandStringRunes(5)
	}

	// 记录当前工作路径
	originDir, err := os.Getwd()
	if err != nil {
		return err
	}
	// 跳转到目标下载文件夹
	err = os.Chdir(outputFolder)
	if err != nil {
		return err
	}

	// 新建输出文件
	output, err := os.Create(fileName)
	if err != nil {
		return err
	}

	// 销毁时关闭文件
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(output)

	// 将数据写入到文件中
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	// 跳转回到原始工作路径
	err = os.Chdir(originDir)
	if err != nil {
		return err
	}

	return nil
}

// InternalDownloaderFilePath 内部下载器,指定输出文件路径
func InternalDownloaderFilePath(url string, outputFile string) (err error) {
	// 使用get方法连接url
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(response.Body)

	// 如果默认文件已存在,则删除已存在文件
	_, err = os.Stat(outputFile)
	if os.IsExist(err) {
		err = os.RemoveAll(outputFile)
		if err != nil {
			return err
		}
	}

	// 新建输出文件
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	// 销毁时关闭文件
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(output)

	// 将数据写入到文件中
	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}

// ExternalDownloaderWithFolder 外部下载器,指定输出文件夹
func ExternalDownloaderWithFolder(url string, outputFolder string) (err error) {
	var downloader string = ""
	// 循环找到是否存在外部下载器
	externalDownloaders := []string{"aria2c", "wget", "curl", "fetch"}
	for _, d := range externalDownloaders {
		exist, _, _ := gek_exec.Exist(d)
		if exist {
			downloader = d
			break
		}
	}

	switch downloader {
	case "aria2c":
		err = gek_exec.Run(exec.Command("aria2c", "-s", "16", "-x", "16", url, "--allow-overwrite=true", "-d", outputFolder))
		if err != nil {
			return err
		}
	case "wget":
		err = gek_exec.Run(exec.Command("wget", url, "-P", outputFolder, "-N"))
		if err != nil {
			return err
		}
	case "curl":
		err = gek_exec.Run(exec.Command("curl", "--create-dirs", "--output-dir", outputFolder, "-LOJ", url))
		if err != nil {
			return err
		}
	case "fetch":
		// 获取原始工作路径,退出时恢复到到原始工作路径
		origin, err := os.Getwd()
		if err != nil {
			return err
		}
		defer func(dir string) {
			err := os.Chdir(dir)
			if err != nil {
				log.Panicln(err)
			}
		}(origin)
		// 输出文件夹不存在则创建
		_, err = os.Stat(outputFolder)
		if os.IsNotExist(err) {
			err = os.MkdirAll(outputFolder, 755)
			if err != nil {
				return err
			}
		}
		// 跳转到输出文件夹
		err = os.Chdir(outputFolder)
		if err != nil {
			return err
		}
		// 下载文件
		err = gek_exec.Run(exec.Command("fetch", url))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("can not find aria2c, wget, curl and fetch")
	}

	return nil
}

// ExternalDownloaderWithFilePath 外部下载器,指定输出文件路径
func ExternalDownloaderWithFilePath(url string, outputFile string) (err error) {
	var downloader string = ""
	// 循环找到是否存在外部下载器
	externalDownloaders := []string{"aria2c", "wget", "curl", "fetch"}
	for _, d := range externalDownloaders {
		exist, _, _ := gek_exec.Exist(d)
		if exist {
			downloader = d
			break
		}
	}

	// 按指定文件名下载
	switch downloader {
	case "aria2c":
		err = gek_exec.Run(exec.Command("aria2c", "-s", "16", "-x", "16", url, "--allow-overwrite=true", "-o", outputFile))
		if err != nil {
			return err
		}
	case "wget":
		err = gek_exec.Run(exec.Command("wget", url, "-O", outputFile))
		if err != nil {
			return err
		}
	case "curl":
		err = gek_exec.Run(exec.Command("curl", "--create-dirs", "-Lo", outputFile, url))
		if err != nil {
			return err
		}
	case "fetch":
		err = gek_exec.Run(exec.Command("fetch", url, "-o", outputFile))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("can not find aria2c, wget, curl and fetch")
	}
	return nil
}

// Downloader 综合下载器,结合外部下载器与内部下载器同时使用
// outputFile,outputFolder可以均为空,或一个为空一个不为空,均不为空的时候按照优先级进行工作
// 优先度 outputFile > outputFolder > 两者均不设置(默认outputFolder="."作为默认设置)
func Downloader(url string, outputFolder string, outputFile string) (err error) {
	// outputFile 优先级第一,如果设置则采用为输出文件名称
	// outputFolder 优先级第二,如果未设置outputFile,则采用为文件输出目录
	// outputFile outputFolder均未设置,则默认采用outputFolder="."作为默认设置
	if outputFile != "" {
		err = ExternalDownloaderWithFilePath(url, outputFile)
		if err != nil {
			err = InternalDownloaderFilePath(url, outputFile)
			if err != nil {
				return err
			}
		}
	} else if outputFolder != "" {
		err = ExternalDownloaderWithFolder(url, outputFolder)
		if err != nil {
			err = InternalDownloaderWithFolder(url, outputFolder)
			if err != nil {
				return err
			}
		}
	} else {
		err = ExternalDownloaderWithFolder(url, ".")
		if err != nil {
			err = InternalDownloaderWithFolder(url, ".")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
