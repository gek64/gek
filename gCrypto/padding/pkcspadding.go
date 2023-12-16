package padding

import "bytes"

func PKCSPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCSUnPadding(paddedText []byte) []byte {
	padding := paddedText[len(paddedText)-1]
	return paddedText[0:(len(paddedText) - int(padding))]
}
