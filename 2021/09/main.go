package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

 const ml = 102
 const mc = 102

// const ml = 7
// const mc = 12

var basins = map[string]map[string]bool{}
var visited = map[string]bool{}
var toVisit = [][]int{}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0

	floor := [ml][mc]int{}

	for i := 0; i < len(floor); i++ {
		for j:= 0; j < len(floor[0]); j++ {
			if i == 0 || j == 0 || i == len(floor) - 1 || j == len(floor[0]) - 1 {
				floor[i][j] = 9
			} else {
				c := lines[i-1][j-1]
				intC, _ := strconv.Atoi(string(c))
				floor[i][j] = intC
			}
		}
	}

	basinCounts := []int{}
	for i := 1; i < len(floor) - 1; i++ {
		for j:= 1; j < len(floor[0]) -1; j++ {

			if isLowest(floor, i, j) {
				ansP1 += floor[i][j] + 1;
				toVisit = [][]int{}
				c := countBasin(floor, i, j, strconv.Itoa(i) + "_" + strconv.Itoa(j))
				basinCounts = append(basinCounts, c)
				// fmt.Println("Basin " + strconv.Itoa(i)  + "_" +  strconv.Itoa(j) + " has count " + strconv.Itoa(c))
			}
		}
	}

	sort.Ints(basinCounts)

	fmt.Println(ansP1)
	fmt.Println(basinCounts[len(basinCounts) - 1] * basinCounts[len(basinCounts) - 2] * basinCounts[len(basinCounts) - 3])
}


func isLowest(floor [ml][mc]int, i int, j int) bool {
	return floor[i][j] < floor[i-1][j] && floor[i][j] < floor[i+1][j] && floor[i][j] < floor[i][j-1] && floor[i][j] < floor[i][j+1]
}


func countBasin(floor [ml][mc]int, i int, j int, basin string) int {
	coords := getHash(i, j)

	r := addToBasin(basin, coords)
	if r {
		visited[coords] = true
	}

	if floor[i-1][j] != 9 {
		addToVistit(i-1, j)
	}
	if floor[i+1][j] != 9 {
		addToVistit(i+1, j)
	}
	if floor[i][j-1] != 9 {
		addToVistit(i, j-1)
	}
	if floor[i][j+1] != 9 {
		addToVistit(i, j+1)
	}

	if (len(toVisit) == 0) {
		return len(basins[basin])
	}

	// pop first el from toVisit and countBasin for it
	next := toVisit[0]
	toVisit = append(toVisit[:0], toVisit[1:]...)

	return countBasin(floor, next[0], next[1], basin)
}

func getHash(i int, j int) string {
	return strconv.Itoa(i)  + "_" +  strconv.Itoa(j)
}

func addToVistit(i int, j int) {
	coords := getHash(i, j)
	
	if _, vis := visited[coords]; ! vis {
		set := true
		for _, c := range toVisit {
			if c[0] == i && c[1] == j {
				set = false
				break
			}
		}
		if (set) {
			toVisit = append(toVisit, []int{i, j})
		}
	}
}

func addToBasin(basin string, coordsinates string) bool {
	m := map[string]bool{}

	if coords, ok := basins[basin]; !ok {
		m[coordsinates] = true
		basins[basin] = m
		return true
	} else {
		if _, ok := coords[coordsinates]; !ok {
			coords[coordsinates] = true
			basins[basin] = coords
			return true
		}
	}

	return false
}
