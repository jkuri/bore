package fs

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

// Exists check if file or directory exists on filesystem.
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

// MakeDir creates directory on filesystem.
func MakeDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0700)
}

// GetHomeDir returns full path to home directory
func GetHomeDir() (string, error) {
	return homedir.Dir()
}

// TempDir returns path to temporary directory
func TempDir() (string, error) {
	var tmp string
	if runtime.GOOS != "darwin" {
		tmp = os.TempDir()
	} else {
		tmp = "/tmp"
	}
	return ioutil.TempDir(tmp, "bore")
}
