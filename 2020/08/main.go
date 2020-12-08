package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solve(inputString []string, noloops bool) int {
	acc := 0
	executed := make(map[int]bool, 0)
	index := 0

	for executed[index] != true {
		if (index > len(inputString) - 1) {
			break
		}

		line := strings.Split(inputString[index], " ")

		executed[index] = true
		instruction := line[0]
		value, _ := strconv.Atoi(line[1])

		if (instruction == "jmp") {
			index += value
		} else {
			if (instruction == "acc") {
				acc += value
			}
			index++
		}
	}

	if noloops {
		if executed[len(inputString) - 1] {
			return acc
		}

		return 0
	}

	return acc
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	inputString := strings.Split(strings.TrimSpace(string(input)), "\n")

	acc := solve(inputString, false)
	fmt.Println("P1", acc)

	switchIndex := make(map[int]bool, 0)
	originalInputString := make([]string, len(inputString))
	copy(originalInputString, inputString)
	noloops := false

	for noloops != true {
		switched := false
		copy(inputString, originalInputString)

		for i, v := range(inputString) {
			line := strings.Split(v, " ")
			instruction := line[0]

			if (switched) {
				break
			}

			if ! switchIndex[i] {
				if instruction == "nop" && line[1] != "+0" {
					inputString[i] = "jmp " + line[1]
					switchIndex[i] = true
					switched = true
				}
			}

			if ! switchIndex[i] {
				if instruction == "jmp" {
					inputString[i] = "nop " + line[1]
					switchIndex[i] = true
					switched = true
				}
			}
		}

		acc = solve(inputString, true)

		if (acc != 0) {
			noloops = true
			fmt.Println("P2", acc)
		}
	}
}
