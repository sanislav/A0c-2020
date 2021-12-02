package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	x := 0
	z := 0
	aim := 0
	zP2 := 0

	for _, l := range lines {
		validFwd := regexp.MustCompile(`forward (\d+)`)
		matchedFwd := validFwd.FindString(l)
		if len(matchedFwd) > 0 {
			fwd := strings.Split(matchedFwd, " ")
			xInt, _ := strconv.Atoi(fwd[1])

			x += xInt
			zP2 += xInt * aim
		}

		validUp := regexp.MustCompile(`up (\d+)`)
		matchedUp := validUp.FindString(l)
		if len(matchedUp) > 0 {
			up := strings.Split(matchedUp, " ")
			zInt, _ := strconv.Atoi(up[1])

			z -= zInt
			aim -= zInt
		}

		validDown := regexp.MustCompile(`down (\d+)`)
		matchedDown := validDown.FindString(l)
		if len(matchedDown) > 0 {
			down := strings.Split(matchedDown, " ")
			zInt, _ := strconv.Atoi(down[1])

			z += zInt
			aim += zInt
		}
	}
	
	fmt.Println(x * z)
	fmt.Println(x * zP2)
}
