package gFile

import (
	"log"
	"os"
)

func CreateFile(filePath string, content string) (name string, err error) {
	return CreateRawFile(filePath, []byte(content))
}

func CreateRawFile(filePath string, content []byte) (name string, err error) {
	// Create file
	f, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	// Close the file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	// write to the file
	_, err = f.Write(content)
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
	return CreateRandomRawFile(dir, pattern, []byte(content))
}

func CreateRandomRawFile(dir string, pattern string, content []byte) (name string, err error) {
	// Create temp file
	f, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return "", err
	}

	// Close the file
	defer func(tmpFile *os.File) {
		err = tmpFile.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)

	// write to the file
	_, err = f.Write(content)
	if err != nil {
		return "", err
	}

	return f.Name(), nil
}

func CreateRandomDir(dir string, pattern string) (name string, err error) {
	// Create temp dir
	name, err = os.MkdirTemp(dir, pattern)
	if err != nil {
		return "", err
	}

	return name, nil
}

func Exist(filePath string) (exist bool, isDir bool) {
	fileInfo, err := os.Stat(filePath)
	return os.IsExist(err), fileInfo.IsDir()
}
