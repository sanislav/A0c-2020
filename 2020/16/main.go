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
	for _, s := range strings.Split(split[2], "\n") {
		if (s == "nearby tickets:") {
			continue
		}
		validTicket := []int{}
		// isValid := true
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
				// isValid = false
				break
			}
		}

		// if (isValid) {
		// 	for i, v := range(validTicket) {
		// }
	}

	fmt.Println(ansP1)
}