package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	numbers := strings.Split(lines[0], ",")
	boards := getBoards(lines)
	winningBoards := [][][]map[string]string{}
	firstWinNum := -1
	lastWinNum := -1

	for _, num := range(numbers) {
		numInt, _ := strconv.Atoi(num)
		boards = checkBoards(num, boards)
		newBoards := [][][]map[string]string{}

		for _, b := range(boards) {
			if (isWinner(b)) {
				if firstWinNum == -1 {
					firstWinNum = numInt
				}
				lastWinNum = numInt
				winningBoards = append(winningBoards, b)
			} else {
				newBoards = append(newBoards, b)
			}
		}
		boards = newBoards

		if (len(boards) == 0) {
			break
		}
	}

	p1 := calculate(winningBoards[0], firstWinNum)
	p2 := calculate(winningBoards[len(winningBoards) - 1], lastWinNum)

	fmt.Println(p1)
	fmt.Println(p2)
}

func getBoards(lines []string) [][][]map[string]string {
	board := [][]map[string]string{}
	boards := [][][]map[string]string{}

	for i := 2; i < len(lines); i++ {
		if (len(lines[i]) == 0) {
			boards = append(boards, board)
			board = [][]map[string]string{}
		}

		num := strings.Split(lines[i], " ")

		ma := []map[string]string{}
		for _, n := range(num) {
			if len(n) == 0 {
				continue
			}

			m := map[string]string{}
			m["num"] = n
			m["checked"] = "false"
			ma = append(ma, m)
		}

		if (len(ma) > 0) {
			board = append(board, ma)
		}

		if (i == len(lines) - 1) {
			boards = append(boards, board)
		}
	}

	return boards
}

func checkBoards(num string, boards [][][]map[string]string) [][][]map[string]string {
	for bi, b := range(boards) {
		for m := 0; m < len(b); m++ {
			for ind, item := range(b[m]) {
				if (item["num"] == num) {
					item["checked"] = "true"
				}
				b[m][ind] = item
			}
		}

		boards[bi] = b
	}

	return boards
}

func isWinner(board [][]map[string]string) bool {
	colWin := map[int]bool{}

	lineWin := false
	for m := 0; m < len(board); m++ {
		lineWin = true
		for n := 0; n < len(board[m]); n++ {
			if _, ok := colWin[n]; !ok {
				colWin[n] = true
			}

			if board[m][n]["checked"] == "false" {
				lineWin = false
				colWin[n] = false
			}
		}

		if lineWin {
			return lineWin
		}
	}
	if lineWin == true {
		return true
	}

	for _, b := range colWin {
		if b != false {
			return true
		}
	}

	return false
}

func calculate(board [][]map[string]string, num int) int {
	sum := 0
	for i := range(board) {
		for j := range(board) {
			if (board[i][j]["checked"] == "false") {
				intV, _ := strconv.Atoi(board[i][j]["num"])
				sum += intV
			}
		}
	}
	return sum * num
}