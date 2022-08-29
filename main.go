package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Enter the size of the world: ")
	var size int
	fmt.Scanln(&size)

	fmt.Println("Enter the number of iterations: ")
	var iterationNumber int
	fmt.Scanln(&iterationNumber)

	var world [][]int = initWorld(size, size)

	generateRandomWorld(world, 0.35)
	displayWorld(world)

	runSimulation(world, iterationNumber)

	//world[1][6] = 1
	//fmt.Println(world)
	//fmt.Println(world[5][6])
	//fmt.Println("Neighbor Matrix : ")
	//displayWorld(generateNeighborMatrix(world))

	//testBed(world)

}

func testBed(world [][]int) {
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
			if world[i][j] == 1 {
				fmt.Print(" ", "■", " ")
			} else {
				fmt.Print(" ", " ", " ")
			}

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

func generateNeighborMatrix(world [][]int) [][]int {
	var neighborMatrix [][]int = initWorld(len(world), len(world[0]))

	for i := range world {
		for j := range world[0] {
			neighborMatrix[i][j] = countNeighbor(world, i, j)
		}
	}
	return neighborMatrix
}

func generateRandomWorld(world [][]int, frequency float64) [][]int {
	for i := range world {
		for j := range world[0] {
			if rand.Float64() < frequency {
				world[i][j] = 1
			}
		}
	}
	return world
}

func runSimulation(world [][]int, iteration int) {
	var currentWorld [][]int = world
	var currentNeighbor [][]int = generateNeighborMatrix(world)

	for i := 0; i < iteration; i++ {
		currentWorld = runIteration(currentWorld, currentNeighbor)
		currentNeighbor = generateNeighborMatrix(currentWorld)
		fmt.Println("Current World iteration", i)
		displayWorld(currentWorld)
		fmt.Println("")
		fmt.Println("")
	}
}

func runIteration(currentWorld [][]int, currentNeighbor [][]int) [][]int {
	var newWorld [][]int = currentWorld
	for i := range currentWorld {
		for j := range currentWorld[0] {
			//if the current cell is alive
			if currentWorld[i][j] == 1 {
				if currentNeighbor[i][j] < 2 || currentNeighbor[i][j] > 3 {
					newWorld[i][j] = 0
				}
			}
			//if the current cell is dead
			if currentWorld[i][j] == 0 {
				if currentNeighbor[i][j] == 3 {
					newWorld[i][j] = 1
				}
			}
		}
	}
	return newWorld
}
