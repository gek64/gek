package tar

import (
	"github.com/gek64/gek/gExec"
	"os"
	"os/exec"
)

// Decompress 解压文件到指定目录
func Decompress(tarFile string, location string, fileList ...string) (err error) {
	// 默认创建输出路径
	err = os.MkdirAll(location, 0755)
	if err != nil {
		return err
	}

	// 不提供解压的部分文件时,解压所有文件
	if len(fileList) == 0 {
		return gExec.Run(exec.Command("tar", "-xvf", tarFile, "-C", location))
	}

	// 提供解压的部分文件时,遍历解压指定文件
	for _, f := range fileList {
		err = gExec.Run(exec.Command("tar", "-xvf", tarFile, "-C", location, f))
		if err != nil {
			return err
		}
	}
	return nil
}

// DecompressFileToByte 解压一个文件,并获取文件内容到比特切片
func DecompressFileToByte(tarFile string, fileInTar string) (fb []byte, err error) {
	return gExec.Output(exec.Command("tar", "-xvOf", tarFile, fileInTar))
}

// DecompressFileToFile 解压一个文件,并获取文件内容存储到另一个文件
func DecompressFileToFile(tarFile string, fileInTar string, newFile string) (err error) {
	bytes, err := DecompressFileToByte(tarFile, fileInTar)
	if err != nil {
		return err
	}
	return os.WriteFile(newFile, bytes, 0755)
}
