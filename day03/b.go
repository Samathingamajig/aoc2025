package day03

import (
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string, isExample bool) (string, error) {
	sum := 0

	const allowedDigits = 12

	for bank := range strings.SplitSeq(input, "\n") {
		maxDigits := make([]int, allowedDigits)

		maxDigitsIndex := -1

		for digitIdx := range allowedDigits {
			for i := maxDigitsIndex + 1; i < len(bank)-allowedDigits+digitIdx+1; i++ {
				digit := int(bank[i] - '0')
				if digit > maxDigits[digitIdx] {
					maxDigits[digitIdx] = digit
					maxDigitsIndex = i
				}
			}
		}

		num := 0
		for _, val := range maxDigits {
			num = num*10 + val
		}

		sum += num
	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(3, registry.B, SolutionB)
}
