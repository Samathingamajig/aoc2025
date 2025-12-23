package day09

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
)

type CellB struct {
	x, y                  float64
	up, down, left, right bool
}

func verifyRow(minX, maxX, y float64, cellsInY map[int][]*CellB) bool {
	parity := false

	prevCorner := 0

	neighbors := cellsInY[int(y)]

	for neighborIdx, neighbor := range neighbors {

		if neighborIdx == 0 {
			if neighbor.x < maxX {
				return false
			}
		} else {
			prevNeighbor := neighbors[neighborIdx-1]
			if prevNeighbor.x-1 != neighbor.x && neighbor.x < maxX {
				if !parity {
					return false
				}
			}
		}

		if neighbor.x < minX {
			break
		}

		// // If cell is already explicitly red/green
		// if neighbor.x == x {
		// 	return true
		// }

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
	}

	return parity
}

func verifyCol(minY, maxY, x float64, cellsInX map[int][]*CellB) bool {
	parity := false

	prevCorner := 0

	neighbors := cellsInX[int(x)]

	for neighborIdx, neighbor := range neighbors {

		if neighborIdx == 0 {
			if neighbor.y < maxY {
				return false
			}
		} else {
			prevNeighbor := neighbors[neighborIdx-1]
			if prevNeighbor.y-1 != neighbor.y && neighbor.y < maxY {
				if !parity {
					return false
				}
			}
		}

		if neighbor.y < minY {
			break
		}

		// // If cell is already explicitly red/green
		// if neighbor.x == x {
		// 	return true
		// }

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
	}

	return parity
}

func verifyRect(c1, c2 CellB, cellsInX, cellsInY map[int][]*CellB) bool {
	// dx := math.Copysign(1, c2.x-c1.x)
	// dy := math.Copysign(1, c2.y-c1.y)
	// if c1.x == c2.x {
	// 	dx = 0
	// }
	// if c1.y == c2.y {
	// 	dy = 0
	// }

	minX := min(c1.x, c2.x)
	maxX := max(c1.x, c2.x)
	minY := min(c1.y, c2.y)
	maxY := max(c1.y, c2.y)

	return verifyRow(minX, maxX, c1.y, cellsInY) &&
		verifyRow(minX, maxX, c2.y, cellsInY) &&
		verifyCol(minY, maxY, c1.x, cellsInX) &&
		verifyCol(minY, maxY, c2.x, cellsInX)
	// for x := c1.x; x != c2.x+dx; x += dx {
	// 	if !verifyRow(x, c1.y, cellsInX, cellsInY) ||
	// 		!verifyRow(x, c2.y, cellsInX, cellsInY) {
	// 		return false
	// 	}
	// }

	// for y := c1.y; y != c2.y+dy; y += dy {
	// 	if !verifyRow(c1.x, y, cellsInX, cellsInY) ||
	// 		!verifyRow(c2.x, y, cellsInX, cellsInY) {
	// 		return false
	// 	}
	// }

	// return true
}

func SolutionB(input string, isExample bool) (string, error) {
	// fmt.Printf("%f %f %f\n", math.Copysign(1, -1), math.Copysign(1, 0), math.Copysign(1, 1))
	cellsInX := make(map[int][]*CellB)
	cellsInY := make(map[int][]*CellB)
	cells := make([]CellB, 0)
	numRedCells := 0
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
			float64(x), float64(y),
			false, false, false, false})
		numRedCells++

		if _, ok := cellsInX[x]; !ok {
			cellsInX[x] = make([]*CellB, 0)
		}
		if _, ok := cellsInY[y]; !ok {
			cellsInY[y] = make([]*CellB, 0)
		}
	}

	// fmt.Printf("Cells len: %d\n", len(cells))
	// fmt.Printf("Red Cells len: %d\n", numRedCells)

	for i := 0; i < numRedCells; i++ {
		curr := &cells[i]
		next := &cells[(i+1)%numRedCells]

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
			curr.up = true
			next.down = true
		} else if dy < 0 {
			dy = -1
			curr.up = true
			next.down = true
		}

		x := curr.x + dx
		y := curr.y + dy

		// log.Printf("curr{%f, %f} next{%f, %f} xy{%f, %f} dxdy{%f, %f}n", curr.x, curr.y, next.x, next.y, x, y, dx, dy)

		for x != next.x || y != next.y {
			cells = append(cells, CellB{x, y, dy != 0, dy != 0, dx != 0, dx != 0})
			if _, ok := cellsInX[int(x)]; !ok {
				cellsInX[int(x)] = make([]*CellB, 0)
			}
			if _, ok := cellsInY[int(y)]; !ok {
				cellsInY[int(y)] = make([]*CellB, 0)
			}
			x += dx
			y += dy
		}
	}

	for i := range cells {
		cell := &cells[i]
		if xcells, ok := cellsInX[int(cell.x)]; ok {
			cellsInX[int(cell.x)] = append(xcells, cell)
		} else {
			return "", fmt.Errorf("cellsInX value doesn't exist")
		}

		if ycells, ok := cellsInY[int(cell.y)]; ok {
			cellsInY[int(cell.y)] = append(ycells, cell)
		} else {
			return "", fmt.Errorf("cellsInY value doesn't exist")
		}
	}

	// fmt.Printf("Cells len: %d\n", len(cells))

	// validXs := make([]int, 0)
	// validYs := make([]int, 0)
	// for x := range cellsInX {
	// 	validXs = append(validXs, x)
	// }
	// for y := range cellsInY {
	// 	validYs = append(validYs, y)
	// }
	// sort.Ints(validXs)
	// sort.Ints(validYs)

	// for _, x := range validXs {
	// 	xs := cellsInX[x]
	// 	log.Printf("Border at x=%d: %d cell(s)\n", x, len(xs))
	// }
	// for _, y := range validYs {
	// 	ys := cellsInY[y]
	// 	log.Printf("Border at y=%d: %d cell(s)\n", y, len(ys))
	// }
	for _, xs := range cellsInX {
		sort.Slice(xs, func(i, j int) bool {
			return xs[i].y > xs[j].y
		})
	}
	for _, ys := range cellsInY {
		sort.Slice(ys, func(i, j int) bool {
			return ys[i].x > ys[j].x
		})
	}

	maxArea := 0
	// maxChangeCount := 0
	for i := 0; i < numRedCells; i++ {
		c1 := cells[i]
		for j := i + 1; j < numRedCells; j++ {
			log.Printf("%d x %d\n", i, j)
			c2 := cells[j]
			area := int(math.Abs(c1.x-c2.x+1) * math.Abs(c1.y-c2.y+1))
			if area > maxArea {
				// maxChangeCount++
				if verifyRect(c1, c2, cellsInX, cellsInY) {
					maxArea = area
				}
			}
		}
	}
	// log.Printf("maxChangeCount = %d\n", maxChangeCount)

	return strconv.Itoa(maxArea), nil
}

func init() {
	registry.Register(9, registry.B, SolutionB)
}
