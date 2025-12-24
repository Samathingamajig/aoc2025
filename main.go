package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	_ "github.com/Samathingamajig/aoc2025/internal"
	"github.com/Samathingamajig/aoc2025/registry"
	"github.com/Samathingamajig/aoc2025/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: main.go 01/a")
	}

	dayAndPart := strings.Split(os.Args[1], "/")
	if len(dayAndPart) != 2 {
		log.Fatalln("Usage: main.go 01/a")
	}

	dayRaw := dayAndPart[0]
	partRaw := strings.ToLower(dayAndPart[1])

	day, err := strconv.Atoi(strings.TrimLeft(dayRaw, "0"))
	if err != nil {
		log.Println("Day could not be parsed as an integer")
		log.Fatalln(err)
	}
	_ = day

	if partRaw != "a" && partRaw != "b" {
		log.Fatalln("Part must be \"a\" or \"b\"")
	}

	part := registry.Part(partRaw)

	solutionFunc, ok := registry.Get(day, part)
	if !ok {
		log.Fatalf("%02d/%s doesn't exist\n", day, part)
	}

	runner(day, part, solutionFunc, true, true)
}

func runner(day int, part registry.Part, solutionFunc registry.SolutionFunc, runExample, runReal bool) {
	paddedDay := fmt.Sprintf("day%02d", day)
	exampleInputRaw, err := os.ReadFile(path.Join(paddedDay, fmt.Sprintf("%s.example.input.txt", part)))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			exampleInputRaw, err = os.ReadFile(path.Join(paddedDay, "example.input.txt"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	exampleInput := strings.TrimRight(string(exampleInputRaw), "\n")
	exampleOutput := strings.TrimRight(utils.ReadFilePanic(paddedDay, fmt.Sprintf("%s.example.output.txt", part)), "\n")
	realInput := strings.TrimRight(utils.ReadFilePanic(paddedDay, "input.txt"), "\n")

	files := []struct {
		name string
		data string
	}{
		{"example.input.txt", exampleInput},
		{fmt.Sprintf("%s.example.output.txt", part), exampleOutput},
		{"input.txt", realInput},
	}

	badFile := false
	for _, file := range files {
		if file.data == "!!! CHANGE ME !!!" {
			log.Printf("⚠️ You need to edit %s/%s\n", paddedDay, file.name)
			badFile = true
		}
	}
	if badFile {
		os.Exit(1)
	}

	if runExample {
		fmt.Println("Running example...")
		answer, err := solutionFunc(exampleInput, true)
		if err != nil {
			fmt.Println("❌ Error while running example:")
			fmt.Println(err)
			os.Exit(1)
		}

		if answer == exampleOutput {
			fmt.Printf("✅ Example correct! `%s`\n", answer)
		} else {
			fmt.Println("❌ Example incorrect")
			fmt.Printf("Expected: `%s`\n", exampleOutput)
			fmt.Printf("Received: `%s`\n", answer)
			os.Exit(1)
		}
	} else {
		fmt.Println("Skipping example due to parameter")
	}

	if runReal {
		fmt.Println("Running real...")
		answer, err := solutionFunc(realInput, false)
		if err != nil {
			fmt.Println("❌ Error while running example:")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Calculated result: `%s`\n", answer)
	} else {
		fmt.Println("Skipping real problem due to parameter")
	}
}
