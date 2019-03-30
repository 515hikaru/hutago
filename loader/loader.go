package loader

import (
	"io/ioutil"
	"path"
)

// ListDirectoryContents gets slice of file paths.
// If recursive is true, it searches file recursivly.
func ListDirectoryContents(directoryName string, _recursive bool) ([]string, error) {
	items, err := ioutil.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}
	var files []string

	for _, item := range items {
		if item.IsDir() {
			// TODO: impl recursive == true
			continue
		} else {
			fileName := item.Name()
			if isMarkdown(fileName) {
				files = append(files, fileName)
			}
		}
	}
	return files, nil
}

func isMarkdown(fileName string) bool {
	if path.Ext(fileName) == ".md" {
		return true
	} else if path.Ext(fileName) == ".markdown" {
		return true
	}
	return false
}
