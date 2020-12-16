package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func solveP1(split []string, rules map[string][]int, mapStruct map[string]map[int]struct{}) int {
	ans := 0

	tickets:
	for _, s := range strings.Split(split[2], "\n")[1:] {
	fields:
		for _, s := range strings.Split(s, ",") {
			n, _ := strconv.Atoi(s)
			for _, v := range rules {
				if n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3] {
					continue fields
				}
			}
			ans += n
			continue tickets
		}

		for i, s := range strings.Split(s, ",") {
			for k, v := range rules {
				if n, _ := strconv.Atoi(s); !(n >= v[0] && n <= v[1] || n >= v[2] && n <= v[3]) {
					delete(mapStruct[k], i)
				}
			}
		}
	}

	return ans
}

func solveP2(split []string, mapStruct map[string]map[int]struct{}) int {
	ans := 1

	for len(mapStruct) > 0 {
		for k, v := range mapStruct {
			if len(v) != 1 {
				continue
			}

			for i := range v {
				for k := range mapStruct {
					delete(mapStruct[k], i)
				}
				delete(mapStruct, k)

				if strings.HasPrefix(k, "departure") {
					n, _ := strconv.Atoi(strings.Split(strings.Split(split[1], "\n")[1], ",")[i])
					ans *= n
				}
			}
		}
	}

	return ans
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	rules := map[string][]int{}
	for _, s := range strings.Split(split[0], "\n") {
		rule := strings.Split(s, ": ")
		rules[rule[0]] = make([]int, 4)
		fmt.Sscanf(rule[1], "%d-%d or %d-%d", &rules[rule[0]][0], &rules[rule[0]][1], &rules[rule[0]][2], &rules[rule[0]][3])
	}

	mapStruct := map[string]map[int]struct{}{}
	for k := range rules {
		mapStruct[k] = map[int]struct{}{}
		for i := 0; i < len(rules); i++ {
			mapStruct[k][i] = struct{}{}
		}
	}

	ans := solveP1(split, rules, mapStruct)

	fmt.Println(ans)

	ans = solveP2(split, mapStruct)

	fmt.Println(ans)
}