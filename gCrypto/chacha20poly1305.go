package gCrypto

import (
	"github.com/tink-crypto/tink-go/v2/aead/subtle"
	"hash"
)

type ChaCha20Poly1305 struct {
	Key            []byte
	AssociatedData []byte
}

func NewChaCha20Poly1305(key []byte, associatedData []byte) *ChaCha20Poly1305 {
	return &ChaCha20Poly1305{
		Key:            key,
		AssociatedData: associatedData,
	}
}

// NewChaCha20Poly1305WithHashAD use hash as associated data, associated data will not be encrypted
func NewChaCha20Poly1305WithHashAD(key []byte, h hash.Hash) *ChaCha20Poly1305 {
	h.Reset()
	h.Write(key)
	return NewChaCha20Poly1305(key, h.Sum(nil))
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
