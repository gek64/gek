package filesystem

import "os"

// Mkdir Create path
func Mkdir(path string) error {
	return os.Mkdir(path, 0755)
}

// MkdirAll Create all path
func MkdirAll(path string) error {
	return os.MkdirAll(path, 0755)
}
