package gek_file

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filePath string, content string) (file *os.File, err error) {
	// Create temp file
	file, err = os.Create(filePath)
	if err != nil {
		return nil, err
	}

	// Close the file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(file)

	// write to the file
	_, err = file.WriteString(content)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func CreateDir(dirPath string) (err error) {
	// Create temp dir
	err = os.MkdirAll(dirPath, 755)
	return err
}

func CreateRandomFile(dir string, pattern string, content string) (f *os.File, err error) {
	// Create temp file
	f, err = ioutil.TempFile(dir, pattern)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return f, nil
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
