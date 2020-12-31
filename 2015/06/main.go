package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

func getRanges(s string) (x1, y1, x2, y2 int) {
	digits := regexp.MustCompile("(\\d+)")
	matches := digits.FindAllString(s, -1)

	x1, _ = strconv.Atoi(matches[0])
	y1, _ = strconv.Atoi(matches[1])
	x2, _ = strconv.Atoi(matches[2])
	y2, _ = strconv.Atoi(matches[3])

	return x1, y1, x2, y2

}

func getAction(s string) (action string) {
	if strings.Contains(s, "on") {
		action = "on"
	} else if strings.Contains(s, "off") {
		action = "off"
	} else {
		action = "toggle"
	}

	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1, ansP2 := 0, 0
	grid := [1000][1000]bool{}
	gridP2 := [1000][1000]int{}

	for _, s := range lines {
		x1, y1, x2, y2 := getRanges(s)
		action := getAction(s)

		for i := x1; i <= x2; i++ {
			for j := y1; j <= y2; j++ {
				if action == "on" {
					grid[i][j] = true
					gridP2[i][j]++
				} else if action == "off" {
					grid[i][j] = false
					gridP2[i][j]--
					if gridP2[i][j] < 0 {
						gridP2[i][j] = 0
					}
				} else {
					if grid[i][j] == true {
						grid[i][j] = false
					} else {
						grid[i][j] = true
					}
					gridP2[i][j] += 2
				}
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == true {
				ansP1++
			}
			ansP2 += gridP2[i][j]
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
