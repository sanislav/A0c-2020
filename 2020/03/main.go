package main

import (
	"fmt"
	"os"
	"bufio"
)

func solve(right int, down int) int {
	col := 0
	row := 0
	treesEncountered := 0
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineText := scanner.Text()
		if (row % down != 0) {
			row++
			continue
		}
		if (col >= len(lineText)) {
			col = col - len(lineText)
		}
		currentChar := string(lineText[col])
		if (currentChar == "#") {
			treesEncountered++
		}
		col += right
		row++
	}

	return treesEncountered
}

func main() {
	fmt.Println("Trees encountered for right 1 down 1:", solve(1, 1))
	fmt.Println("Trees encountered for right 3 down 1:", solve(3, 1))
	fmt.Println("Trees encountered for right 5 down 1:", solve(5, 1))
	fmt.Println("Trees encountered for right 7 down 1:", solve(7, 1))
	fmt.Println("Trees encountered for right 1 down 2:", solve(1, 2))
}