package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)


func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0

	chars := 0
	memoryChars := 0
	memoryCharsQ := 0
	for _, s := range lines {
		chars += len(s)

		unquoted, _ := strconv.Unquote(s)
		memoryChars += len(unquoted)

		quoted := strconv.Quote(s)
		memoryCharsQ += len(quoted)
	}

	ansP1 = chars - memoryChars
	ansP2 = memoryCharsQ - chars
	fmt.Println(ansP1)
	fmt.Println(ansP2)
}