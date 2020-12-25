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
	card, _ := strconv.Atoi(lines[0])
	door, _ := strconv.Atoi(lines[1])

	loopSize := 0
	k := 1
	for k != card {
		k = 7 * k % 20201227
		loopSize++
	}

	key := 1
	for l := 0; l < loopSize; l++ {
		key = key * door % 20201227
	}

	fmt.Println(key)
}