package main

import (
	"os"
	"path/filepath"
)

// Retrieve a list of directories to iterate through
func GetDirectoriesInPath(directoryPath string) {
	filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			println(info.Name())
		}
		return nil
	})
}
