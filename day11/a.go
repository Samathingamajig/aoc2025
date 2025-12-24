package day00

import (
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string, isExample bool) (string, error) {
	edges := make(map[string][]string)

	for line := range strings.SplitSeq(input, "\n") {
		chunks := strings.Split(line, " ")
		edges[chunks[0][:len(chunks[0])-1]] = chunks[1:]
	}

	stack := make([]string, 0)
	endings := 0

	stack = append(stack, "you")

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == "out" {
			endings++
		} else {
			for _, next := range edges[curr] {
				stack = append(stack, next)
			}
		}
	}

	return strconv.Itoa(endings), nil
}

func init() {
	registry.Register(11, registry.A, SolutionA)
}
