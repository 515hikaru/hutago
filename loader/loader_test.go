package loader

import (
	"testing"
)

func TestIsMarkdown(t *testing.T) {
	validMdFileName := "foo.md"
	validMarkdownFileName := "foo.markdown"
	invalidMdFileName := "foo.d"
	directoryName := "foo"

	if isMarkdown(validMdFileName) != true {
		t.Errorf("isMarkdown(\"foo.md\") = %v, want true\n", isMarkdown(validMdFileName))
	}

	if isMarkdown(validMarkdownFileName) != true {
		t.Errorf("isMarkdown(\"foo.markdown\") = %v, want true\n", isMarkdown(validMarkdownFileName))
	}
	if isMarkdown(invalidMdFileName) != false {
		t.Errorf("isMarkdown(\"foo.d\") = %v, want false\n", isMarkdown(invalidMdFileName))
	}
	if isMarkdown(directoryName) != false {
		t.Errorf("isMarkdown(\"foo\") = %v, want false\n", isMarkdown(directoryName))
	}
}
