package day08

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string) (string, error) {
	type Box struct {
		x         float64
		y         float64
		z         float64
		circuitId int
	}

	nextCircuitId := 0
	boxes := make([]Box, 0)

	for line := range strings.SplitSeq(input, "\n") {
		numsRaw := strings.Split(line, ",")
		x, err := strconv.Atoi(numsRaw[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(numsRaw[1])
		if err != nil {
			return "", err
		}
		z, err := strconv.Atoi(numsRaw[2])
		if err != nil {
			return "", err
		}

		boxes = append(boxes, Box{float64(x), float64(y), float64(z), nextCircuitId})
		nextCircuitId++
	}

	pairs := make([][]*Box, 0)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			pair := make([]*Box, 2)
			pair[0] = &boxes[i]
			pair[1] = &boxes[j]
			pairs = append(pairs, pair)
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		p1a := pairs[i][0]
		p1b := pairs[i][1]
		dist1 := math.Pow(p1a.x-p1b.x, 2) + math.Pow(p1a.y-p1b.y, 2) + math.Pow(p1a.z-p1b.z, 2)
		p2a := pairs[j][0]
		p2b := pairs[j][1]
		dist2 := math.Pow(p2a.x-p2b.x, 2) + math.Pow(p2a.y-p2b.y, 2) + math.Pow(p2a.z-p2b.z, 2)
		return dist1 < dist2
	})

	for _, pair := range pairs {
		a := pair[0]
		b := pair[1]
		if a.circuitId != b.circuitId {
			oldId := b.circuitId
			newId := a.circuitId
			diffFound := false
			for j := 0; j < len(boxes); j++ {
				if boxes[j].circuitId == oldId {
					boxes[j].circuitId = newId
				} else if boxes[j].circuitId != newId {
					diffFound = true
				}
			}

			if !diffFound {
				product := int(a.x) * int(b.x)
				return strconv.Itoa(product), nil
			}
		}
	}

	return "", fmt.Errorf("Never joined the structure.")

}

func init() {
	registry.Register(8, registry.B, SolutionB)
}
