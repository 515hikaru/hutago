package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

const yamlDelimiter = "---"

type ArticleHeader struct {
	Title string
	Draft bool
	Tag   []string
}

type PairTitleAndTags map[string][]string
type PairTagAndCount map[string]int

func listDirectoryContens(directoryName string, _recursive bool) ([]string, error) {
	items, err := ioutil.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}
	// TODO: calc size for `recursive` true case
	var files []string

	for _, item := range items {
		if item.IsDir() {
			// TODO: impl recursive == true
			continue
		} else {
			files = append(files, item.Name())
		}
	}
	return files, nil
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
		fmt.Fprintf(os.Stderr, "error: %v", err)
	}
	return h
}

func getTagMap(headers []ArticleHeader) PairTitleAndTags {
	m := make(map[string][]string)
	for _, header := range headers {
		m[header.Title] = header.Tag
	}
	return m
}

func countTag(m PairTitleAndTags) PairTagAndCount {
	countTagMap := make(PairTagAndCount)
	for _, tags := range m {
		for _, tag := range tags {
			countTagMap[tag] += 1
		}
	}
	return countTagMap
}

func printTags(c PairTagAndCount) {
	for k, v := range c {
		fmt.Printf("%s:\t%d\n", k, v)
	}
}

func main() {
	recursive := flag.Bool("r", false, "search directories and their contents recursively")
	flag.Parse()
	targetDirectory := flag.Arg(0)
	fileNames, err := listDirectoryContens(targetDirectory, *recursive)
	if err != nil {
		panic(err)
	}
	headers := make([]ArticleHeader, len(fileNames))
	for _, fileName := range fileNames {
		filePath := path.Join(targetDirectory, fileName)
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		yamlLines := takeYamlLines(buf)
		h := parseTags(yamlLines)
		headers = append(headers, h)
	}
	if err != nil {
		fmt.Println("boo")
		panic(err)
	}
	m := getTagMap(headers)
	c := countTag(m)
	printTags(c)
}
