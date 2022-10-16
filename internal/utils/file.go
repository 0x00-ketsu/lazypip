package utils

import (
	"os"
)

// FilePathExist returns true if file is exist, else false
func FilePathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}

// CreateFileNotExist creates file if not exist
func CreateFileNotExist(filePath string) error {
	if !FilePathExist(filePath) {
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	}

	return nil
}
