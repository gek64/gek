package gCrypto

import (
	"crypto/md5"
	"github.com/tink-crypto/tink-go/v2/aead/subtle"
)

type ChaCha20Poly1305 struct {
	Key            []byte
	AssociatedData []byte
}

func NewChaCha20Poly1305(key []byte, associatedDataSize uint) *ChaCha20Poly1305 {
	// associated data size
	if associatedDataSize > md5.Size {
		associatedDataSize = md5.Size
	}
	// use md5 to generate associated data
	associatedDataCandidate := md5.Sum(key)
	return &ChaCha20Poly1305{
		Key:            key,
		AssociatedData: associatedDataCandidate[0:associatedDataSize],
	}
}

func (c *ChaCha20Poly1305) Encrypt(plaintext []byte) ([]byte, error) {
	ca, err := subtle.NewChaCha20Poly1305(c.Key)
	if err != nil {
		return nil, err
	}
	return ca.Encrypt(plaintext, c.AssociatedData)
}

func (c *ChaCha20Poly1305) Decrypt(ciphertext []byte) ([]byte, error) {
	ca, err := subtle.NewChaCha20Poly1305(c.Key)
	if err != nil {
		return nil, err
	}
	return ca.Decrypt(ciphertext, c.AssociatedData)
}
