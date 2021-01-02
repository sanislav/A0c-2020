package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)


// DP formula: g(start, {nodes...}) = min(Cost[start][k] + g(k, {nodes...}-{k})) for each k in {nodes}
func travellingSalesperson(start string,  persons map[string]bool, distances map[string]map[string]int, memo map[string]int, isMin bool) int {
	key := start
	for k := range persons {
		key += "_" + k
	}
	if v, exists := memo[key]; exists {
        return v
	}

	// base case - we reached a leaf - no need to go back to start
	if len(persons) == 0 {
		return distances[start]["Alice"] + distances["Alice"][start]
	}

    values := []int{}

    for k := range persons {
		subList := map[string]bool{}

		for k2 := range persons {
			if k != k2 {
				subList[k2] = true
			}
		}

        result := travellingSalesperson(k, subList, distances, memo, isMin)
        values = append(values, distances[start][k] + distances[k][start] + result)
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
	persons := map[string]bool{}
	start := ""
	for _, s := range lines {
		parts := strings.Split(s, " ")

		distance, _ := strconv.Atoi(strings.TrimSpace(parts[3]))

		if parts[2] == "lose" {
			distance = 0 - distance
		}

		person1 := parts[0]
		person2 := strings.TrimRight(parts[len(parts) - 1], ".")

		if start == "" {
			start = person1
		}

		if d, exists := distances[person1]; exists {
			d[person2] = distance
			distances[person1] = d
		} else {
			d := map[string]int{}
			d[person2] = distance
			distances[person1] = d
		}

		persons[person1] = true
		persons[person2] = true
	}

	var memo = map[string]int{}
	delete(persons, start)

	ansP1 = travellingSalesperson(start, persons, distances, memo, false)
	fmt.Println(ansP1)

	persons["Sabin"] = true
	ansP2 = travellingSalesperson(start, persons, distances, memo, false)
	fmt.Println(ansP2)
}
