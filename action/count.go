package action

import (
	"fmt"
	"sort"
	"strings"

	"github.com/515hikaru/hutago/parser"
)

// PairTagAndCount holds the tag name and its frequency
type PairTagAndCount map[string]int

type tc struct {
	Tag   string
	Count int
}

func countTag(headers []parser.ArticleHeader) PairTagAndCount {
	countTagMap := make(PairTagAndCount)
	for _, header := range headers {
		for _, tag := range header.Tags {
			countTagMap[tag]++
		}
	}
	return countTagMap
}

func getMaxLength(c PairTagAndCount) int {
	var maxLength int
	for k := range c {
		if l := len(k); l >= maxLength {
			maxLength = l
		}
	}
	return maxLength
}

func sortCount(c PairTagAndCount) []tc {
	var tcs []tc
	for k, v := range c {
		t := tc{Tag: k, Count: v}
		tcs = append(tcs, t)
	}
	sort.Slice(tcs, func(i, j int) bool {
		return tcs[i].Count > tcs[j].Count
	})
	return tcs
}

// PrintTags prints tag name and the tag frequency
func PrintTags(headers []parser.ArticleHeader) {
	c := countTag(headers)
	sorted := sortCount(c)
	width := getMaxLength(c)
	for _, t := range sorted {
		diff := width - len(t.Tag)
		interval := strings.Repeat(" ", diff)
		fmt.Printf("%s%s\t%d\n", t.Tag, interval, t.Count)
	}
}
