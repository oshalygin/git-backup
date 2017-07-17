package main

import (
	"errors"
)

func checkRequiredFlags(path string) error {

	if path == "" {
		err := `The "path" flag must be set`
		return errors.New(err)
	}
	return nil

}
