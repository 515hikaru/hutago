package loader

import (
	"io/ioutil"
	"path"
)

// ListDirectoryContents gets slice of file paths.
// If recursive is true, it searches file recursivly.
func ListDirectoryContents(directoryName string, recursive bool) ([]string, error) {
	items, err := ioutil.ReadDir(directoryName)
	if err != nil {
		return nil, err
	}
	var files []string

	for _, item := range items {
		if item.IsDir() {
			if recursive == false {
				continue
			}
			targetPath := path.Join(directoryName, item.Name())
			childFiles, err := ListDirectoryContents(targetPath, recursive)
			if err != nil {
				return nil, err
			}
			for _, file := range childFiles {
				files = append(files, file)
			}
		} else {
			filePath := path.Join(directoryName, item.Name())
			if isMarkdown(filePath) {
				files = append(files, filePath)
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
