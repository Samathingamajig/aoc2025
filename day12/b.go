package day12

import (
	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string, isExample bool) (string, error) {
	return "Merry Christmas!", nil
}

func init() {
	registry.Register(12, registry.B, SolutionB)
}
