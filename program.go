package gopkg

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ProgramExist Check if the program exists in environment variable or input path
func ProgramExist(programName string, programPath ...interface{}) (bool, string, error) {
	var found bool
	var path string

	// cd to exec file path
	execFilePath, err := PathExecFile()
	log.Print(err)
	os.Chdir(execFilePath)

	// check input path
	if len(programPath) != 0 {
		switch programPath[0].(type) {
		case string:
			path, err = exec.LookPath(programPath[0].(string))
			if err != nil {
				found = false
			} else {
				//When path is found, get its absolute path and convert all "\"to "/"
				path, err = filepath.Abs(path)
				path = strings.Replace(path, "\\", "/", -1)
				found = true
			}
		default:
			log.Fatal("Error the parameters:", programPath[0], "is not a string")
		}
	}

	// check environment variable
	if found == false {
		path, err = exec.LookPath(programName)
		if err == nil {
			found = true
		}
	}
	os.Chdir(OriginalWorkingPath)
	return found, path, err
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
