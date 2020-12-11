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

	return matrix
}

func isSeat(row int, col int, seatMap [][]string) bool {
	if row < 0 || row > len(seatMap) || col < 0 || col > len(seatMap[row]) {
		return false
	}

	return seatMap[row][col] != "."
}

func isSeatOccupied(row int, col int, seatMap [][]string) bool {
	return seatMap[row][col] == "#"
}

func countVisibleOccupiedSeatsAround(row int, col int, seatMap [][]string) int {
	res := 0

TopLeft:
	for shift := 1; row-shift >= 0 && col-shift >= 0; shift++ {
		r := row - shift
		s := col - shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break TopLeft
		}
	}
Top:
	for shift := 1; row-shift >= 0; shift++ {
		r := row - shift
		s := col

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break Top
		}
	}
TopRight:
	for shift := 1; row-shift >= 0 && col+shift < len(seatMap[row-shift]); shift++ {
		r := row - shift
		s := col + shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break TopRight
		}
	}
Left:
	for shift := 1; col-shift >= 0; shift++ {
		r := row
		s := col - shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break Left
		}
	}
Right:
	for shift := 1; col+shift < len(seatMap[row]); shift++ {
		r := row
		s := col + shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break Right
		}
	}
BottomLeft:
	for shift := 1; row+shift < len(seatMap) && col-shift >= 0; shift++ {
		r := row + shift
		s := col - shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break BottomLeft
		}
	}
Bottom:
	for shift := 1; row+shift < len(seatMap); shift++ {
		r := row + shift
		s := col

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break Bottom
		}
	}
BottomRight:
	for shift := 1; row+shift < len(seatMap) && col+shift < len(seatMap[row+shift]); shift++ {
		r := row + shift
		s := col + shift

		if isSeat(r, s, seatMap) {
			if isSeatOccupied(r, s, seatMap) {
				res++
			}
			break BottomRight
		}
	}

	return res
}

// count number of occupied adjacent positions
func countOccupiedArround(i int, j int, matrix [][]string) int {
	occupied := 0

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

func solve(matrix [][]string, around bool) [][]string {
	duplicate := make([][]string, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]string, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	counter := countOccupiedArround
	occupiedThreshold := 3
	if (around == false) {
		counter = countVisibleOccupiedSeatsAround
		occupiedThreshold = 4
	}

	for i:= 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			char := matrix[i][j]
			if matrix[i][j] == "L" && counter(i, j, matrix) == 0 {
				char = "#"
			} else if matrix[i][j] == "#" && counter(i, j, matrix) > occupiedThreshold {
				char = "L"
			}

			duplicate[i][j] = char
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if (matrix[i][j] != duplicate[i][j]) {
				return solve(duplicate, around)
			}
		}
	}

 	return duplicate
}

func main() {
	matrix := getMatrix(98)

	occupied := solve(matrix, true)
	count := 0
	for i := 0; i < len(occupied); i++ {
		for j := 0; j < len(occupied[i]); j++ {
			if (occupied[i][j] == "#") {
				count++
			}
		}
	}


	fmt.Println("P1:", count)

	occupied = solve(matrix, false)
	count = 0

	for i := 0; i < len(occupied); i++ {
		for j := 0; j < len(occupied[i]); j++ {
			if (occupied[i][j] == "#") {
				count++
			}
		}
	}

	fmt.Println("P2:", count)
}
