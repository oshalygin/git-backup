package main

import (
	"io/ioutil"
)

// GetDirectoriesInPath returns a list of directories to iterate through
func GetDirectoriesInPath(directoryPath string) []string {

	directories := make([]string, 0)
	files, _ := ioutil.ReadDir(directoryPath)

	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file.Name())
		}
	}

	return directories
}
