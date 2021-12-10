package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var digits = map[int]map[string]bool{
    0: {
		"a": true,
		"c": true,
		"f": true,
		"g": true,
		"e": true,
		"b": true,
	},
	1: {
		"c": true,
		"f": true,
	},
	2: {
		"a": true,
		"c": true,
		"d": true,
		"e": true,
		"g": true,
	},
	3: {
		"a": true,
		"c": true,
		"d": true,
		"f": true,
		"g": true,
	},
	4: {
		"b": true,
		"d": true,
		"c": true,
		"f": true,
	},
	5: {
		"a": true,
		"b": true,
		"d": true,
		"f": true,
		"g": true,
	},
	6: {
		"a": true,
		"b": true,
		"d": true,
		"f": true,
		"g": true,
		"e": true,
	},
	7: {
		"a": true,
		"c": true,
		"f": true,
	},
	8: {
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"e": true,
		"f": true,
		"g": true,
	},
	9: {
		"a": true,
		"b": true,
		"c": true,
		"d": true,
		"f": true,
		"g": true,
	},
}

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
	for _, signal := range(signals) {
		fmt.Println(signal)
		found := true
		for i := 0; i <= 9; i++ {
			if len(digits[i]) == len(signal) {
				for _, c := range signal {
					if _, ok := digits[i][c]; !ok {
						found = false
						break
					}
				}
				if found
				i
			}
		}
	}
	return 0;
}