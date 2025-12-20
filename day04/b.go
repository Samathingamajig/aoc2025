package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string, isExample bool) (string, error) {
	gridRaw := strings.Split(input, "\n")
	gridRaw = append([]string{strings.Repeat(".", len(gridRaw[0]))}, gridRaw...)
	gridRaw = append(gridRaw, strings.Repeat(".", len(gridRaw[0])))

	for i := range gridRaw {
		gridRaw[i] = fmt.Sprintf(".%s.", gridRaw[i])
	}
	grid := make([][]byte, 0, len(gridRaw))
	for _, row := range gridRaw {
		grid = append(grid, []byte(row))
	}

	count := 0
	changed := true
	for changed {
		changed = false
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
						grid[i][j] = '.'
						changed = true
					}
				}

			}
		}
	}

	return strconv.Itoa(count), nil
}

func init() {
	registry.Register(4, registry.B, SolutionB)
}
