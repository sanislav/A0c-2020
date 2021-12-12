package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var routes = [][]string{}
var graph = map[string][]string{} // adjacency map

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	for _, l := range lines {
		nodes := strings.Split(l, "-")
		if nodes[1] != "start" {
			graph[nodes[0]] = append(graph[nodes[0]], nodes[1])
		}
		if nodes[0] != "start" {
			graph[nodes[1]] = append(graph[nodes[1]], nodes[0])
		}
	}

	ansP1 := countPaths("start", []string{}, true)

	ansP2 := countPaths("start", []string{}, false);
	fmt.Println(ansP1)
	fmt.Println(ansP2)
}

func countPaths(room string, small_visited []string, revisit bool) int {
    if room == "end" {
        return 1
    }

    if stringInSlice(room, small_visited) {
        if revisit || room == "start" {
            return 0
        } else {
            revisit = true
        }
    }

    if room == strings.ToLower(room) {
        add := true
        for _, v := range small_visited {
            if v == room {
                add = false
                break
            }
        }
        if add {
            small_visited = append(small_visited, room)
        }
    }

    sum := 0
    for _, neighbor := range graph[room] {
        sum += countPaths(neighbor, small_visited, revisit)
    }
    return sum
}


func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}