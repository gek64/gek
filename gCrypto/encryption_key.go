package gCrypto

import "bytes"

func KeyZeroPadding(key []byte, size int) []byte {
	if len(key) >= size {
		return key
	}
	return append(key, bytes.Repeat([]byte{0}, size-len(key))...)
}

func KeyCropping(key []byte, size int) []byte {
	if len(key) <= size {
		return key
	}
	return key[0:size]
}
