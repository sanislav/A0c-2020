package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	gama := ""
	epsilon := ""
	counts := getCounts(lines)

	for i := 0; i < len(counts); i++ {
		if (counts[i] < 0) {
			epsilon += "0"
			gama += "1"
		} else {
			epsilon += "1"
			gama += "0"
		}

	}

	e, _ := strconv.ParseInt(epsilon, 2, 64)
	g, _ := strconv.ParseInt(gama, 2, 64)
	fmt.Println(e * g)

	pos := 0
	linesOxygen := lines
	linesScrubbing := lines

	for len(linesOxygen) > 1 {
		counts := getCounts(linesOxygen)
		newlines := []string{}
		for _, l := range linesOxygen {
			if (l[pos] == 48 && counts[pos] < 0 || l[pos] == 49 && counts[pos] >= 0) {
				newlines = append(newlines, l)
				continue;
			}
		}

		linesOxygen = newlines
		pos++
	}

	pos = 0
	for len(linesScrubbing) > 1 {
		counts := getCounts(linesScrubbing)
		newlines := []string{}
		for _, l := range linesScrubbing {
			if (l[pos] == 48 && counts[pos] >= 0 || l[pos] == 49 && counts[pos] < 0) {
				newlines = append(newlines, l)
				continue;
			}
		}

		linesScrubbing = newlines
		pos++
	}

	o, _ := strconv.ParseInt(linesOxygen[0], 2, 64)
	s, _ := strconv.ParseInt(linesScrubbing[0], 2, 64)
	fmt.Println(o * s)
}

func getCounts(lines []string) map[int]int {
	counts := map[int]int{}

	for _, l := range lines {
		for p, c := range l {
			if (c == 48) {
				counts[p]--
			} else {
				counts[p]++
			}
		}
	}

	return counts
}