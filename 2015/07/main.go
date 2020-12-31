package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)


type Link struct {
	input1, input2, operation string
}

var circuit = map[string]Link{}
var memo = map[string]uint16{}

func compute(label string) (v uint16) {
	// if we get a number
	if v, err := strconv.ParseUint(label, 10, 16); err == nil {
		return uint16(v)
	}

	if v, exists := memo[label]; exists {
		return v
	}

	link := circuit[label]

	switch link.operation{
		case "":
			v = compute(link.input1)
		case "AND":
			v = compute(link.input1) & compute(link.input2)
		case "OR":
			v = compute(link.input1) | compute(link.input2)
		case "LSHIFT":
			v = compute(link.input1) << compute(link.input2)
		case "RSHIFT":
			v = compute(link.input1) >> compute(link.input2)
		case "NOT":
			v = ^compute(link.input1)
	}

	memo[label] = v

	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := uint16(0)
	ansP2 := uint16(0)

	for _, s := range lines {
		var w1, w2, op, dest string

		if n, _ := fmt.Sscanf(s, "%s -> %s\n", &w1, &dest); n == 2 {
			op = ""
		} else if n, _ := fmt.Sscanf(s, "%s AND %s -> %s\n", &w1, &w2, &dest); n == 3 {
			op = "AND"
		} else if n, _ := fmt.Sscanf(s, "%s OR %s -> %s\n", &w1, &w2, &dest); n == 3 {
			op = "OR"
		} else if n, _ := fmt.Sscanf(s, "%s RSHIFT %s -> %s\n", &w1, &w2, &dest); n == 3 {
			op = "RSHIFT"
		} else if n, _ := fmt.Sscanf(s, "%s LSHIFT %s -> %s\n", &w1, &w2, &dest); n == 3 {
			op = "LSHIFT"
		} else if n, _ := fmt.Sscanf(s, "NOT %s -> %s\n", &w1, &dest); n == 2 {
			op = "NOT"
		}

		circuit[dest] = Link{
			input1: w1,
			input2: w2,
			operation: op,
		}
	}

	ansP1 = compute("a")
	fmt.Println(ansP1)

	memo = map[string]uint16{}
	memo["b"] = uint16(ansP1)
	ansP2 = compute("a")

	fmt.Println(ansP2)
}
