package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
	"regexp"
)

// DP formula: g(start, {nodes...}) = min(Cost[start][k] + g(k, {nodes...}-{k})) for each k in {nodes}
func travellingSalesperson(start string,  locations map[string]bool, distances map[string]map[string]int, memo map[string]int, isMin bool) int {
	key := start
	for k := range locations {
		key += "_" + k
	}
	if v, exists := memo[key]; exists {
        return v
	}

	// base case - we reached a leaf - no need to go back to start
	if len(locations) == 0 {
		return 0
	}

    values := []int{}

    for k := range locations {
		subList := map[string]bool{}

		for k2 := range locations {
			if k != k2 {
				subList[k2] = true
			}
		}

        result := travellingSalesperson(k, subList, distances, memo, isMin)
        values = append(values, distances[start][k] + result)
	}

	sort.Ints(values)
	if isMin == true {
		memo[key] = values[0]
	} else {
		memo[key] = values[len(values) -1 ]
	}

    return memo[key]
}


func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0

	distances := map[string]map[string]int{}
	locations := map[string]bool{}
	for _, s := range lines {
		cities := regexp.MustCompile("(\\w+) to (\\w+)")
		matches := cities.FindAllStringSubmatch(s, -1)

		parts := strings.Split(s, " = ")
		distance, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

		city1 := matches[0][1]
		city2 := matches[0][2]
		if d, exists := distances[city1]; exists {
			d[city2] = distance
			distances[city1] = d
		} else {
			d := map[string]int{}
			d[city2] = distance
			distances[city1] = d
		}

		if d, exists := distances[city2]; exists {
			d[city1] = distance
			distances[city2] = d
		} else {
			d := map[string]int{}
			d[city1] = distance
			distances[city2] = d
		}

		locations[city1] = true
		locations[city2] = true
	}

	var memo = map[string]int{}
	ansP1 = travellingSalesperson("", locations, distances, memo, true)

	memo = map[string]int{}
	ansP2 = travellingSalesperson("", locations, distances, memo, false)

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
