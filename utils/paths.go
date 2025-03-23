package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var cortexDir string

func EnsureCortexDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home directory: %v", err)
	}

	dirPath := filepath.Join(homeDir, ".cortex")

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create directory %s: %v", dirPath, err)
	}

	cortexDir = dirPath

	return nil
}

func GetCortexPath(path ...string) string {
	return filepath.Join(append([]string{cortexDir}, path...)...)
}
