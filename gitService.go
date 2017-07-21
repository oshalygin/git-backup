package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

// PullLatest retrieves the most recent code from the codebase
func PullLatest(directory string) bool {

	color.Green("PULL: %s", directory)

	commandName := "git"
	args := []string{"pull", "origin", "master"}

	command := exec.Command(commandName, args...)
	command.Dir = directory

	commandOutput := &bytes.Buffer{}
	commandErrorOutput := &bytes.Buffer{}
	command.Stdout = commandOutput
	command.Stderr = commandErrorOutput

	err := command.Run()

	if commandErrorOutput != nil {
		fmt.Print(string(commandErrorOutput.Bytes()))
	}
	if commandOutput != nil {
		fmt.Print(string(commandOutput.Bytes()))
	}

	if err != nil {
		fmt.Println(err.Error())
		color.Red("PULL ERROR: %s", directory)
		return false
	}

	color.Green("PULL SUCCESS: %s", directory)
	return true

}
