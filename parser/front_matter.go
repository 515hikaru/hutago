package parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type ArticleHeader struct {
	Title string
	Draft bool
	Tag   []string
}

type PairTitleAndTags map[string][]string
type PairTagAndCount map[string]int

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

func CreateMapTitleWithTag(fileNames []string, parentPath string) (PairTitleAndTags, error) {
	headers := make([]ArticleHeader, len(fileNames))
	for _, fileName := range fileNames {
		filePath := path.Join(parentPath, fileName)
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		yamlLines := takeYamlLines(buf)
		h := parseTags(yamlLines)
		headers = append(headers, h)
	}
	m := getTagMap(headers)
	return m, nil
}
