package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
	// "strings"
)

func solveP1(groupSize int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 1
	groupCodes := make([]int, 0)
	for scanner.Scan() {
		lineText := scanner.Text()

		if (len(lineText) == 0) {
			continue
		}

		val, _ := strconv.Atoi(lineText)

		if (len(groupCodes) > groupSize) {
			groupCodes = groupCodes[1:len(groupCodes)]
		}

		if len(groupCodes) == groupSize && ! isSumOfTwoElements(groupCodes, val) {
			return val
		}

		groupCodes = append(groupCodes, val)

		count++
	}

	return 0
}

func solveP2(desiredSum int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	groupCodes := make([]int, 0)
	sum := 0

	for scanner.Scan() {
		lineText := scanner.Text()

		if (len(lineText) == 0) {
			continue
		}

		val, _ := strconv.Atoi(lineText)

		if (sum < desiredSum) {
			groupCodes = append(groupCodes, val)
			sum += val
			continue
		} else if (sum > desiredSum) {
			for sum > desiredSum && len(groupCodes) > 0 {
				sum -= groupCodes[0]
				groupCodes = groupCodes[1:len(groupCodes)]
			}

			if (sum == desiredSum) {
				sort.Ints(groupCodes)
				return groupCodes[0] + groupCodes[len(groupCodes) -1]
			}

			groupCodes = append(groupCodes, val)
			sum += val
		}
	}

	return 0
}

func isSumOfTwoElements(groupCodes []int, val int) bool {
	arr := make([]int, len(groupCodes))
	copy(arr, groupCodes)

	left := 0
	right := len(arr) - 1
	sort.Ints(arr)

    for (left < right) {
        if(arr[left] + arr[right] == val) {
            return true
		} else if (arr[left] + arr[right] < val) {
			left++
		} else {
			right--
		}
	}

    return false
}

func main() {
	p1 := solveP1(25)
	fmt.Println("P1:", p1)

	p2 := solveP2(p1)
	fmt.Println("P2:", p2)
}
