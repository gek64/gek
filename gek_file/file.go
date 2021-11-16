package gek_file

import (
	"io/ioutil"
	"os"
)

func CreateFile(filePath string, content string) (file *os.File, err error) {
	// Create temp file
	file, err = os.Create(filePath)
	if err != nil {
		return nil, err
	}

	// write to the file
	text := []byte(content)
	_, err = file.Write(text)
	if err != nil {
		return nil, err
	}

	// Close the file
	err = file.Close()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func CreateDir(dirPath string) (err error) {
	// Create temp dir
	err = os.MkdirAll(dirPath, 751)
	if err != nil {
		return err
	}
	return nil
}

func CreateRandomFile(dir string, pattern string, content string) (tmpFile *os.File, err error) {
	// Create temp file
	tmpFile, err = ioutil.TempFile(dir, pattern)
	if err != nil {
		return nil, err
	}

	// Remember to clean up the file afterwards
	// defer os.Remove(tmpFile.Name())

	// write to the file
	text := []byte(content)
	_, err = tmpFile.Write(text)
	if err != nil {
		return nil, err
	}

	// Close the file
	err = tmpFile.Close()
	if err != nil {
		return nil, err
	}

	return tmpFile, nil
}

func CreateRandomDir(dir string, pattern string) (tmpDir string, err error) {
	// Create temp dir
	tmpDir, err = ioutil.TempDir(dir, pattern)
	if err != nil {
		return "", err
	}

	// Remember to clean up the file afterwards
	// defer os.RemoveAll(tmpDir)

	return tmpDir, nil
}

func Exist(filePath string) (exist bool, isDir bool, err error) {
	exist = false
	isDir = false

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, false, err
	}

	return true, fileInfo.IsDir(), nil
}
