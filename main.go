package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
)

var path string

func init() {
	flag.StringVar(&path, "path", "", "Project repository location")
	flag.Parse()
}

func main() {

	pathArgument := flag.Arg(0)
	directoryPath, err := checkRequiredFlags(pathArgument, path)

	if err != nil {
		color.Red(err.Error())
		flag.PrintDefaults()

		os.Exit(1)
	}

	directories := GetDirectoriesInPath(directoryPath)
	for _, dir := range directories {
		FetchLatest(dir)
		PullLatest(dir)
		PushLatest(dir)
	}

}
