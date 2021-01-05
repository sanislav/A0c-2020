package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	ON = 1
	OFF = 0
)

func countOnNeighbours(grid [102][102]int, i int, j int) int {
	return grid[i - 1][j - 1] + grid[i - 1][j] + grid[i - 1][j + 1] + grid[i][j - 1] + grid[i][j + 1] + grid[i + 1][j - 1] + grid[i + 1][j] + grid[i + 1][j + 1];
}

func makeGeneration(grid [102][102]int, isP2 bool) [102][102]int {
	prevGrid := [102][102]int{}

	if isP2 {
		grid[1][1] = ON
		grid[1][100] = ON
		grid[100][1] = ON
		grid[100][100] = ON
	}

	for i := range grid {
		for j, l := range grid[i] {
			prevGrid[i][j] = l
		}
	}

	// A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
	// A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			c := prevGrid[i][j]

			on := countOnNeighbours(prevGrid, i, j)
			if c == ON && on != 2 && on != 3 {
				grid[i][j] = OFF
			}
			if c == OFF && on == 3 {
				grid[i][j] = ON
			}
		}
	}

	return grid
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0
	grid := [102][102]int{}
	gridP2 := [102][102]int{}

	for i, s := range lines {
		for j, c := range s {
			n := OFF
			if (string(c) == "#") {
				n = ON
			}

			grid[i+1][j+1] = n
			gridP2[i+1][j+1] = n
		}
	}

	for c := 0; c < 100; c++ {
		grid = makeGeneration(grid, false)
	}

	for c := 0; c < 100; c++ {
		gridP2 = makeGeneration(gridP2, true)
	}

	for i := range grid {
		for _, c := range grid[i] {
			if c == ON {
				ansP1++
			}
		}
	}

	gridP2[1][1] = ON
	gridP2[1][100] = ON
	gridP2[100][1] = ON
	gridP2[100][100] = ON

	for i := range gridP2 {
		for _, c := range gridP2[i] {
			if c == ON {
				ansP2++
			}
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
