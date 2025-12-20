package day06

import (
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string, isExample bool) (string, error) {
	lines := strings.Split(input, "\n")
	sums := make([]int, len(lines[0]))
	products := make([]int, len(lines[0]))
	for i := range products {
		products[i] = 1
	}
	result := 0

	arrIdx := 0
	for c := 0; c < len(lines[0]); c++ {
		allSpaces := true

		val := 0
		for r := 0; r < len(lines)-1; r++ {
			if lines[r][c] != ' ' {
				allSpaces = false
				val = val*10 + int(lines[r][c]-'0')
			}
		}

		if allSpaces {
			arrIdx += 1
		} else {
			sums[arrIdx] += val
			products[arrIdx] *= val
		}
	}

	arrIdx = 0
	for c := 0; c < len(lines[0]); c++ {
		value := lines[len(lines)-1][c]
		if value == '+' {
			result += sums[arrIdx]
			arrIdx++
		} else if value == '*' {
			result += products[arrIdx]
			arrIdx++
		}
	}

	return strconv.Itoa(result), nil
}

func init() {
	registry.Register(6, registry.B, SolutionB)
}
