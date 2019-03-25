package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	buf, err := ioutil.ReadFile("test.md")
	if err != nil {
		panic(err)
	}
	yamlLines := takeYamlLines(buf)
	h := parseTags(yamlLines)
	m := getTagMap([]ArticleHeader{h})
	c := countTag(m)
	printTags(c)
}
