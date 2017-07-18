package main

import (
	"errors"
)

func checkRequiredFlags(pathArgument string, path string) (string, error) {

	if pathArgument != "" && path != "" && pathArgument != path {
		err := "Cannot set a path flag AND an argument"
		return "", errors.New(err)
	}

	if path == "" && pathArgument == "" {
		err := `The "path" flag must be set`
		return "", errors.New(err)
	}

	if path != "" {
		return path, nil
	}

	return pathArgument, nil

}
