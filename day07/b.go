package day07

import (
	"strconv"

	"github.com/Samathingamajig/aoc2025/registry"
	"github.com/Samathingamajig/aoc2025/utils"
)

func SolutionB(input string, isExample bool) (string, error) {
	type cell struct {
		row        int
		col        int
		activated  bool
		isSplitter bool
		weight     int
	}

	grid := make([][]cell, 0)

	{
		r := 0
		c := 0

		for _, char := range input {
			if char == '\n' {
				r++
				c = 0
				continue
			}
			if r >= len(grid) {
				grid = append(grid, make([]cell, 0))
			}
			grid[r] = append(grid[r], cell{r, c, char == 'S', char == '^', utils.Bool2int(char == 'S')})
			c++
		}
	}

	timelines := 1

	for r := 1; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			above := grid[r-1][c]
			current := &grid[r][c]

			if !above.isSplitter && above.activated {
				current.activated = true
				current.weight += above.weight
				if current.isSplitter {
					timelines += current.weight
					if c-1 > 0 {
						grid[r][c-1].activated = true
						grid[r][c-1].weight += current.weight
					}
					if c+1 < len(grid[r]) {
						grid[r][c+1].activated = true
						grid[r][c+1].weight += current.weight
					}
				}
			}
		}
	}

	return strconv.Itoa(timelines), nil
}

func init() {
	registry.Register(7, registry.B, SolutionB)
}
