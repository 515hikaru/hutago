package action

import (
	"fmt"
	"sort"
	"strings"

	"github.com/515hikaru/mhugo/parser"
)

type PairTagAndCount map[string]int

type tc struct {
	Tag   string
	Count int
}

func countTag(tt []parser.TitleAndTags) PairTagAndCount {
	countTagMap := make(PairTagAndCount)
	for _, t := range tt {
		for _, tag := range t.Tags {
			countTagMap[tag] += 1
		}
	}
	return countTagMap
}

func sortCount(c PairTagAndCount) ([]tc, int) {
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

func PrintTags(tt []parser.TitleAndTags) {
	c := countTag(tt)
	sorted, width := sortCount(c)
	for _, t := range sorted {
		diff := width - len(t.Tag)
		interval := strings.Repeat(" ", diff)
		fmt.Printf("%s%s\t%d\n", t.Tag, interval, t.Count)
	}
}
