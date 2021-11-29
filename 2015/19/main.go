package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	// "strconv"
	// "sort"
)


func main() {
	input, _ := ioutil.ReadFile("input.txt")
	parts := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	lines := strings.Split(strings.TrimSpace(string(parts[0])), "\n")
	ansP1 := 0
	ansP2 := 0

	combinations := map[string]int{}
	repl := map[string][]string{}
	molecule := parts[1]

	for _, s := range lines {
		mapping := strings.Split(strings.TrimSpace(string(s)), " => ")
		repl[mapping[0]] = append(repl[mapping[0]], mapping[1]);
	}

	for i := 0; i < len(molecule); i++ {
		for k, values := range repl {
			for _, v := range values {
				offset := i + len(k)
				if parts[1][i:offset] == k {
					parts[1] = parts[1][:i] + v + parts[1][offset:]
					combinations[parts[1]] = 1
				}
			}
		}
	}

	ansP1 = len(combinations)
	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
