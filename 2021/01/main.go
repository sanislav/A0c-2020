package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0
	beforeP1 := 0
	beforeP2 := 0
	groups := map[int]int{}
	keys := make([]int, 0)

	for i, l := range lines {
		mesasurement, _ := strconv.Atoi(l)

		if (mesasurement == 0) { // last line
			break
		}

		if (i + 2 < len(lines)) {
			mesasurement2, _ := strconv.Atoi(lines[i+1])
			mesasurement3, _ := strconv.Atoi(lines[i+2])

			groups[i] = mesasurement + mesasurement2 + mesasurement3
			keys = append(keys, i)
		}

		if beforeP1 == 0 {
			beforeP1 = mesasurement
			continue
		}

		if mesasurement > beforeP1 {
			ansP1++;
		}

		beforeP1 = mesasurement
	}

	sort.Ints(keys)

	for _, k := range keys {
		if beforeP2 == 0 {
			beforeP2 = groups[k]
			continue
		}

		if groups[k] > beforeP2 {
			ansP2++;
		}

		beforeP2 = groups[k]

	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
