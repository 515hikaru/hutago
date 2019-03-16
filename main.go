package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const yamlDelimiter = "---"

func takeYamlLines(buf []byte) []string {
	fileContent := string(buf)
	fileContentArray := strings.Split(fileContent, "\n")
	yamlLines := make([]string, 0)
	flg := false
	for _, line := range fileContentArray {
		if line == yamlDelimiter && flg == false {
			flg = true
			continue
		} else if line == yamlDelimiter && flg == true {
			break
		}

		yamlLines = append(yamlLines, line)
	}
	return yamlLines
}

func main() {
	buf, err := ioutil.ReadFile("test.md")
	if err != nil {
		panic(err)
	}
	yamlLines := takeYamlLines(buf)
	for i, item := range yamlLines {
		fmt.Printf("%d: %s\n", i, item)
	}
}
