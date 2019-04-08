package parser

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type ArticleHeader struct {
	Title string
	Date  string
	Draft bool
	Tags  []string
}

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
		if flg == true {
			yamlLines = append(yamlLines, line)
		}
	}
	return yamlLines
}

func parseTags(yamlLines []string) (ArticleHeader, error) {
	yamlContent := (strings.Join(yamlLines, "\n"))
	yamlBytes := []byte(yamlContent)
	h := ArticleHeader{}
	err := yaml.Unmarshal(yamlBytes, &h)
	if err != nil {
		return ArticleHeader{}, err
	}
	return h, nil
}

func CreateHeaders(filePaths []string, parentPath string) ([]ArticleHeader, error) {
	headers := make([]ArticleHeader, len(filePaths))
	for _, filePath := range filePaths {
		buf, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		yamlLines := takeYamlLines(buf)
		h, err := parseTags(yamlLines)
		if err != nil {
			return nil, err
		}
		headers = append(headers, h)
	}
	return headers, nil
}
