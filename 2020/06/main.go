package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func solve() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	sumEveryone := 0
	annyoneYes := make(map[string]bool)
	everyoneYes := make(map[string]bool)
	newGroup := true

	for scanner.Scan() {
		lineText := scanner.Text()

		if (len(lineText) == 0) {
			sumEveryone += len(everyoneYes)
			annyoneYes = make(map[string]bool)
			everyoneYes = make(map[string]bool)
			newGroup = true
			continue
		}

		for c := 0; c < len(lineText); c++ {
			char := string(lineText[c])

			if ! annyoneYes[char] {
				annyoneYes[char] = true
				sum ++
			}

			if newGroup {
				everyoneYes[char] = true
			}
		}

		if ! newGroup {
			for k := range everyoneYes {
				if ! strings.Contains(lineText, k) {
					delete(everyoneYes, k)
				}
			}
		}

		newGroup = false
	}

	sumEveryone += len(everyoneYes)

	fmt.Println("Questions answered by everyone:", sumEveryone)
	fmt.Println("Questions answered by anyone:", sum)
}


func main() {
	solve()
}