package day00

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

type MachineA struct {
	lights  []bool
	buttons [][]int
	joltage []int
}

func solveMachineA(m *MachineA) (int, bool) {
	buttonsAsXors := make([]int, len(m.buttons))
	for i, button := range m.buttons {
		value := 0
		for _, idx := range button {
			value |= 1 << idx
		}
		buttonsAsXors[i] = value
	}

	initialState := 0
	for i, val := range m.lights {
		if val {
			initialState |= 1 << i
		}
	}

	queueState := make([]int, 0)
	queueLength := make([]int, 0)

	nextState := initialState
	nextLength := 0
	for nextState != 0 {
		for _, button := range buttonsAsXors {
			queueState = append(queueState, nextState^button)
			queueLength = append(queueLength, nextLength+1)
		}

		nextState = queueState[0]
		queueState = queueState[1:]
		nextLength = queueLength[0]
		queueLength = queueLength[1:]
	}

	return nextLength, true
}

func SolutionA(input string, isExample bool) (string, error) {
	machines := make([]*MachineA, 0)
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

		joltage := make([]int, 0)
		for numRaw := range strings.SplitSeq(joltageRaw[1:len(joltageRaw)-1], ",") {
			num, err := strconv.Atoi(numRaw)
			if err != nil {
				return "", err
			}
			joltage = append(joltage, num)
		}

		machines = append(machines, &MachineA{lights, buttons, joltage})
	}

	sum := 0
	for _, machine := range machines {
		numPresses, ok := solveMachineA(machine)
		if !ok {
			return "", fmt.Errorf("Machine unsolvable")
		}
		sum += numPresses
	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(10, registry.A, SolutionA)
}
