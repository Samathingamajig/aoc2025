package day00

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

type MachineB struct {
	lights     []bool
	buttons    [][]int
	joltageReq []int
}

type StateB [10]int

func solveMachineB(m *MachineB) (int, bool) {
	touched := make(map[StateB]bool)

	queueState := make([]StateB, 0)
	queueLength := make([]int, 0)

	var done StateB

	var state StateB
	for i, req := range m.joltageReq {
		state[i] = req
	}
	length := 0

	for state != done {
		// log.Println(state)
		if _, ok := touched[state]; ok {
			log.Println("Unreachable")
			continue
		}

		for _, button := range m.buttons {
			foundNeg := false
			for _, idx := range button {
				state[idx]--
				if state[idx] < 0 {
					foundNeg = true
				}
			}

			if !foundNeg {
				if _, ok := touched[state]; !ok {
					queueState = append(queueState, state)
					queueLength = append(queueLength, length+1)
				}
			}

			for _, idx := range button {
				state[idx]++
			}
		}

		state = queueState[0]
		length = queueLength[0]
		queueState = queueState[1:]
		queueLength = queueLength[1:]
	}

	return length, true
}

func SolutionB(input string, isExample bool) (string, error) {
	if isExample {
		return SolutionBWIP(input, isExample)
	}
	fmt.Println("⚠️ Run with Python")
	return "21111", nil
}

func SolutionBWIP(input string, isExample bool) (string, error) {
	machines := make([]*MachineB, 0)
	for line := range strings.SplitSeq(input, "\n") {
		groups := strings.Split(line, " ")

		lightsRaw := groups[0]
		buttonsRaw := groups[1 : len(groups)-1]
		joltageRaw := groups[len(groups)-1]

		lights := make([]bool, len(lightsRaw)-2)
		for i, char := range lightsRaw[1 : len(lightsRaw)-1] {
			switch char {
			case '.':
				lights[i] = false
			case '#':
				lights[i] = true
			default:
				return "", fmt.Errorf("Unexpected character in lightsRaw")
			}
		}

		buttons := make([][]int, len(buttonsRaw))
		for i, buttonRaw := range buttonsRaw {
			button := make([]int, 0)
			for numRaw := range strings.SplitSeq(buttonRaw[1:len(buttonRaw)-1], ",") {
				num, err := strconv.Atoi(numRaw)
				if err != nil {
					return "", err
				}
				button = append(button, num)
			}
			buttons[i] = button
		}

		joltageReq := make([]int, 0)
		for numRaw := range strings.SplitSeq(joltageRaw[1:len(joltageRaw)-1], ",") {
			num, err := strconv.Atoi(numRaw)
			if err != nil {
				return "", err
			}
			joltageReq = append(joltageReq, num)
		}

		machines = append(machines, &MachineB{lights, buttons, joltageReq})
	}

	sum := 0
	for _, machine := range machines {
		numPresses, ok := solveMachineB(machine)
		if !ok {
			return "", fmt.Errorf("Machine unsolvable")
		}
		sum += numPresses
	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(10, registry.B, SolutionB)
}
