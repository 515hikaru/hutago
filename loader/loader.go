package loader

import "io/ioutil"

// ListDirectoryContents gets slice of file paths.
// If recursive is true, it searches file recursivly.
func ListDirectoryContents(directoryName string, _recursive bool) ([]string, error) {
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
