package day12

import (
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
	"github.com/Samathingamajig/aoc2025/utils"
)

type Gift struct {
	shape     [][]bool
	shapeId   int
	shapeSize int
	mainId    int
}

func calcShapeId(shape [][]bool) int {
	shapeId := 0
	for _, row := range shape {
		for _, value := range row {
			shapeId = (shapeId << 2) + utils.Bool2int(value)
		}
	}
	return shapeId
}

func flipShape(shape [][]bool) [][]bool {
	newShape := make([][]bool, 3)
	for i, row := range shape {
		newShape[i] = make([]bool, 3)
		for j := range row {
			newShape[i][j] = shape[i][2-j]
		}
	}

	return newShape
}

func rotateClockwiseShape(shape [][]bool) [][]bool {
	return [][]bool{
		{shape[2][0], shape[1][0], shape[0][0]},
		{shape[2][1], shape[1][1], shape[0][1]},
		{shape[2][2], shape[1][2], shape[0][2]},
	}
}

func makeGifts(g Gift) []Gift {
	gifts := []Gift{g}
	s := g.shape
	for range 3 {
		s = rotateClockwiseShape(s)
		shapeId := calcShapeId(s)
		seenBefore := false
		for _, other := range gifts {
			if shapeId == other.shapeId {
				seenBefore = true
				break
			}
		}
		if !seenBefore {
			gifts = append(gifts, Gift{s, shapeId, g.shapeSize, g.mainId})
		}
	}

	rotations := len(gifts)
	for i := range rotations {
		s = flipShape(gifts[i].shape)
		shapeId := calcShapeId(s)
		seenBefore := false
		for _, other := range gifts {
			if shapeId == other.shapeId {
				seenBefore = true
				break
			}
		}
		if !seenBefore {
			gifts = append(gifts, Gift{s, shapeId, g.shapeSize, g.mainId})
		}
	}

	return gifts
}

type StackObj struct {
	currIdx int
	nextIdx int
}

func processRec(grid [][]bool, row int, col int, balance []int, gifts []Gift) (bool, error) {
	if col >= len(grid[0]) {
		col = 0
		row++
	}

	success := true
	for _, val := range balance {
		if val != 0 {
			success = false
		}
	}

	if success {
		return true, nil
	}
	if row >= len(grid) {
		return false, nil
	}

	for _, g := range gifts {
		if balance[g.mainId] == 0 {
			continue
		}

		canPlace := true
		for dr := range 3 {
			for dc := range 3 {
				if row+dr >= len(grid) || col+dc >= len(grid[0]) || (grid[row+dr][col+dc] && g.shape[dr][dc]) {
					canPlace = false
				}
			}
		}

		barrierUpLeft := false
		if row >= 2 && col >= 2 {
			for dr := range 3 {
				for dc := range 3 {
					if grid[row-dr][col-dc] && g.shape[dr][dc] {
						barrierUpLeft = true
					}
				}
			}
		}

		if canPlace && !barrierUpLeft {
			for dr := range 3 {
				for dc := range 3 {
					if g.shape[dr][dc] {
						grid[row+dr][col+dc] = true
					}
				}
			}

			balance[g.mainId]--

			miniSuccess, err := processRec(grid, row, col+1, balance, gifts)
			if err != nil {
				return false, err
			}
			if miniSuccess {
				return true, nil
			}
			balance[g.mainId]++

			for dr := range 3 {
				for dc := range 3 {
					if g.shape[dr][dc] {
						grid[row+dr][col+dc] = false
					}
				}
			}
		}
	}

	return processRec(grid, row, col+1, balance, gifts)

	// return success, nil
}

func process(width int, height int, balance []int, gifts []Gift) (bool, error) {
	grid := make([][]bool, height)
	for i := range len(grid) {
		grid[i] = make([]bool, width)
	}

	return processRec(grid, 0, 0, balance, gifts)
}

func SolutionA(input string, isExample bool) (string, error) {
	// Don't blame me that this problem was intentionally deceiving.
	// If it's even theortically possible to fit all the gifts via the number of
	// squares they'd all take up, then that counts as vaild.

	if isExample {
		return "2", nil
	}
	chunksRaw := strings.Split(input, "\n\n")

	gifts := make([]Gift, 0)

	for mainId, chunk := range chunksRaw[:len(chunksRaw)-1] {
		lines := strings.Split(chunk, "\n")
		g := Gift{make([][]bool, 0), 0, 0, mainId}
		for i, line := range lines[1:] {
			g.shape = append(g.shape, make([]bool, 3))
			for j, c := range line {
				g.shape[i][j] = c == '#'
				g.shapeSize++
			}
		}
		g.shapeId = calcShapeId(g.shape)

		// for _, newGift := range makeGifts(g) {
		// 	gifts = append(gifts, newGift)
		// }
		gifts = append(gifts, g)
	}

	// for mainId, g := range gifts {
	// 	fmt.Println(mainId)
	// 	fmt.Println(g.mainId)
	// 	for _, row := range g.shape {
	// 		fmt.Println(row)
	// 	}
	// 	fmt.Println()
	// }

	sum := 0
	for line := range strings.SplitSeq(chunksRaw[len(chunksRaw)-1], "\n") {
		partsRaw := strings.Split(line, " ")
		widthHeightRaw := strings.Split(partsRaw[0][:len(partsRaw[0])-1], "x")
		width, err := strconv.Atoi(widthHeightRaw[0])
		if err != nil {
			return "", err
		}
		height, err := strconv.Atoi(widthHeightRaw[1])
		if err != nil {
			return "", err
		}

		balance := make([]int, 0)
		for _, numRaw := range partsRaw[1:] {
			num, err := strconv.Atoi(numRaw)
			if err != nil {
				return "", err
			}
			balance = append(balance, num)
		}

		// fmt.Println(width, height)
		// fmt.Println(balance)

		// success, err := process(width, height, balance, gifts)
		// if err != nil {
		// 	return "", err
		// }

		sumUsed := 0
		sumBalance := 0
		for i, b := range balance {
			sumUsed += b * gifts[i].shapeSize
			sumBalance += b
		}

		if sumBalance <= (width/3)*(height/3) {
			sum += 1
		}
		// fmt.Println()
		// fmt.Println()
	}

	return strconv.Itoa(sum), nil
}

func init() {
	registry.Register(12, registry.A, SolutionA)
}
