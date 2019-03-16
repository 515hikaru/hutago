package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"log"
)

const yamlDelimiter = "---"

type ArticleHeader struct {
	Title string
	Draft bool
	Tag   []string
}

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

func parseTags(yamlLines []string) ArticleHeader {
	yamlContent := (strings.Join(yamlLines, "\n"))
	yamlBytes := []byte(yamlContent)
	h := ArticleHeader{}
	err := yaml.Unmarshal(yamlBytes, &h)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return h
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
	h := parseTags(yamlLines)
	fmt.Println(h)
}
