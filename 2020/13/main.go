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

	for {
		for _, t := range(strings.Split(inputString[1], ",")) {
			bus, _ := strconv.Atoi(t)

			if (bus == 0) {
				continue
			}

			if earliest % bus == 0 {
				return bus * (earliest - originalEarlyest)
			}
		}

		earliest += 1
	}
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
