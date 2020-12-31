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
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0

	for _, s := range lines {
		fmt.Println(s)
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
