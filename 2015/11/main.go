package main

import (
	"fmt"
)

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
func isValid(pass []byte) bool {
	threeConsecutite := false
	countDifferentPairs := 0
	lastPairIndex := 0
	for i := 0; i < len(pass); i++ {
		if pass[i] == byte(105) || pass[i] == byte(108) || pass[i] == byte(111) {
			return false
		}
		if i >= 2 {
			if pass[i - 2] + 2 == pass[i - 1] + 1 && pass[i - 1] + 1 == pass[i] {
				threeConsecutite = true
			}
		}
		if i >= 1 {
			if pass[i - 1] == pass[i] && lastPairIndex + 1 != i {
				countDifferentPairs++
				lastPairIndex = i
			}
		}
	}

	return threeConsecutite && countDifferentPairs > 1
}

func nextPass(pass []byte) []byte {
	index := len(pass) - 1
	increment := pass[index]

	for increment == 122 {
		increment = 97
		pass[index] = increment
		index--
		increment = pass[index]
	}

	increment++
	pass[index] = increment

	return pass
}

func main() {
	input := []byte("vzbxkghb")

	for ! isValid([]byte(input)) {
		input = nextPass(input)
	}

	fmt.Println(string(input))

	input = nextPass(input)
	for ! isValid([]byte(input)) {
		input = nextPass(input)
	}

	fmt.Println(string(input))
}
