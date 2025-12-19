package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

type idRange struct {
	min int
	max int
}

func SolutionA(input string) (string, error) {
	sections := strings.Split(input, "\n\n")
	if len(sections) != 2 {
		return "", fmt.Errorf("expected 2 sections")
	}
	rangesRaw := sections[0]
	idsRaw := sections[1]

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

	// fmt.Printf("%+v\n", ranges)

	numFresh := 0
	for idRaw := range strings.SplitSeq(idsRaw, "\n") {
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			return "", err
		}

		for _, theRange := range ranges {
			if theRange.min <= id && id <= theRange.max {
				numFresh++
				break
			}
		}
	}

	return strconv.Itoa(numFresh), nil
}

func init() {
	registry.Register(5, registry.A, SolutionA)
}
