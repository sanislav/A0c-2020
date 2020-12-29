package main

import (
	"fmt"
	"io/ioutil"
	// "strings"
	"strconv"
	// "sort"
)

func hash(x, y int) string {
	return strconv.Itoa(x) + "_" + strconv.Itoa(y)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	grid := map[string]bool{}
	gridP2 := map[string]bool{}

	x, y := 0, 0

	xS, yS := 0, 0
	xRS, yRS := 0, 0

	grid[hash(x, y)] = true
	gridP2[hash(xS, yS)] = true

	for i, c := range input {
		if string(c) == "^" {
			y++
			if i % 2 == 0 {
				yS++
			} else {
				yRS++
			}
		} else if string(c) == ">" {
			x++
			if i % 2 == 0 {
				xS++
			} else {
				xRS++
			}
		} else if string(c) == "v" {
			y--
			if i % 2 == 0 {
				yS--
			} else {
				yRS--
			}
		} else {
			x--
			if i % 2 == 0 {
				xS--
			} else {
				xRS--
			}
		}

		grid[hash(x, y)] = true
		if i % 2 == 0 {
			gridP2[hash(xS, yS)] = true
		} else {
			gridP2[hash(xRS, yRS)] = true
		}
	}

	fmt.Println(len(grid))
	fmt.Println(len(gridP2))
}
