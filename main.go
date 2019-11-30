package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/515hikaru/hutago/action"
	"github.com/515hikaru/hutago/loader"
	"github.com/515hikaru/hutago/parser"
)

const (
	failCode    = 1
	successCode = 0
)

func run() {
	recursive := flag.Bool("r", false, "search directories and their contents recursively")
	flag.Parse()
	targetDirectory := flag.Arg(0)
	if targetDirectory == "" {
		targetDirectory = "contents"
		*recursive = true
	}
	fileNames, err := loader.ListDirectoryContents(targetDirectory, *recursive)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(failCode)
	}
	if len(fileNames) == 0 {
		fmt.Fprintf(os.Stderr, "%s has no markdown files.\n", targetDirectory)
		os.Exit(failCode)
	}
	mmap, err := parser.CreateHeaders(fileNames, targetDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(failCode)
	}
	action.PrintTags(mmap)
}

func main() {
	run()
	os.Exit(successCode)
}
