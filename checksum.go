package vivycore

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
	"io"
	"log"
	"os"
)

// Checksum 校验和计算
func Checksum(mode string, fileURL string) string {
	var result string
	switch mode {
	case "crc32":
		result = crc32Sum(fileURL)
	case "md5":
		result = md5Sum(fileURL)
	case "sha1":
		result = sha1Sum(fileURL)
	case "sha256":
		result = sha256Sum(fileURL)
	default:
		return ""
	}
	return result
}

// crc32Sum crc32校验和
func crc32Sum(fileURL string) string {
	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	// copy file data
	hash := crc32.NewIEEE()
	_, err = io.Copy(hash, fileObj)
	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	// close file
	err = fileObj.Close()
	if err != nil {
		log.Fatal(err)
	}
	return hashString
}

// md5Sum md5校验和
func md5Sum(fileURL string) string {
	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	// copy file data
	hash := md5.New()
	_, err = io.Copy(hash, fileObj)
	// close file
	err = fileObj.Close()
	if err != nil {
		log.Fatal(err)
	}
	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}

// sha1Sum sha1校验和
func sha1Sum(fileURL string) string {
	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	// copy file data
	hash := sha1.New()
	_, err = io.Copy(hash, fileObj)
	// close file
	err = fileObj.Close()
	if err != nil {
		log.Fatal(err)
	}
	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}

// sha256Sum sha256校验和
func sha256Sum(fileURL string) string {
	// open file
	fileObj, err := os.Open(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	// copy file data
	hash := sha256.New()
	_, err = io.Copy(hash, fileObj)
	// close file
	err = fileObj.Close()
	if err != nil {
		log.Fatal(err)
	}
	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
