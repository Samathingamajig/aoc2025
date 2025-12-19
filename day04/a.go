package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string) (string, error) {
	grid := strings.Split(input, "\n")
	grid = append([]string{strings.Repeat(".", len(grid[0]))}, grid...)
	grid = append(grid, strings.Repeat(".", len(grid[0])))

	for i := range grid {
		grid[i] = fmt.Sprintf(".%s.", grid[i])
	}

	count := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == '@' {
				neighbors := -1
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if grid[i+di][j+dj] == '@' {
							neighbors++
						}
					}
				}
				if neighbors < 4 {
					count++
				}
			}

		}
	}

	return strconv.Itoa(count), nil
}

func init() {
	registry.Register(4, registry.A, SolutionA)
}
