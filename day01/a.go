package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string, isExample bool) (string, error) {
	splitLines := strings.Split(input, "\n")

	count := 0
	pos := 50
	for _, line := range splitLines {
		dirChar := line[0]
		dir := 0
		if dirChar == 'L' {
			dir = -1
		} else if dirChar == 'R' {
			dir = 1
		} else {
			return "", fmt.Errorf("didn't start with L or R")
		}

		movement, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}

		pos = ((pos+movement*dir)%100 + 100) % 100
		if pos == 0 {
			count++
		}

		// for _ := range count {
		// 	pos += dir
		// }
	}

	return strconv.Itoa(count), nil
}

func init() {
	registry.Register(1, registry.A, SolutionA)
}
