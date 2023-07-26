package gExec

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

// Exist 检查指定程序是否存在于环境变量或者可执行文件的同一目录中
func Exist(program string) (path string, err error) {
    // 获取当前运行的可执行文件的目录
    exd, err := FindLocation()
    if err != nil {
        return "", err
    }
    // 当前运行的可执行文件的目录是否存在program
    exp := filepath.Join(exd, program)
    path, err = exec.LookPath(exp)
    if err != nil {
        // 当拼接路径不存在时,则从环境路径中查找program
        path, err = exec.LookPath(program)
        //当拼接路径及环境路径中均找不到program时返回false 空字符串 错误
        if err != nil {
            return "", err
        }
    }
    return path, nil
}

// FindLocation 返回当前运行的可执行文件的地址
// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
func FindLocation() (string, error) {
    // 当前可执行文件地址
    ex, err := os.Executable()
    if err != nil {
        return "", err
    }

    // 截取可执行文件地址的文件夹地址部分
    ex = filepath.Dir(ex)
    return ex, err
}

// StringToCmd 字符串按空格分词转换为 *exec.Cmd
func StringToCmd(cmdString string) (cmd *exec.Cmd) {
    cmdArgs := strings.Fields(cmdString)
    return exec.Command(cmdArgs[0], cmdArgs[1:]...)
}

// Run 执行命令并等待命令执行完成,命令可为 string类型或者 *exec.Cmd类型
func Run(command interface{}) (err error) {
    var cmd = &exec.Cmd{}

    // 同时处理输入的命令,string类型 或者 *exec.Cmd类型
    switch command.(type) {
    case string:
        cmd = StringToCmd(command.(string))
    case *exec.Cmd:
        cmd = command.(*exec.Cmd)
    default:
        return fmt.Errorf("the type of command %v is not supported", command)
    }

    // 实时输出结果
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout
    cmd.Stdin = os.Stdin

    // 运行命令并等待命令执行完成
    err = cmd.Run()
    return err
}

// CombinedOutput 执行命令并等待命令执行完成,返回运行后的输出和错误,命令可为 string类型或者 *exec.Cmd类型
func CombinedOutput(command interface{}) (output string, err error) {
    var cmd = &exec.Cmd{}

    // 同时处理输入的命令,string类型 或者 *exec.Cmd类型
    switch command.(type) {
    case string:
        cmd = StringToCmd(command.(string))
    case *exec.Cmd:
        cmd = command.(*exec.Cmd)
    default:
        return "", fmt.Errorf("the type of command %v is not supported", command)
    }

    // 运行命令,获取输出和错误
    outputByteSlice, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(outputByteSlice), err
}

// Output 执行命令并等待命令执行完成,返回运行后的输出,命令可为 string类型或者 *exec.Cmd类型
func Output(command interface{}) (output string, err error) {
    var cmd = &exec.Cmd{}

    // 同时处理输入的命令,string类型 或者 *exec.Cmd类型
    switch command.(type) {
    case string:
        cmd = StringToCmd(command.(string))
    case *exec.Cmd:
        cmd = command.(*exec.Cmd)
    default:
        return "", fmt.Errorf("the type of command %v is not supported", command)
    }

    // 运行命令,获取输出
    outputByteSlice, err := cmd.Output()
    if err != nil {
        return "", err
    }
    return string(outputByteSlice), err
}
