package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func getVoltages() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	voltages := make([]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		if (len(lineText) == 0) {
			continue
		}

		val, _ := strconv.Atoi(lineText)

		voltages = append(voltages, val)
	}

	sort.Ints(voltages)
	voltages = append([]int{0}, voltages...)
	voltages = append(voltages, voltages[len(voltages) - 1] + 3)

	return voltages
}

func solveP1(voltages []int) int {
	diffOneCount := 0
	diffThreeCount := 0
	for i := 1; i < len(voltages); i++ {
		if voltages[i] - voltages[i - 1] == 1 {
			diffOneCount++
		}
		if voltages[i] - voltages[i - 1] == 3 {
			diffThreeCount++
		}
	}

	return diffOneCount * diffThreeCount
}

func solveP2(voltages []int, solutions map[int]int, index int) int {
	if index == len(voltages) - 1 {
		return 1
	}

	if val, ok := solutions[index]; ok {
		return val
	}

	permutations := 0
	for j := index + 1; j < len(voltages); j++ {
		if voltages[j] - voltages[index] <= 3 {
			permutations += solveP2(voltages, solutions, j)
		}
	}

	solutions[index] = permutations

	return permutations
}

func main() {
	voltages := getVoltages()

	p1 := solveP1(voltages)
	fmt.Println("P1:", p1)

	solutionsMap := make(map[int]int, 0)
	p2 := solveP2(voltages, solutionsMap, 0)
	fmt.Println("P2:", p2)
}
