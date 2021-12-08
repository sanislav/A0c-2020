package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Matrix struct {
	data [1000][1000]int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	pos := strings.Split(lines[0], ",")
	positions := []int{}

	for _, p := range pos {
		pint, _ := strconv.Atoi(p)
		positions = append(positions, pint)
	}

	sort.Ints(positions)

	ansP1 := 0
	median := positions[len(positions) / 2]

	for _, p := range positions {
		ansP1 += int(math.Abs(float64(p) - float64(median)))
	}
	fmt.Println(ansP1)

	ansP2 := math.MaxInt32
	for pos := range (positions) {
		score := 0
		for _, p := range positions {
			s := int(math.Abs(float64(p) - float64(pos)))
			// each step adrs 1 cost to fuel
			// 1 + 2 + 3 + ... + s == s * (s + 1) / 2
			score += s*(s + 1) / 2
		}

		if score < ansP2 {
			ansP2 = score
		}
	}

	fmt.Println(ansP2)
}