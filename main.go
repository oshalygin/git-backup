package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
)

var path string
var directoryPath string

func init() {
	flag.StringVar(&path, "path", "", "Project repository location")
	flag.Parse()
}

func main() {

	pathArgument := flag.Arg(0)
	_, err := checkRequiredFlags(pathArgument, path)

	if err != nil {
		color.Red(err.Error())
		flag.PrintDefaults()

		os.Exit(1)
	}

	//GetDirectoriesInPath()

}
