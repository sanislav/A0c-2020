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
	const l = 20
	okIntervals := [l][l]bool{}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
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
	dict := make(map[int]int, 0)
	for len(dict) != len(okIntervals) {
		for lineInd, i := range(okIntervals) {
			c := 0
			ind := -1
			for trueInd, j := range(i) {
				if j == true {
					if _, ok := dict[trueInd]; ok {
						continue;
					} else {
						c++
						ind = trueInd
					}
				}
				if c > 1 {
					break
				}
			}
			if c == 1 {
				dict[ind] = lineInd
			}
		}
	}

	ansP2 := 1
	myTicket := strings.Split(strings.ReplaceAll(split[1], "your ticket:\n", ""), ",");

	for i := 0; i < 6; i++ {
		pos := dict[i]
		ticketVal, _ := strconv.Atoi(myTicket[pos])
		ansP2 *= ticketVal
	}

	fmt.Println(ansP2)
}