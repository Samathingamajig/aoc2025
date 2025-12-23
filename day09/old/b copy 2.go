package day09

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
	"github.com/Samathingamajig/aoc2025/utils"
)

type CellB struct {
	x, y                  int
	up, down, left, right bool
}

// func verifyRow(minX, maxX, y int, cellsInY map[int][]*CellB) bool {
func verifyRow(minX, maxX, y int, cellsInY map[int][]CellB) bool {
	parity := false

	prevCorner := 0

	neighbors := cellsInY[y]

	if (minX+1 < neighbors[0].x) || (neighbors[len(neighbors)-1].x < maxX-1) {
		minNX := neighbors[0].x
		maxNX := neighbors[len(neighbors)-1].x
		for _, n := range neighbors {
			minNX = min(minNX, n.x)
			maxNX = max(maxNX, n.x)
		}

		if minNX != neighbors[0].x {
			log.Println("minNX differs")
		}
		if maxNX != neighbors[len(neighbors)-1].x {
			log.Println("maxNX differs")
		}

		return false
	}
	return true

	// for neighborIdx, neighbor := range neighbors {
	for _, neighbor := range neighbors {
		if neighbor.x < minX {
			// Only care about updating parity/prevCorner
			if neighbor.up && neighbor.down {
				parity = !parity
			} else if neighbor.left != neighbor.right {
				corner := 1
				if neighbor.left {
					corner = -1
				}

				if prevCorner != 0 {
					if corner == prevCorner {
						prevCorner = 0
						parity = !parity
					} else {
						prevCorner = corner
					}
				} else {
					prevCorner = corner
					parity = !parity
				}
			}
		} else if neighbor.x < maxX {
			if minX < neighbor.x && !parity {
				return false
			}
			// Within
			if minX < neighbor.x && neighbor.up && neighbor.down {
				return false
			}
		} else {
			// Beyond
			// if
		}

	}

	return true
}

// func verifyRect(c1, c2 CellB, cellsInY map[int][]*CellB) bool {
func verifyRect(c1, c2 CellB, cellsInY map[int][]CellB) bool {
	minX := min(c1.x, c2.x)
	maxX := max(c1.x, c2.x)
	minY := min(c1.y, c2.y)
	maxY := max(c1.y, c2.y)

	for y := minY; y <= maxY; y++ {
		if !verifyRow(minX, maxX, y, cellsInY) {
			return false
		}
	}

	return true
}

func SolutionB(input string, isExample bool) (string, error) {
	// fmt.Printf("%f %f %f\n", math.Copysign(1, -1), math.Copysign(1, 0), math.Copysign(1, 1))
	// cellsInX := make(map[int][]*CellB)
	// cellsInY := make(map[int][]*CellB)
	cellsInY := make(map[int][]CellB)
	cells := make([]CellB, 0)
	numRedCells := 0
	dupes := make(map[int]bool)
	for line := range strings.SplitSeq(input, "\n") {
		nums := strings.Split(line, ",")
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(nums[1])
		if err != nil {
			return "", err
		}

		cells = append(cells, CellB{
			x, y,
			false, false, false, false})
		numRedCells++

		if _, ok := dupes[x*1000000+y]; ok {
			log.Println("BAD!!! Red cell duplicated")
		}
		dupes[x*1000000+y] = true

		// if _, ok := cellsInX[x]; !ok {
		// 	cellsInX[x] = make([]*CellB, 0)
		// }
		if _, ok := cellsInY[y]; !ok {
			// cellsInY[y] = make([]*CellB, 0)
			cellsInY[y] = make([]CellB, 0)
		}
	}

	// fmt.Printf("Cells len: %d\n", len(cells))
	// fmt.Printf("Red Cells len: %d\n", numRedCells)

	for i := 0; i < numRedCells; i++ {
		curr := &cells[i]
		next := &cells[(i+1)%numRedCells]
		// log.Printf("%d %d\n", i, (i+1)%numRedCells)

		dx := next.x - curr.x
		dy := next.y - curr.y

		if dx > 0 {
			dx = 1
			curr.right = true
			next.left = true
		} else if dx < 0 {
			dx = -1
			curr.left = true
			next.right = true
		}

		if dy > 0 {
			dy = 1
			// curr.up = true
			// next.down = true
			curr.down = true
			next.up = true
		} else if dy < 0 {
			dy = -1
			// curr.down = true
			// next.up = true
			curr.up = true
			next.down = true
		}

		x := curr.x + dx
		y := curr.y + dy

		if dx != 0 && dy != 0 {
			log.Printf("BAD!!! dx and dy are 0\n")
		}

		// log.Printf("curr{%d, %d} next{%d, %d} xy{%d, %d} dxdy{%d, %d}\n", curr.x, curr.y, next.x, next.y, x, y, dx, dy)

		for x != next.x || y != next.y {
			cells = append(cells, CellB{x, y, dy != 0, dy != 0, dx != 0, dx != 0})
			if _, ok := dupes[x*1000000+y]; ok {
				log.Println("BAD!!! Green cell duplicated")
			}
			dupes[x*1000000+y] = true
			// if _, ok := cellsInX[int(x)]; !ok {
			// 	cellsInX[int(x)] = make([]*CellB, 0)
			// }
			if _, ok := cellsInY[y]; !ok {
				// cellsInY[y] = make([]*CellB, 0)
				cellsInY[y] = make([]CellB, 0)
			}
			x += dx
			y += dy
		}
	}

	for i := 0; i < numRedCells; i++ {
		cell := cells[i]
		if (cell.down && cell.up) || (cell.left && cell.right) {
			log.Printf("BAD!!! Found red tile not as corner: i=%d\n", i)
		}
	}

	for i := range cells {
		cell := &cells[i]
		// if xcells, ok := cellsInX[int(cell.x)]; ok {
		// 	cellsInX[int(cell.x)] = append(xcells, cell)
		// } else {
		// 	return "", fmt.Errorf("cellsInX value doesn't exist")
		// }

		if ycells, ok := cellsInY[cell.y]; ok {
			// if !(cell.left && cell.right) {
			// cellsInY[cell.y] = append(ycells, cell)
			cellsInY[cell.y] = append(ycells, *cell)
			// }
		} else {
			return "", fmt.Errorf("cellsInY value doesn't exist")
		}
	}

	// fmt.Printf("Cells len: %d\n", len(cells))

	// validXs := make([]int, 0)
	validYs := make([]int, 0)
	// for x := range cellsInX {
	// 	validXs = append(validXs, x)
	// }
	for y := range cellsInY {
		validYs = append(validYs, y)
	}
	// sort.Ints(validXs)
	sort.Ints(validYs)

	// for _, x := range validXs {
	// 	xs := cellsInX[x]
	// 	log.Printf("Border at x=%d: %d cell(s)\n", x, len(xs))
	// }
	maxYLength := -1
	numBadPairs := 0
	for _, y := range validYs {
		ys := cellsInY[y]
		// log.Printf("Border at y=%d: %d cell(s)\n", y, len(ys))
		maxYLength = max(maxYLength, len(ys))
		for i := range ys {
			if i > 0 && ys[i-1].x+1 == ys[i].x {
				numBadPairs++
			}
		}
	}
	log.Printf("maxYLength=%d\n", maxYLength)
	log.Printf("numBadPairs=%d\n", numBadPairs)
	// for _, xs := range cellsInX {
	// 	sort.Slice(xs, func(i, j int) bool {
	// 		return xs[i].y > xs[j].y
	// 	})
	// }
	stupidPrinted := false
	for _, ys := range cellsInY {
		sort.Slice(ys, func(i, j int) bool {
			return ys[i].x < ys[j].x
		})
		if len(ys) == 4 && !stupidPrinted {
			log.Printf("ys: [%v, %v, %v, %v]\n", ys[0].x, ys[1].x, ys[2].x, ys[3].x)
			stupidPrinted = true
		}
	}

	printBoard(cells, 14, 9)

	return "-1", nil

	// for _, y := range validYs {
	// 	ys := cellsInY[y]
	// 	log.Printf("%d: %+v\n", y, ys)
	// }

	maxArea := -1
	// if !isExample {
	// 	maxArea = 1550726012
	// }
	maxChangeCount := 0
	for i := 0; i < numRedCells; i++ {
		c1 := cells[i]
		for j := i + 1; j < numRedCells; j++ {
			// log.Printf("%d x %d\n", i, j)
			c2 := cells[j]
			area := utils.AbsInt(c1.x-c2.x+1) * utils.AbsInt(c1.y-c2.y+1)
			if area > maxArea {
				maxChangeCount++
				if verifyRect(c1, c2, cellsInY) {
					maxArea = area
				}
			}
		}
	}
	log.Printf("maxChangeCount = %d\n", maxChangeCount)

	return strconv.Itoa(maxArea), nil
}

func printBoard(cells []CellB, width int, height int) {
	grid := make([][]rune, height)
	for i := range height {
		grid[i] = make([]rune, width)
		for j := range width {
			grid[i][j] = '.'
		}
	}

	for _, cell := range cells {
		char := '.'
		switch dirAsInt(cell.up, cell.down, cell.left, cell.right) {
		case UP_2 + DOWN_2:
			char = '║'
		case UP_2 + LEFT_2:
			char = '╝'
		case UP_2 + RIGHT_2:
			char = '╚'
		case LEFT_2 + RIGHT_2:
			char = '═'
		case DOWN_2 + LEFT_2:
			char = '╗'
		case DOWN_2 + RIGHT_2:
			char = '╔'
		}
		grid[cell.y][cell.x] = char
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func printBoardFromMaps(cellsInY map[int][]CellB, width int, height int) {
	grid := make([][]rune, height)
	for i := range height {
		grid[i] = make([]rune, width)
		for j := range width {
			grid[i][j] = '.'
		}
	}

	for _, cells := range cellsInY {
		for _, cell := range cells {
			char := '.'
			switch dirAsInt(cell.up, cell.down, cell.left, cell.right) {
			case UP_2 + DOWN_2:
				char = '║'
			case UP_2 + LEFT_2:
				char = '╝'
			case UP_2 + RIGHT_2:
				char = '╚'
			case LEFT_2 + RIGHT_2:
				char = '═'
			case DOWN_2 + LEFT_2:
				char = '╗'
			case DOWN_2 + RIGHT_2:
				char = '╔'
			}
			grid[cell.y][cell.x] = char
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

const (
	UP_2 = 1 << iota
	DOWN_2
	LEFT_2
	RIGHT_2
)

func dirAsInt(up, down, left, right bool) int {
	result := 0
	if up {
		result += UP_2
	}
	if down {
		result += DOWN_2
	}
	if left {
		result += LEFT_2
	}
	if right {
		result += RIGHT_2
	}
	return result
}

func init() {
	registry.Register(9, registry.B, SolutionB)
}
