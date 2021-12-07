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

	best := math.MaxInt32
	for pos := range (positions) {
		score := 0
		for _, p := range positions {
			s := int(math.Abs(float64(p) - float64(pos)))
			for i := 1; i <= s; i++ {
				score += i
			}
		}

		if score < best {
			best = score
		}
	}

	fmt.Println(best)
}