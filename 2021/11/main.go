package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const ml = 10
const mc = 10

var flashCount = 0
var octo = [ml][mc]int{}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	for i := 0; i < ml; i++ {
		for j:= 0; j < mc; j++ {
			octo[i][j], _ = strconv.Atoi(string(lines[i][j]))
		}
	}

	c := 0
	for true {
		c += 1
		for i := 0; i < ml; i++ {
			for j:= 0; j < mc; j++ {
				octo[i][j] += 1
			}
		}
		for i := 0; i < ml; i++ {
			for j:= 0; j < mc; j++ {
				if octo[i][j] == 10 {
					flash(i,j)
				}
			}
		}

		done := true
		for i := 0; i < ml; i++ {
			for j:= 0; j < mc; j++ {
				if octo[i][j] == -1 {
					octo[i][j] = 0
				} else {
					done = false
				}
			}
		}

		if c == 100 {
			fmt.Println(flashCount)
		}
		if done {
			fmt.Println(c)
			break
		}
	}
}


func flash(i int, j int) {
	flashCount += 1
	octo[i][j] = -1
	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			rr := i + dr
			cc := j + dc
			if 0 <= rr && rr < ml && 0 <= cc && cc < mc && octo[rr][cc] != -1 {
				octo[rr][cc] += 1
				if octo[rr][cc] >= 10 {
					flash(rr, cc)
				}
			}
		}
	}
}
