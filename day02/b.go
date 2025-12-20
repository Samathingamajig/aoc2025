package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string, isExample bool) (string, error) {
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

			success := false
			for myLen := 1; myLen < len(idStr) && !success; myLen++ {
				if len(idStr)%myLen == 0 {
					valid := true
					for i := 1; i < len(idStr)/myLen && valid; i++ {
						if idStr[:myLen] != idStr[i*myLen:(i+1)*myLen] {
							valid = false
						}
					}
					success = valid
				}
			}

			if success {
				sum += idInt
			}
		}

	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(2, registry.B, SolutionB)
}
