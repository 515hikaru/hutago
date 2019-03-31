package action

import (
	"testing"

	"github.com/515hikaru/mhugo/parser"
)

func TestSimpleCountTag(t *testing.T) {
	tt := []parser.TitleAndTags{
		parser.TitleAndTags{
			Title: "title1",
			Tags:  []string{"foo", "bar"},
		},
	}
	m := countTag(tt)
	if m["foo"] != 1 {
		t.Errorf("m[\"foo\"] should be 1. got %d", m["foo"])
	}
	if m["bar"] != 1 {
		t.Errorf("m[\"bar\"] should be 1. got %d", m["bar"])
	}
}

func TestMultiCountTag(t *testing.T) {
	tt := []parser.TitleAndTags{
		parser.TitleAndTags{
			Title: "title1",
			Tags:  []string{"foo", "bar"},
		},
		parser.TitleAndTags{
			Title: "title2",
			Tags:  []string{"foo", "boo", "boo"},
		},
	}
	m := countTag(tt)
	if m["foo"] != 2 {
		t.Errorf("m[\"foo\"] should be 2. got %d", m["foo"])
	}
	if m["bar"] != 1 {
		t.Errorf("m[\"bar\"] should be 1. got %d", m["bar"])
	}
	if m["boo"] != 1 {
		t.Skipf("Feature of making tags uniq is not implemented.")
	}
}
