package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/515hikaru/mhugo/action"
	"github.com/515hikaru/mhugo/loader"
	"github.com/515hikaru/mhugo/parser"
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
		fmt.Fprintf(os.Stderr, "Specify directory name\n")
		os.Exit(failCode)
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
