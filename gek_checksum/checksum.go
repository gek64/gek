package gek_checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

// Checksum 校验和计算
func Checksum(mode string, fileURL string) (string, error) {
	var result string
	var err error
	switch mode {
	case "crc32":
		result, err = crc32Sum(fileURL)
	case "md5":
		result, err = md5Sum(fileURL)
	case "sha1":
		result, err = sha1Sum(fileURL)
	case "sha256":
		result, err = sha256Sum(fileURL)
	default:
		return "", fmt.Errorf("%s is not a valid mode", mode)
	}
	return result, err
}

// crc32Sum crc32校验和
func crc32Sum(fileURL string) (string, error) {
	// check fileURL is no link to a folder
	fileInfo, err := os.Stat(fileURL)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", fmt.Errorf("%s is a folder,and will be skiped", fileURL)
	}

	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		return "", err
	}

	// copy file data
	hash := crc32.NewIEEE()
	_, err = io.Copy(hash, fileObj)

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))

	// close file
	err = fileObj.Close()
	if err != nil {
		return "", err
	}
	return hashString, nil
}

// md5Sum md5校验和
func md5Sum(fileURL string) (string, error) {
	// check fileURL is no link to a folder
	fileInfo, err := os.Stat(fileURL)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", fmt.Errorf("%s is a folder,and will be skiped", fileURL)
	}

	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		return "", err
	}

	// copy file data
	hash := md5.New()
	_, err = io.Copy(hash, fileObj)

	// close file
	err = fileObj.Close()
	if err != nil {
		return "", err
	}

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString, nil
}

// sha1Sum sha1校验和
func sha1Sum(fileURL string) (string, error) {
	// check fileURL is no link to a folder
	fileInfo, err := os.Stat(fileURL)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", fmt.Errorf("%s is a folder,and will be skiped", fileURL)
	}

	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		return "", err
	}

	// copy file data
	hash := sha1.New()
	_, err = io.Copy(hash, fileObj)

	// close file
	err = fileObj.Close()
	if err != nil {
		return "", err
	}

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString, nil
}

// sha256Sum sha256校验和
func sha256Sum(fileURL string) (string, error) {
	// check fileURL is no link to a folder
	fileInfo, err := os.Stat(fileURL)
	if err != nil {
		return "", err
	}
	if fileInfo.IsDir() {
		return "", fmt.Errorf("%s is a folder,and will be skiped", fileURL)
	}

	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		return "", err
	}

	// copy file data
	hash := sha256.New()
	_, err = io.Copy(hash, fileObj)

	// close file
	err = fileObj.Close()
	if err != nil {
		return "", err
	}

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString, nil
}
