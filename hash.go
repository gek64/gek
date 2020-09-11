package mypkg

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

// Checksum return checksum
func Checksum(mode string, file string) string {
	var result string
	switch mode {
	case "crc32":
		result = hashCRC32(file)
	case "md5":
		result = hashMD5(file)
	case "sha1":
		result = hashSHA1(file)
	case "sha256":
		result = hashSHA256(file)
	default:
		return ""
	}
	return result
}
func hashCRC32(file string) string {
	// open file
	fileObj, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// close file
	defer fileObj.Close()

	// copy filedata
	hash := crc32.NewIEEE()
	_, err = io.Copy(hash, fileObj)

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
func hashMD5(file string) string {
	// open file
	fileObj, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// close file
	defer fileObj.Close()

	// copy filedata
	hash := md5.New()
	_, err = io.Copy(hash, fileObj)

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
func hashSHA1(file string) string {
	// open file
	fileObj, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// close file
	defer fileObj.Close()

	// copy filedata
	hash := sha1.New()
	_, err = io.Copy(hash, fileObj)

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
func hashSHA256(file string) string {
	// open file
	fileObj, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// close file
	defer fileObj.Close()

	// copy filedata
	hash := sha256.New()
	_, err = io.Copy(hash, fileObj)

	// hex to string
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
