package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionA(input string, isExample bool) (string, error) {
	sum := 0

	for _, entry := range strings.Split(input, ",") {
		ids := strings.Split(entry, "-")
		if len(ids) != 2 {
			return "", fmt.Errorf("expected ids to be length 2")
		}

		low, err := strconv.Atoi(ids[0])
		if err != nil {
			return "", err
		}
		high, err := strconv.Atoi(ids[1])
		if err != nil {
			return "", err
		}

		for idInt := low; idInt <= high; idInt++ {
			idStr := strconv.Itoa(idInt)
			if len(idStr)%2 == 0 {
				if idStr[:len(idStr)/2] == idStr[len(idStr)/2:] {
					sum += idInt
				}
			}
		}

	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(2, registry.A, SolutionA)
}
