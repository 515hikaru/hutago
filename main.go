package main

import (
	"flag"
	"fmt"

	"github.com/515hikaru/mhugo/action"
	"github.com/515hikaru/mhugo/loader"
	"github.com/515hikaru/mhugo/parser"
)

func run() {
	recursive := flag.Bool("r", false, "search directories and their contents recursively")
	flag.Parse()
	targetDirectory := flag.Arg(0)
	fileNames, err := loader.ListDirectoryContents(targetDirectory, *recursive)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	mmap, err := parser.CreateMapTitleWithTag(fileNames, targetDirectory)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	action.PrintTags(mmap)
}

func main() {
	run()
}
