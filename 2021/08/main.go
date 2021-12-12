package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0
	ansP2 := 0

	for _, line := range(lines) {
		// fmt.Println(line)

		parts := strings.Split(line, " | ")

		signals := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		for _, d := range(output) {
			// 1 4, 7, 8
			if len(d) == 2 || len(d) == 4 || len(d) == 3 || len(d) == 7 {
				ansP1 += 1
			}

			ansP2 += calculateLineSum(signals, output)
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}


func calculateLineSum(signals []string, output []string) int {
	digitList := [10]string{}

	for _, signal := range(signals) {
		if len(signal) == 2 {
			digitList[1] = signal
		} else if len(signal) == 3 {
			digitList[7] = signal
		} else if len(signal) == 4 {
			digitList[4] = signal
		} else if len(signal) == 5 {
			if len(digitList[2]) == 0 {
				digitList[2] = signal
			} else {
				if len(digitList[3]) == 0 {
					digitList[3] = signal
				} else {
					digitList[5] = signal
				}
			}
		} else if len(signal) == 6 {
			if len(digitList[0]) == 0 {
				digitList[0] = signal
			} else {
				if len(digitList[6]) == 0 {
					digitList[6] = signal
				} else {
					digitList[9] = signal
				}
			}
		} else if len(signal) == 7 {
			digitList[8] = signal
		}
	}

	digitSegments := [7]string{}
	digitSegments[0] = substract(digitList[7], digitList[1])
	o1 := intersect(intersect(digitList[2], digitList[3]), digitList[5])
	o2 := intersect(intersect(digitList[0], digitList[6]), digitList[9])
	digitSegments[1] = intersect(substract(digitList[4], digitList[1]), o2)
	digitSegments[3] = intersect(substract(digitList[4], digitList[1]), o1)
	digitSegments[6] = intersect(substract(substract(digitList[8], digitList[4]), digitSegments[0]), o1)
	digitSegments[4] = substract(substract(substract(digitList[8], digitList[4]), digitSegments[0]), digitSegments[6])
	digitSegments[5] = substract(substract(substract(o2, digitSegments[0]), digitSegments[1]), digitSegments[6])
	digitSegments[2] = substract(digitList[1], digitSegments[5])

	fmt.Println(digitSegments)
	// fmt.Println(digitSegments[2])
	return 0;
}

func substract(a string, b string) string {
	diff := ""
	for _, c := range a {
		found := strings.Contains(b, string(c))

		if (! found) {
			diff += string(c)
		}
	}

	return diff
}


func intersect(a string, b string) string {
	common := ""
	for _, c := range a {
		found := strings.Contains(b, string(c))

		if (found) {
			common += string(c)
		}
	}

	return common
}