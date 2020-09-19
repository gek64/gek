package gopkg

import (
	"io/ioutil"
	"os"
)

func CreateTempFile(dir string, pattern string, content string) (*os.File, error) {
	// Create temp file
	tmpFile, err := ioutil.TempFile(dir, pattern)
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

func CreateTempDir(dir string, pattern string) (string, error) {
	// Create temp dir
	tmpDir, err := ioutil.TempDir(dir, pattern)
	if err != nil {
		return "", err
	}

	// Remember to clean up the file afterwards
	// defer os.RemoveAll(tmpDir)

	return tmpDir, nil
}
