package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func intersect(s1 []string, s2 []string) []string {
	if len(s1) == 0 {
		return s2
	}

	intersection := make([]string, 0)
	for _, v := range s2 {
		exists := false
		for _, v2 := range(s1) {
			if v2 == v {
				exists = true
				break
			}
		}

		if exists {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

func appendIfMissing(slice []string, s string) []string {
    for _, ele := range slice {
        if ele == s {
            return slice
        }
    }
    return append(slice, s)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	alToIng := make(map[string][]string, 0)
	allIngredients := make([]string, 0)

	for _, l := range(lines) {
		parts := strings.Split(l, "(contains ")
		ingredients := strings.Split(strings.TrimRight(parts[0], " "), " ")
		alergens := strings.Split(strings.TrimRight(parts[1], ")"), ", ")
		for _, al := range(alergens) {
			alToIng[al] = intersect(alToIng[al], ingredients)
		}

		for _, ing := range ingredients {
			allIngredients = append(allIngredients, ing)
		}
	}

	fmt.Println(alToIng)


	// filter out known aleren ingredients from other alergens that might be it
	dict := make(map[string]string, 0)
	for len(dict) < len(alToIng) {
		for alergen, ingredients := range(alToIng) {
			if (len(ingredients) == 1) {
				dict[ingredients[0]] = alergen
				continue
			}

			filteredIng := make([]string, 0)
			for i, ing := range ingredients {
				_, ok := dict[ing]
				if ! ok {
					filteredIng = append(filteredIng, ing)
				} else {
					if (len(ingredients) > i) {
						ingredients = append(ingredients[:i], ingredients[i + 1:]...)
					}

					if (len(ingredients) == 1) {
						dict[ingredients[0]] = alergen
						continue
					}
				}
			}

			if (len(filteredIng) == 1) {
				dict[filteredIng[0]] = alergen
				continue
			}

			alToIng[alergen] = filteredIng
		}
	}

	ansP1 := 0
	ansP2 := ""
	for _, i := range allIngredients {
		_, exists := dict[i]
		if ! exists {
			ansP1++
		}
	}

	keys := make([]string, 0)
    for _, al := range dict {
        keys = append(keys, al)
	}
	sort.Strings(keys)
	fmt.Println(dict)

	p2 := make([]string, 0)
	for _, a := range keys {
		for ing, al := range dict {
			if a == al {
				p2 = appendIfMissing(p2, ing)
			}
		}
	}

	ansP2 += strings.Join(p2, ",")
	fmt.Println(ansP1, ansP2)
}