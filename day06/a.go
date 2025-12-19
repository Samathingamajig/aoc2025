package day06

import (
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string) (string, error) {
	sums := make([]int, 0)
	products := make([]int, 0)
	result := 0

	for line := range strings.SplitSeq(input, "\n") {
		i := 0
		for value := range strings.SplitSeq(line, " ") {
			if value == "" {
				continue
			}

			if value == "+" {
				result += sums[i]
			} else if value == "*" {
				result += products[i]
			} else {
				num, err := strconv.Atoi(value)
				if err != nil {
					return "", err
				}
				if i >= len(sums) {
					sums = append(sums, 0)
				}
				if i >= len(products) {
					products = append(products, 1)
				}

				sums[i] += num
				products[i] *= num
			}

			i++
		}
	}

	return strconv.Itoa(result), nil
}

func init() {
	registry.Register(6, registry.A, SolutionA)
}
