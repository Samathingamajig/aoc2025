package day03

import (
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string) (string, error) {
	sum := 0

	for bank := range strings.SplitSeq(input, "\n") {
		maxLeftDigit := 0
		maxLeftDigitIndex := 0

		for i := 0; i < len(bank)-1; i++ {
			if int(bank[i]-'0') > maxLeftDigit {
				maxLeftDigit = int(bank[i] - '0')
				maxLeftDigitIndex = i
			}
		}

		maxRightDigit := 0
		for i := maxLeftDigitIndex + 1; i < len(bank); i++ {
			if int(bank[i]-'0') > maxRightDigit {
				maxRightDigit = int(bank[i] - '0')
			}
		}

		sum += maxLeftDigit*10 + maxRightDigit
	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(3, registry.A, SolutionA)
}
