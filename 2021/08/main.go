package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Matrix struct {
	data [1000][1000]int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0

	for _, line := range(lines) {
		// fmt.Println(line)

		parts := strings.Split(line, " | ")

		// siglans := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		for _, d := range(output) {
			// 1 4, 7, 8

			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				ansP1 += 1
			}
		}
	}

	// m := map[int][]string{}


	fmt.Println(ansP1)
}