package day05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

// delcared in a.go
// type idRange struct {
// 	min int
// 	max int
// }

func SolutionB(input string, isExample bool) (string, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return "", fmt.Errorf("expected 2 sections")
	}
	rangesRaw := sections[0]

	ranges := make([]idRange, 0)
	for rangeRaw := range strings.SplitSeq(rangesRaw, "\n") {
		rangeRawSplit := strings.Split(rangeRaw, "-")
		if len(rangeRawSplit) != 2 {
			return "", fmt.Errorf("expected 2 numbers in range")
		}
		low, err := strconv.Atoi(rangeRawSplit[0])
		if err != nil {
			return "", nil
		}
		high, err := strconv.Atoi(rangeRawSplit[1])
		if err != nil {
			return "", nil
		}

		ranges = append(ranges, idRange{low, high})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	changed := true
	for changed {
		changed = false
		for i := 0; i < len(ranges)-1; i++ {

			for j := len(ranges) - 1; j > i; j-- {
				if ranges[j].min <= ranges[i].max {
					ranges[i].max = max(ranges[i].max, ranges[j].max)
					ranges = append(ranges[:j], ranges[j+1:]...)
					changed = true
				}
			}
		}
	}

	numFresh := 0

	for _, theRange := range ranges {
		numFresh += theRange.max - theRange.min + 1
	}

	return strconv.Itoa(numFresh), nil
}

func init() {
	registry.Register(5, registry.B, SolutionB)
}
