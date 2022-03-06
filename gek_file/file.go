package gek_file

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filePath string, content string) (name string, err error) {
	// Create temp f
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	// Close the f
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(f)

	// write to the f
	_, err = f.WriteString(content)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func CreateDir(dirPath string) (err error) {
	// Create temp dir
	err = os.MkdirAll(dirPath, 0755)
	return err
}

func CreateRandomFile(dir string, pattern string, content string) (name string, err error) {
	// Create temp file
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", err
	}

	defer func(tmpFile *os.File) {
		err = tmpFile.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(f)

	// write to the file
	_, err = f.WriteString(content)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func CreateRandomDir(dir string, pattern string) (name string, err error) {
	// Create temp dir
	name, err = ioutil.TempDir(dir, pattern)
	if err != nil {
		return "", err
	}

	return name, nil
}

func Exist(filePath string) (exist bool, isDir bool) {
	fileInfo, err := os.Stat(filePath)
	return os.IsExist(err), fileInfo.IsDir()
}
