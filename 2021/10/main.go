package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var costs = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var completeCosts = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0
	completionCosts := []int{}

	for _, line := range lines {
		score := calculateScore(line)

		if score > 0 {
			ansP1 += score
		} else {
			completionCosts = append(completionCosts, calculateCompletionCost(line))
		}
	}
	sort.Ints(completionCosts)

	fmt.Println(ansP1)
	fmt.Println(completionCosts[len(completionCosts) / 2])
}


func calculateScore(line string) int {
	opened := []string{}

	for _, c := range line {
		sc := string(c)

		if sc == "(" || sc == "{" || sc == "<" || sc == "[" {
			opened = append(opened, sc)
		} else {
			toclose := opened[len(opened) - 1]
			opened = opened[:len(opened)-1]
			if sc == ")" && toclose != "(" {
				return costs[sc]
			} else if sc == "]" && toclose != "[" {
				return costs[sc]
			} else if sc == ">" && toclose != "<" {
				return costs[sc]
			} else if sc == "}" && toclose != "{" {
				return costs[sc]
			}
		}
	}

	return 0
}

func calculateCompletionCost(line string) int {
	opened := []string{}

	for _, c := range line {
		sc := string(c)

		if sc == "(" || sc == "{" || sc == "<" || sc == "[" {
			opened = append(opened, sc)
		} else {
			opened = opened[:len(opened)-1]
		}
	}

	lineCost := 0
	for i := len(opened) - 1; i >= 0; i-- {
		lineCost = 5 * lineCost + completeCosts[opened[i]]
	}

	return lineCost
}