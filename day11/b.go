package day00

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func getIncomingEdges(edges map[string][]string) map[string][]string {
	incomingEdges := make(map[string][]string)
	incomingEdges["svr"] = make([]string, 0)
	for source, dests := range edges {
		for _, dest := range dests {
			sources, ok := incomingEdges[dest]
			if !ok {
				sources = make([]string, 0)
			}

			sources = append(sources, source)

			incomingEdges[dest] = sources
		}
	}

	return incomingEdges
}

func getOrdering(edges map[string][]string) ([]string, error) {
	// https://en.wikipedia.org/wiki/Topological_sorting#Kahn's_algorithm

	ordering := make([]string, 0)
	noIncomingEdges := make([]string, 0)
	incomingEdges := getIncomingEdges(edges)

	for dest, sources := range incomingEdges {
		if len(sources) == 0 {
			noIncomingEdges = append(noIncomingEdges, dest)
		}
	}

	for len(noIncomingEdges) > 0 {
		node := noIncomingEdges[len(noIncomingEdges)-1]
		noIncomingEdges = noIncomingEdges[:len(noIncomingEdges)-1]
		ordering = append(ordering, node)

		for _, dest := range edges[node] {
			// https://zetcode.com/golang/filter-slice/
			incomingEdgesOfDest := incomingEdges[dest]
			n := 0
			for _, source := range incomingEdgesOfDest {
				if source != node {
					incomingEdgesOfDest[n] = source
					n++
				}
			}
			incomingEdges[dest] = incomingEdgesOfDest[:n]

			if n == 0 {
				noIncomingEdges = append(noIncomingEdges, dest)
			}
		}
	}

	if len(ordering) != len(edges) {
		return nil, fmt.Errorf("Could not make ordering, there is a cycle")
	}

	return ordering, nil
}

func SolutionB(input string, isExample bool) (string, error) {
	edges := make(map[string][]string)

	edges["out"] = make([]string, 0)
	for line := range strings.SplitSeq(input, "\n") {
		chunks := strings.Split(line, " ")
		edges[chunks[0][:len(chunks[0])-1]] = chunks[1:]
	}

	ordering, err := getOrdering(edges)
	if err != nil {
		return "", err
	}
	incomingEdges := getIncomingEdges(edges)

	type State struct {
		raw       int
		afterDac  int
		afterFft  int
		afterBoth int
	}

	states := make(map[string]State)

	for _, node := range ordering {
		state := State{0, 0, 0, 0}
		if node == "svr" {
			state.raw = 1
		}

		for _, source := range incomingEdges[node] {
			sourceState := states[source]
			state.raw += sourceState.raw
			state.afterDac += sourceState.afterDac
			state.afterFft += sourceState.afterFft
			state.afterBoth += sourceState.afterBoth
		}

		if node == "dac" {
			if state.afterFft > 0 {
				state.afterBoth = state.afterFft
				state.raw = 0
				state.afterFft = 0
				state.afterDac = 0
			} else {
				state.afterDac = state.raw
				state.raw = 0
			}
		} else if node == "fft" {
			if state.afterDac > 0 {
				state.afterBoth = state.afterDac
				state.raw = 0
				state.afterFft = 0
				state.afterDac = 0
			} else {
				state.afterFft = state.raw
				state.raw = 0
			}
		}

		states[node] = state
	}

	return strconv.Itoa(states["out"].afterBoth), nil
}

func init() {
	registry.Register(11, registry.B, SolutionB)
}
