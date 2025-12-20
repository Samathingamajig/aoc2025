package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

func SolutionB(input string, isExample bool) (string, error) {
	splitLines := strings.Split(input, "\n")

	count := 0
	pos := 50
	for _, line := range splitLines {
		dirChar := line[0]
		dir := 0
		if dirChar == 'L' {
			dir = -1
		} else if dirChar == 'R' {
			dir = 1
		} else {
			return "", fmt.Errorf("didn't start with L or R")
		}

		movement, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", err
		}

		// pos = ((pos+movement*dir)%100 + 100) % 100
		// if pos == 0 {
		// 	count++
		// }

		for _ = range movement {
			pos = ((pos+dir)%100 + 100) % 100
			if pos == 0 {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}

// func main() {
// 	exampleInput := strings.TrimRight(readFilePanic(DAY_FOLDER, "example.input.txt"), "\n")
// 	exampleOutput := strings.TrimRight(readFilePanic(DAY_FOLDER, "example.output.txt"), "\n")
// 	realInput := strings.TrimRight(readFilePanic(DAY_FOLDER, "input.txt"), "\n")

// 	if RUN_EXAMPLE {
// 		fmt.Println("Runing example...")
// 		answer, err := solution(exampleInput)
// 		if err != nil {
// 			fmt.Println("❌ Error while running example:")
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		if answer == exampleOutput {
// 			fmt.Printf("✅ Example correct! `%s`\n", answer)
// 		} else {
// 			fmt.Println("❌ Example incorrect")
// 			fmt.Printf("Expected: `%s`\n", exampleOutput)
// 			fmt.Printf("Received: `%s`\n", answer)
// 			os.Exit(1)
// 		}
// 	} else {
// 		fmt.Println("Skipping example due to constants")
// 	}

// 	if RUN_REAL {
// 		fmt.Println("Runing real...")
// 		answer, err := solution(realInput)
// 		if err != nil {
// 			fmt.Println("❌ Error while running example:")
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		fmt.Printf("Calculated result: `%s`\n", answer)
// 	} else {
// 		fmt.Println("Skipping real problem due to constants")
// 	}
// }

// func readFilePanic(name ...string) string {
// 	raw, err := os.ReadFile(path.Join(name...))
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	return string(raw)
// }

func init() {
	registry.Register(1, registry.B, SolutionB)
}
