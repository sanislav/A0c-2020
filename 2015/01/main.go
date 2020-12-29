package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	ans := 0
	basementPos := 0
	for i, c := range input {
		if string(c) == "(" {
			ans++
		} else {
			ans--
		}
		if ans < 0 && basementPos == 0 {
			basementPos = i + 1
		}
	}
	fmt.Println(ans)
	fmt.Println(basementPos)
}
