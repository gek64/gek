package unzip

import (
	"github.com/gek64/gek/gExec"
	"os"
	"os/exec"
)

// Decompress 解压文件到指定目录
func Decompress(zipFile string, location string, fileList ...string) (err error) {
	// 创建默认输出路径
	if location != "" {
		err = os.MkdirAll(location, 0755)
		if err != nil {
			return err
		}
	}

	// 不提供解压的部分文件时,解压所有文件
	if len(fileList) == 0 {
		return gExec.Run(exec.Command("unzip", "-d", location, "-o", zipFile))
	}

	// 提供解压的部分文件时,遍历解压指定文件
	for _, f := range fileList {
		err = gExec.Run(exec.Command("unzip", "-d", location, "-o", zipFile, f))
		if err != nil {
			return err
		}
	}
	return nil
}

// DecompressFileToByte 解压一个文件,并获取文件内容到比特切片
func DecompressFileToByte(zipFile string, fileInZip string) (fb []byte, err error) {
	return exec.Command("unzip", "-p", zipFile, fileInZip).Output()
}

// DecompressFileToFile 解压一个文件,并获取文件内容存储到另一个文件
func DecompressFileToFile(zipFile string, fileInZip string, newFile string) (err error) {
	bytes, err := DecompressFileToByte(zipFile, fileInZip)
	if err != nil {
		return err
	}
	return os.WriteFile(newFile, bytes, 0644)
}
