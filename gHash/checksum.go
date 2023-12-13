package gHash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
	"log"
	"os"
)

func Crc32Sum(file string) ([]byte, error) {
	return Hash(crc32.NewIEEE(), file)
}

func Crc64Sum(file string) ([]byte, error) {
	return Hash(crc64.New(crc64.MakeTable(crc64.ECMA)), file)
}

func Md5Sum(file string) ([]byte, error) {
	return Hash(md5.New(), file)
}

func Sha1Sum(file string) ([]byte, error) {
	return Hash(sha1.New(), file)
}

func Sha256Sum(file string) ([]byte, error) {
	return Hash(sha256.New(), file)
}

func Sha512Sum(file string) ([]byte, error) {
	return Hash(sha512.New(), file)
}

func Blake2s256Sum(file string) ([]byte, error) {
	h, err := blake2s.New256(nil)
	if err != nil {
		return nil, err
	}
	return Hash(h, file)
}

func Blake2b256Sum(file string) ([]byte, error) {
	h, err := blake2b.New256(nil)
	if err != nil {
		return nil, err
	}
	return Hash(h, file)
}

func Blake2b384Sum(file string) ([]byte, error) {
	h, err := blake2b.New384(nil)
	if err != nil {
		return nil, err
	}
	return Hash(h, file)
}

func Blake2b512Sum(file string) ([]byte, error) {
	h, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	return Hash(h, file)
}

func Hash(h hash.Hash, file string) ([]byte, error) {
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
