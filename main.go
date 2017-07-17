package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "Project repository location")
}

func main() {

	err := checkRequiredFlags(path)
	if err != nil {
		color.Red("A project repository path MUST be provided.")
		flag.PrintDefaults()

		os.Exit(1)
	}

}
