package padding

import "bytes"

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padText...)
}

func ZeroUnPadding(paddedText []byte) []byte {
	return bytes.TrimRightFunc(paddedText, func(r rune) bool {
		return r == rune(0)
	})
}
