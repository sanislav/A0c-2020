package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func solveP1(inputString []string) int {
	earliest, _ := strconv.Atoi(inputString[0])
	originalEarlyest := earliest
	ts := make([]int, 0)

	for _, c := range(strings.Split(inputString[1], ",")) {
		if string(c) == "x" {
			continue
		}
		val, _ := strconv.Atoi(c)
		ts = append(ts, val)
	}

	bid := 0
	found := false

	for ! found {
		for _, t := range(ts) {
			if earliest % t == 0 {
				bid = t
				found = true
				break
			}
		}
		if ! found {
			earliest++
		}
	}

	return bid * (earliest - originalEarlyest)
}

func solveP2(inputString []string) int {
	ans, step := 0, 1

	for i, s := range strings.Split(inputString[1], ",") {
		bus, _ := strconv.Atoi(s)

		if (bus == 0) {
			continue
		}

		for (ans+i)%bus != 0 {
			ans += step
		}

		step *= bus
	}

	return ans
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	inputString := strings.Split(strings.TrimSpace(string(input)), "\n")

	ans := solveP1(inputString)
	fmt.Println("P1", ans)

	ans = solveP2(inputString)
	fmt.Println("P2", ans)
}
