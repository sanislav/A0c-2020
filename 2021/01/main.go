package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0

	for i := range lines {
		if (i == 0) {
			continue
		}
		current, _ := strconv.Atoi(lines[i])
		previous, _ := strconv.Atoi(lines[i-1])

		if (current > previous) {
			ansP1++;
		}

		if (i >= 3) {
			previous3, _ := strconv.Atoi(lines[i-3])
			if (current > previous3) {
				ansP2++;
			}
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
