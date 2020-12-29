package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)


func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ans := 0
	ribbon := 0
	for _, l := range lines {
		parts := strings.Split(l, "x")

		l, _ := strconv.Atoi(parts[0])
		w, _ := strconv.Atoi(parts[1])
		h, _ := strconv.Atoi(parts[2])
		order := []int{l, w, h}
		sort.Ints(order)

		side1 := l*w
		side2 := w*h
		side3 := l*h
		slack := side1

		if side2 < slack {
			slack = side2
		}
		if side3 < slack {
			slack = side3
		}

		ribbon += 2 * order[0] + 2 * order[1] + w * h * l
		ans += 2*side1 + 2*side2 + 2*side3 + slack
	}

	fmt.Println(ans)
	fmt.Println(ribbon)
}
