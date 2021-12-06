package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Matrix struct {
	data [1000][1000]int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fishDaysStr := strings.Split(lines[0], ",")
	fish := []int{}
	daysFish := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}

	for _, f := range (fishDaysStr) {
		v, _ := strconv.Atoi(f)
		fish = append(fish, v)

		if _, ok := daysFish[v]; ok {
			daysFish[v] += 1
		} else {
			daysFish[v] = 0
		}
	}

	for day := 0; day < 256; day++ {
		fishy := map[int]int{}

		for v := range daysFish {
			if v == 0 {
				fishy[6] += daysFish[v]
				fishy[8] += daysFish[v]
			} else {
				fishy[v - 1] += daysFish[v]
			}
		}

		for m := range fishy {
			daysFish[m] = fishy[m]
		}
	}

	ans := 0
	for m := range daysFish {
		ans += daysFish[m]
	}

	fmt.Println(ans)
}
