package main

import (
	"fmt"
	"os"
	"bufio"
	// "sort"
)

func solve() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	highestID := 0
	var seats []int

	for scanner.Scan() {
		lineText := scanner.Text()
		if (len(lineText) == 0) {
			continue
		}

		rowID := decode(lineText[0:6], 0, 127, "F", "B")
		colID := decode(lineText[7:10], 0, 7, "L", "R")
		seatID := (rowID * 8) + colID
		seats = append(seats, seatID)
		if seatID > highestID {
			highestID = seatID
		}
	}

	return highestID
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func decode(s string, lowerCount int, upperCount int, lowerIndicator string, upperIndicator string) int {
	for c := 0; c < len(s); c++ {
		if string(s[c]) == lowerIndicator {
			upperCount = upperCount - 1 - (upperCount - lowerCount) / 2
		}

		if string(s[c]) == upperIndicator {
			lowerCount = lowerCount + 1 + (upperCount - lowerCount) / 2
		}
	}

	if (string(s[len(s) - 1]) != lowerIndicator) {
		return lowerCount
	}

	return upperCount
}


func main() {
	id := solve()
	fmt.Println("Highest seat ID:", id)
}