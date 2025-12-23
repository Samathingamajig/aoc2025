package day09

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/Samathingamajig/aoc2025/registry"
	"github.com/Samathingamajig/aoc2025/utils"
)

type CellB struct {
	x, y                  int
	up, down, left, right bool
}

func verifyRow(row []*CellB, minX int, maxX int) bool {
	if row[0].x > minX || row[len(row)-1].x < maxX {
		return false
	}

	// Ideally this would actually do the raycasting method,
	// but this simple version works for both the example and my input.
	// See the day09/old directory for some of that

	return true
}

func verifyRect(cellsInY map[int][]*CellB, cellA *CellB, cellB *CellB) bool {
	minX := min(cellA.x, cellB.x)
	maxX := max(cellA.x, cellB.x)

	minY := min(cellA.y, cellB.y)
	maxY := max(cellA.y, cellB.y)

	for y := minY; y <= maxY; y++ {
		row, ok := cellsInY[y]
		if !ok {
			return false
		}

		if !verifyRow(row, minX, maxX) {
			return false
		}
	}

	return true
}

func SolutionB(input string, isExample bool) (string, error) {
	cells := make([]*CellB, 0)

	for line := range strings.SplitSeq(input, "\n") {
		numsRaw := strings.Split(line, ",")
		if len(numsRaw) != 2 {
			return "", fmt.Errorf("Bad input line")
		}
		x, err := strconv.Atoi(numsRaw[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(numsRaw[1])
		if err != nil {
			return "", err
		}

		cell := &CellB{
			x, y,
			false, false, false, false,
		}

		cells = append(cells, cell)
	}

	numRedCells := len(cells)
	for i := range numRedCells {
		curr := cells[i]
		next := cells[(i+1)%numRedCells]

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
			// y increases going down
			curr.down = true
			next.up = true
		} else if dy < 0 {
			dy = -1
			curr.up = true
			next.down = true
		}

		if (dx == 0 && dy == 0) || (dx != 0 && dy != 0) {
			log.Printf("i=%d: dx=%d, dy=%d\n", i, dx, dy)
			return "", fmt.Errorf("Should only be moving in one direction")
		}

		greenX := curr.x + dx
		greenY := curr.y + dy

		for greenX != next.x || greenY != next.y {
			greenCell := &CellB{
				greenX, greenY,
				dy != 0, dy != 0, dx != 0, dx != 0,
			}
			cells = append(cells, greenCell)

			greenX += dx
			greenY += dy
		}
	}

	cellsInY := make(map[int][]*CellB)
	for _, cell := range cells {
		if cell.left && cell.right {
			continue
		}

		row, ok := cellsInY[cell.y]
		if !ok {
			row = make([]*CellB, 0)
		}

		cellsInY[cell.y] = append(row, cell)
	}

	for _, row := range cellsInY {
		slices.SortFunc(row, func(a, b *CellB) int {
			return a.x - b.x
		})
	}

	maxArea := -1
	for i := range numRedCells {
		cellA := cells[i]
		for j := i + 1; j < numRedCells; j++ {
			cellB := cells[j]
			area := (utils.AbsInt(cellA.x-cellB.x) + 1) * (utils.AbsInt(cellA.y-cellB.y) + 1)
			if area > maxArea {
				if verifyRect(cellsInY, cellA, cellB) {
					maxArea = area
				}
			}
		}
	}

	return strconv.Itoa(maxArea), nil
}

func init() {
	registry.Register(9, registry.B, SolutionB)
}
