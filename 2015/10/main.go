package main

import (
	"fmt"
)

func lastIndexWithSameChar(s []byte, startIndex int) int {
	for i := startIndex + 1; i < len(s); i++ {
		if s[i] != s[startIndex] {
			return i
		}
	}

	return len(s)
}

func main() {
	input := []byte("1113122113")

	ansP1 := 0
	ansP2 := 0

	generations := 50

	for {
	 	if generations == 0 {
			break
		}

		newInput := []byte{}

		for i := 0; i < len(input); {
			lastIndex := lastIndexWithSameChar(input, i)
			groupSize := byte(lastIndex - i)

			newInput = append(newInput, groupSize + byte(48), input[i])

			i = lastIndex
		}

		input = newInput

		if (generations == 11) {
			ansP1 = len(input)
		}

		generations--
	}

	fmt.Println(ansP1)
	ansP2 = len(input)
	fmt.Println(ansP2)
}
