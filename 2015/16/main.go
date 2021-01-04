package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := ""
	ansP2 := ""

	r := regexp.MustCompile("\\w+[0-9]*")
	matchingSue := map[string]int{"children": 3,
		"cats": 7,
		"samoyeds": 2,
		"pomeranians": 3,
		"akitas": 0,
		"vizslas": 0,
		"goldfish": 5,
		"trees": 3,
		"cars": 2,
		"perfumes": 1,
	}

	for _, s := range lines {
		matches := r.FindAllString(s, -1)

		info := map[string]int{}
		info[matches[2]], _ = strconv.Atoi(matches[3])
		info[matches[4]], _ = strconv.Atoi(matches[5])
		info[matches[6]], _ = strconv.Atoi(matches[7])

		matchedP1 := true
		matchedP2 := true
		for k, v := range info {
			if k == "cats" || k == "trees" {
				if matchingSue[k] >= v{
					matchedP2 = false
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if matchingSue[k] <= v{
					matchedP2 = false
				}
			}

			if matchingSue[k] != v{
				matchedP1 = false
				if k != "cats" && k != "trees" && k != "pomeranians" && k != "goldfish" {
					matchedP2 = false
				}
			}

			if matchedP1 == false && matchedP2 == false {
				break
			}
		}

		if matchedP1 && len(ansP1) == 0 {
			ansP1 = matches[1]
		}
		if matchedP2 && len(ansP2) == 0 {
			ansP2 = matches[1]
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
