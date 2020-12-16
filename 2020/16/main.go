package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"regexp"
	"strings"
)


func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	// limit rules section
	limits := [][]int{}
	for _, s := range strings.Split(split[0], "\n") {
		limRule := regexp.MustCompile("\\d+")
		lim := limRule.FindAllString(s, -1)
		limInt := []int{}
		for _, i := range lim {
			j, _ := strconv.Atoi(i)
			limInt = append(limInt, j)
		}

		limits = append(limits, limInt)
	}

	// nearby tickts section
	ansP1 := 0

	// okIntervals := make([][]bool, 20)
	okIntervals := [20][20]bool{}

	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			okIntervals[i][j] = true
		}
	}

	for _, s := range strings.Split(split[2], "\n") {
		if (s == "nearby tickets:") {
			continue
		}
		validTicket := []int{}
		isValid := true
		vs := strings.Split(s, ",")

		for _, i := range vs {
			j, _ := strconv.Atoi(i)
			validTicket = append(validTicket, j)
		}

		for _, v := range(validTicket) {
			valid := false

			for _, limit := range(limits) {
				if (limit[0] <= v && v <= limit[1]) || (limit[2] <= v && v <= limit[3]) {
					valid = true
					break
				}
			}

			if ! valid {
				ansP1 += v
				isValid = false
				break
			}
		}

		if (isValid) {
			for i, v := range(validTicket) {
				for j, limit := range(limits) {
					if ! ((limit[0] <= v && v <= limit[1]) || (limit[2] <= v && v <= limit[3])) {
						okIntervals[i][j] = false
					}
				}
			}
		}
	}

	fmt.Println(ansP1)

	// figure out which position is which field
	fmt.Println(okIntervals)
}