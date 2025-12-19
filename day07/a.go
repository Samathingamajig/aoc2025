package day07

import (
	"strconv"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string) (string, error) {
	type cell struct {
		row        int
		col        int
		activated  bool
		isSplitter bool
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
			grid[r] = append(grid[r], cell{r, c, char == 'S', char == '^'})
			c++
		}
	}

	timesSplit := 0

	for r := 1; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			above := grid[r-1][c]
			current := &grid[r][c]
			if !above.isSplitter && above.activated {
				current.activated = true
				if current.isSplitter {
					timesSplit++
					if c-1 > 0 {
						grid[r][c-1].activated = true
					}
					if c+1 < len(grid[r]) {
						grid[r][c+1].activated = true
					}
				}
			}
		}
	}

	return strconv.Itoa(timesSplit), nil
}

func init() {
	registry.Register(7, registry.A, SolutionA)
}
