package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

var ansP1 int
var combinations = make(map[int]int)
var min int

func subsetSum(numbers []int, target int, partial []int) {
	s := 0

	for _, i := range partial {
		s += i
	}

	if s == target {
		ansP1++
		combinations[len(partial)]++

		if len(partial) < min || min == 0 {
			min = len(partial)
		}
		return
	} else if s > target {
        return
	}

	for i:= 0; i < len(numbers); i++ {
        n := numbers[i]
        remaining := numbers[i+1:]
        subsetSum(remaining, target, append(partial, n))
	}

	return
}


func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	nums := []int{}

	for _, s := range lines {
		i, _ := strconv.Atoi(s)
		nums = append(nums, i)
	}

	subsetSum(nums, 150, []int{})
	fmt.Println(ansP1)
	ansP2 := combinations[min]
	fmt.Println(ansP2)
}
