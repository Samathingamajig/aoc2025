package day09

import (
	"math"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string, isExample bool) (string, error) {
	type Cell struct {
		x, y float64
	}

	cells := make([]Cell, 0)
	for line := range strings.SplitSeq(input, "\n") {
		nums := strings.Split(line, ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			return "", err
		}
		cells = append(cells, Cell{float64(x), float64(y)})
	}

	maxArea := 0
	for i := 0; i < len(cells); i++ {
		a := cells[i]
		for j := i + 1; j < len(cells); j++ {
			b := cells[j]
			area := math.Abs(a.x-b.x+1) * math.Abs(a.y-b.y+1)
			maxArea = max(maxArea, int(area))
		}
	}

	return strconv.Itoa(maxArea), nil
}

func init() {
	registry.Register(9, registry.A, SolutionA)
}
