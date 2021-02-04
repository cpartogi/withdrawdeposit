package utils

import "os"

// IsFileorDirExist returns a boolean indicating whether file is exist or not
func IsFileorDirExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}
