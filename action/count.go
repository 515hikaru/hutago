package action

import (
	"fmt"

	"github.com/515hikaru/mhugo/parser"
)

func countTag(m parser.PairTitleAndTags) parser.PairTagAndCount {
	countTagMap := make(parser.PairTagAndCount)
	for _, tags := range m {
		for _, tag := range tags {
			countTagMap[tag] += 1
		}
	}
	return countTagMap
}

func PrintTags(m parser.PairTitleAndTags) {
	c := countTag(m)
	for k, v := range c {
		fmt.Printf("%s:\t%d\n", k, v)
	}
}
