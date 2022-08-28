package main

import "fmt"

func main() {

	var world [][]int = initWorld(10, 10)

	world[1][6] = 1
	//fmt.Println(world)
	//fmt.Println(world[5][6])
	displayWorld(world)

	fmt.Println("Test countNeighbor Classic for (2,6) :", countNeighbor(world, 2, 6), " neighbors  (expected 1)")
	fmt.Println("Test countNeighbor Exception Border for (0,6) :", countNeighbor(world, 0, 6), " neighbors  (expected 1)")
	fmt.Println("Test countNeighbor Exception Corner for (0,0) :", countNeighbor(world, 0, 0), " neighbors  (expected 0)")
}

func initWorld(width int, length int) [][]int {
	var worldGrid [][]int = make([][]int, int(length))
	for i := range worldGrid {
		worldGrid[i] = make([]int, int(width))
	}
	return worldGrid
}

func displayWorld(world [][]int) {
	for i := range world {
		for j := range world[0] {
			fmt.Print(" ", world[i][j], " ")
		}
		fmt.Println()
	}
}

func countNeighbor(world [][]int, line int, col int) int {
	var numOfNeighbor int = 0
	var minRangeCol int = -1
	var maxRangeCol int = 1
	var minRangeLine int = -1
	var maxRangeLine int = 1

	switch line {
	case 0:
		minRangeLine = 0
	case len(world[0]) - 1:
		maxRangeLine = 0
	}

	switch col {
	case 0:
		minRangeCol = 0
	case len(world) - 1:
		maxRangeCol = 0
	}

	numOfNeighbor = searchNeighbor(world, line, col, minRangeCol, maxRangeCol, minRangeLine, maxRangeLine, numOfNeighbor)

	if world[line][col] == 1 && numOfNeighbor != 0 {
		numOfNeighbor -= 1
	}
	return numOfNeighbor
}

func searchNeighbor(world [][]int, line int, col int, minRangeCol int, maxRangeCol int, minRangeLine int, maxRangeLine int, numOfNeighbor int) int {
	for i := minRangeLine; i <= maxRangeLine; i++ {
		for j := minRangeCol; j <= maxRangeCol; j++ {
			if world[line+i][col+j] == 1 {
				numOfNeighbor += 1
			}
		}
	}
	return numOfNeighbor
}
