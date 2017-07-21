package main

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func printOutput(buffer ...*Buffer) {

	for _, output := range buffer {
		if output != nil {
			fmt.Print(string(output.Bytes()))
		}
	}
}

// FetchLatest fetches all remote branches
func FetchLatest(directory string) bool {

	color.Green("FETCH: %s", directory)

	commandName := "git"
	args := []string{"fetch", "--all"}

	command := exec.Command(commandName, args...)
	command.Dir = directory

	commandOutput := &bytes.Buffer{}
	commandErrorOutput := &bytes.Buffer{}
	command.Stdout = commandOutput
	command.Stderr = commandErrorOutput

	err := command.Run()

	printOutput(commandOutput, commandErrorOutput)

	if err != nil {
		fmt.Println(err.Error())
		color.Red("FETCH ERROR: %s", directory)
		return false
	}

	color.Green("FETCH SUCCESS: %s", directory)
	return true

}

// PullLatest retrieves the most recent code from the codebase
func PullLatest(directory string) bool {

	color.Green("PULL: %s", directory)

	commandName := "git"
	args := []string{"pull", "--all"}

	command := exec.Command(commandName, args...)
	command.Dir = directory

	commandOutput := &bytes.Buffer{}
	commandErrorOutput := &bytes.Buffer{}
	command.Stdout = commandOutput
	command.Stderr = commandErrorOutput

	err := command.Run()

	printOutput(commandOutput, commandErrorOutput)

	if err != nil {
		fmt.Println(err.Error())
		color.Red("PULL ERROR: %s", directory)
		return false
	}

	color.Green("PULL SUCCESS: %s", directory)
	return true

}
