package gek

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

//// ProgramExist Check if the program exists in environment variable or input path
//func ProgramExist(programName string, programPath ...interface{}) (bool, string, error) {
//	var found bool
//	var path string
//
//	// cd to exec file path
//	execFilePath, err := FindExecutableDir()
//	log.Print(err)
//	os.Chdir(execFilePath)
//
//	// check input path
//	if len(programPath) != 0 {
//		switch programPath[0].(type) {
//		case string:
//			path, err = exec.LookPath(programPath[0].(string))
//			if err != nil {
//				found = false
//			} else {
//				//When path is found, get its absolute path and convert all "\"to "/"
//				path, err = filepath.Abs(path)
//				path = strings.Replace(path, "\\", "/", -1)
//				found = true
//			}
//		default:
//			log.Fatal("Error the parameters:", programPath[0], "is not a string")
//		}
//	}
//
//	// check environment variable
//	if found == false {
//		path, err = exec.LookPath(programName)
//		if err == nil {
//			found = true
//		}
//	}
//	os.Chdir(OriginalWorkingPath)
//	return found, path, err
//}

// ProgramExist Check if the program exists in environment variable or input path
func ProgramExist(program string) (string, error) {
	var path string
	// 获取当前运行的可执行文件的目录
	exd, err := FindExecutableDir()
	if err != nil {
		return "", err
	}
	// 当前运行的可执行文件的目录与program拼接
	exp := filepath.Join(exd, program)
	// 首先查看拼接后的program是否存在
	path, err = exec.LookPath(exp)
	if err != nil {
		// 当拼接路径不存在时,则从环境路径中查找program
		path, err = exec.LookPath(program)
		//当拼接路径及环境路径中均找不到program时返回空字符串和错误
		if err != nil {
			return "", err
		}
		return path, nil
	}
	return path, nil

}

// FindExecutableDir 返回当前运行的可执行文件的目录
// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
func FindExecutableDir() (string, error) {
	// 当前可执行文件地址
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	// 截取可执行文件地址的文件夹地址部分
	ex = filepath.Dir(ex)
	return ex, err
}

// ProgramRealtimeOutput Execute commands and get real-time output
func ProgramRealtimeOutput(cmd *exec.Cmd) error {

	// Only works if the full file path has no spaces
	// cmdArgs := strings.Fields(cmdString)
	// cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)

	writer := io.Writer(os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return err
}
