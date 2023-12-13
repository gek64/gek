package gExec

import (
	"os"
	"os/exec"
)

// Run 执行命令, 等待命令执行完成
func Run(cmd *exec.Cmd) (err error) {
	// 实时输出结果
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// Start 执行命令, 不等待命令执行完成
func Start(cmd *exec.Cmd) (err error) {
	// 实时输出结果
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Start()
}

// Output 执行命令, 等待命令执行完成, 返回运行后的输出
func Output(cmd *exec.Cmd) (output []byte, err error) {
	return cmd.Output()
}

// CombinedOutput 执行命令, 等待命令执行完成, 返回运行后的输出和错误
func CombinedOutput(cmd *exec.Cmd) (output []byte, err error) {
	return cmd.CombinedOutput()
}
