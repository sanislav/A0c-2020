package main

import (
	"fmt"
	"bufio"
	"os"
)

func getMatrix(dim int) [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]string, dim)

	line := 0
	for scanner.Scan() {
		lineText := scanner.Text()

		if (len(lineText) == 0) {
			break
		}

		matrix[line] = make([]string, 0, dim)
		vector := make([]string, dim)

		for col, c := range(lineText) {
			vector[col] = string(c)
			matrix[line] = append(matrix[line], vector[col])
		}

		line++
	}
	fmt.Println(matrix)
	return matrix
}

// count number of occupied adjacent positions
func seeOccupied(matrix [][]string, i int, j int, seeAdjacent int) int {
	occupied := 0

	if (seeAdjacent == -1) {
		dr := []int{-1, 0, 1}
		dc := []int{-1, 0, 1}

		for _, ddr := range dr {
			r := i + ddr
			if (r < 0 || r >= len(matrix) ) {
				continue
			}
			for _, ddc := range dc {
				c := j + ddc

				if (c < 0 || c >= len(matrix[i])) {
					continue
				}

				if matrix[r][c] == "#" {
					occupied++
					break
				}
			}
		}

		return occupied
	}

	if i > 0 {
		if j > 0 && matrix[i-1][j-1] == "#" {
			occupied++
		}
		if matrix[i-1][j] == "#" {
			occupied++
		}

		if j < len(matrix[i]) - 1 && matrix[i-1][j+1] == "#" {
			occupied++
		}
	}

	if j > 0 && matrix[i][j-1] == "#" {
		occupied++
	}
	if j < len(matrix[i]) - 1 && matrix[i][j+1] == "#" {
		occupied++
	}

	if i < len(matrix) - 1 {
		if j > 0 {
			if matrix[i+1][j-1] == "#" {
				occupied++
			}
		}

		if matrix[i+1][j] == "#" {
			occupied++
		}

		if j < len(matrix[i]) - 1 && matrix[i+1][j+1] == "#" {
			occupied++
		}
	}

	return occupied
}

func solve(matrix [][]string, seeAdjacent int, occupiedThreshold int) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	for i:= 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			char := matrix[i][j]
			if matrix[i][j] == "L" && seeOccupied(matrix, i, j, seeAdjacent) == 0 {
				char = "#"
			} else if matrix[i][j] == "#" && seeOccupied(matrix, i, j, seeAdjacent) > occupiedThreshold {
				char = "L"
			}

			duplicate[i][j] = char
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if (matrix[i][j] != duplicate[i][j]) {
				return solve(duplicate, seeAdjacent, occupiedThreshold)
			}
		}
	}

 	return duplicate
}

func main() {
	matrix := getMatrix(98)

	occupied := solve(matrix, 1, 3)
	count := 0
	for i := 0; i < len(occupied); i++ {
		for j := 0; j < len(occupied[i]); j++ {
			if (occupied[i][j] == "#") {
				count++
			}
		}
	}


	fmt.Println("P1:", count)

	occupied = solve(matrix, -1, 4)
	count = 0
	for i := 0; i < len(occupied); i++ {
		for j := 0; j < len(occupied[i]); j++ {
			if (occupied[i][j] == "#") {
				count++
			}
		}
	}

	fmt.Println("P1:", count)
}
