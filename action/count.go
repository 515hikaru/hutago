package action

import (
	"fmt"
	"sort"
	"strings"

	"github.com/515hikaru/mhugo/parser"
)

type tc struct {
	Tag   string
	Count int
}

func countTag(m parser.PairTitleAndTags) parser.PairTagAndCount {
	countTagMap := make(parser.PairTagAndCount)
	for _, tags := range m {
		for _, tag := range tags {
			countTagMap[tag] += 1
		}
	}
	return countTagMap
}

func sortCount(c parser.PairTagAndCount) ([]tc, int) {
	var tcs []tc
	var maxLength int
	for k, v := range c {
		if len(k) >= maxLength {
			maxLength = len(k)
		}
		t := tc{Tag: k, Count: v}
		tcs = append(tcs, t)
	}
	sort.Slice(tcs, func(i, j int) bool {
		return tcs[i].Count > tcs[j].Count
	})
	return tcs, maxLength
}

func PrintTags(m parser.PairTitleAndTags) {
	c := countTag(m)
	sorted, width := sortCount(c)
	for _, t := range sorted {
		diff := width - len(t.Tag)
		interval := strings.Repeat(" ", diff)
		fmt.Printf("%s%s\t%d\n", t.Tag, interval, t.Count)
	}
}
