package gHash

import (
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func Hash(h hash.Hash, data []byte) []byte {
	h.Reset()
	h.Write(data)
	return h.Sum(nil)
}

func HashFile(h hash.Hash, file string) ([]byte, error) {
	// hash reset
	h.Reset()

	// check file
	fileInfo, err := os.Stat(file)
	if err != nil || fileInfo.IsDir() {
		return nil, fmt.Errorf("%s is not a valid file", file)
	}

	// open file
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(f)

	// copy data to hash
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
