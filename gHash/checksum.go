package gHash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"hash/crc32"
	"hash/crc64"
)

func Crc32Sum(file string) ([]byte, error) {
	return HashFile(crc32.NewIEEE(), file)
}

func Crc64Sum(file string) ([]byte, error) {
	return HashFile(crc64.New(crc64.MakeTable(crc64.ECMA)), file)
}

func Md5Sum(file string) ([]byte, error) {
	return HashFile(md5.New(), file)
}

func Sha1Sum(file string) ([]byte, error) {
	return HashFile(sha1.New(), file)
}

func Sha256Sum(file string) ([]byte, error) {
	return HashFile(sha256.New(), file)
}

func Sha512Sum(file string) ([]byte, error) {
	return HashFile(sha512.New(), file)
}

func Blake2s256Sum(file string) ([]byte, error) {
	h, err := blake2s.New256(nil)
	if err != nil {
		return nil, err
	}
	return HashFile(h, file)
}

func Blake2b256Sum(file string) ([]byte, error) {
	h, err := blake2b.New256(nil)
	if err != nil {
		return nil, err
	}
	return HashFile(h, file)
}

func Blake2b384Sum(file string) ([]byte, error) {
	h, err := blake2b.New384(nil)
	if err != nil {
		return nil, err
	}
	return HashFile(h, file)
}

func Blake2b512Sum(file string) ([]byte, error) {
	h, err := blake2b.New512(nil)
	if err != nil {
		return nil, err
	}
	return HashFile(h, file)
}
