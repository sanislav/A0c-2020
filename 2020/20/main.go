package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func flip(s string) string {
	r := ""
	for _, c := range s {
		r = string(c) + r
	}

	return r
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	sensors := map[string][]string{}
	for _, s := range sections {
		parts := strings.Split(s, "\n")
		id := strings.Split(strings.TrimRight(parts[0], ":"), "Tile ")

		top := string(parts[1])
		bottom := string(parts[10])
		left := ""
		right := ""

		for i := 1; i <= 10; i++ {
			left += string(parts[i][0])
			right += string(parts[i][9])
		}

		sensors[id[1]] = []string{top, flip(top), bottom, flip(bottom), left, flip(left), right, flip (right)}
	}

	m := map[string][]string{}
	for i1, e1 := range sensors {
		for i2, e2 := range sensors {
			if i1 != i2 {
				for _, v1 := range e1 {
					for _, v2 := range e2 {
						if v1 == v2 {
							alreadyExists := false
							for _, v := range(m[i1]) {
								if v == i2 {
									alreadyExists = true
									break
								}
							}
							if ! alreadyExists {
								m[i1] = append(m[i1], i2)
							}
						}
					}
				}
			}
		}
	}

	ans := 1
	for id, l := range m {
		if len(l) == 2 {
			intID, _ := strconv.Atoi(id)
			ans *= intID
		}
	}
	// fmt.Println(m)

	fmt.Println(ans)
}