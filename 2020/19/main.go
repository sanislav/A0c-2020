package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func test(t string, ruleNo int, dict map[int]string) []int {
	rule := dict[ruleNo]
	if (string(rule[0]) == "\"") {
		rule = strings.ReplaceAll(rule, "\"", "")
		if strings.HasPrefix(t, rule) {
			re := make([]int, 0)
			return append(re, len(rule))
		}

		return make([]int, 0)
	}

	result := make([]int, 0)
	parts := strings.Split(rule, " | ")
	for _, option := range(parts) {
		acc := make([]int, 0)
		acc = append(acc, 0)

		subrules := strings.Split(option, " ")

		for _, rn := range(subrules) {
			nacc := make([]int, 0)
			intRuleNo, _ := strconv.Atoi(rn)

			for _, ac := range(acc) {
				res := test(t[ac:], intRuleNo, dict)
				for _, c := range(res) {
					nacc = append(nacc, c + ac)
				}
			}
			acc = nacc
		}

		result = append(result, acc...)
	}

	return result
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	rules := strings.Split(parts[0], "\n")
	tests := strings.Split(parts[1], "\n")
	dict := make(map[int]string, 0)
	dict2 := make(map[int]string, 0)

	for _, rule := range(rules) {
		splitParts := strings.Split(string(rule), ": ")
		ruleNo, _ := strconv.Atoi(splitParts[0])
		dict[ruleNo] = splitParts[1]
		if (ruleNo == 8) {
			dict2[ruleNo] = "42 | 42 8"
		} else if (ruleNo == 11) {
			dict2[ruleNo] = "42 31 | 42 11 31"
		} else {
			dict2[ruleNo] = splitParts[1]
		}
	}

	ansP1 := 0
	ansP2 := 0
	for _, t := range(tests) {
		for _, r := range(test(t, 0, dict)) {
			if len(t) == r {
				ansP1++
			}
		}

		for _, r := range(test(t, 0, dict2)) {
			if len(t) == r {
				ansP2++
			}
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}