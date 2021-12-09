package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0
	ansP2 := 0

	for _, line := range(lines) {
		// fmt.Println(line)

		parts := strings.Split(line, " | ")

		signals := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		for _, d := range(output) {
			// 1 4, 7, 8
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				ansP1 += 1
			}

			ansP2 += calculateLineSum(signals, output)
		}
	}

	// m := map[int][]string{}


	fmt.Println(ansP1)
	fmt.Println(ansP2)
}


func calculateLineSum(signals []string, output []string) int {
	m := map[int]map[string]bool{}

	for _, signal := range(signals) {
		digit := map[string]bool{}

		if len(signal) == 2 {
			digit["c"] = true;
			digit["f"] = true;

			m[1] = digit;
		}

		if len(signal) == 3 {
			digit["a"] = true;
			digit["c"] = true;
			digit["f"] = true;

			m[1] = digit;
		}

		if len(signal) == 2 {
			digit["c"] = true;
			digit["f"] = true;

			m[1] = digit;
		}

		if len(signal) == 2 {
			digit["c"] = true;
			digit["f"] = true;

			m[1] = digit;
		}
	}
	return 0;
}