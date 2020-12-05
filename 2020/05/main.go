package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
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
	seats := make([]int, 0)

	for scanner.Scan() {
		lineText := scanner.Text()
		if (len(lineText) == 0) {
			continue
		}
		l := lineText[0:7]
		rowStr := strings.ReplaceAll(l, "F", "0")
		rowStr = strings.ReplaceAll(rowStr, "B", "1")
		rowID, _ := strconv.ParseInt(rowStr, 2, 64)

		ll := lineText[7:10]
		colStr := strings.ReplaceAll(ll, "L", "0")
		colStr = strings.ReplaceAll(colStr, "R", "1")
		colID, _ := strconv.ParseInt(colStr, 2, 16)


		seatID := (int(rowID) * 8) + int(colID)

		seats = append(seats, seatID)

		if seatID > highestID {
			highestID = seatID
		}
	}

	sort.Ints(seats)
    prevSeat := 0
	for i := 0; i < len(seats); i++ {
		if (seats[i] - prevSeat) == 2 {
			fmt.Println("My seat:", seats[i] - 1)
		}

		prevSeat = seats[i]
	}

	return highestID
}


func main() {
	id := solve()
	fmt.Println("Highest seat ID:", id)
}