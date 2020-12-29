package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func isValidP1(s string) bool {
	numVouls := 0
	twiceInRow := false
	notContainsStringPairs := true

	for i, cr := range s {
		c := string(cr)
		if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
			numVouls++
		}

		if (i > 0) {
			prev := string(s[i-1])
			if (c == prev) {
				twiceInRow = true
			}

			pair := prev + c
			if pair == "ab" || pair == "cd"  || pair == "pq"  || pair == "xy" {
				notContainsStringPairs = false
			}
		}
	}

	return numVouls >= 3 && twiceInRow && notContainsStringPairs
}

func isValidP2(s string) bool {
	pairsRepeat := false
	letterRepeats := false

	pairs := map[string]int{}

	for i, cr := range s {
		c := string(cr)
		if i > 1 {
			if c == string(s[i-2]) {
				letterRepeats = true
			}
		}

		if (i > 0) {
			pair := c + string(s[i-1])

			if lastIndex, exists := pairs[pair]; exists {
				if lastIndex != i-2 {
					pairsRepeat = true
				}
			} else {
				pairs[pair] = i - 1
			}
		}
	}

	return pairsRepeat && letterRepeats
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0

	for _, s := range lines {
		if isValidP1(s) {
			ansP1++
		}
		if isValidP2(s) {
			ansP2++
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
