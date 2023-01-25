// +build !js

package fs

import "os"

func LoadFile(name string) []byte {
	file, _ := os.ReadFile(string(name))
	return file
}