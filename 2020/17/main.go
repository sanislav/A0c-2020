package main

import (
	"fmt"
	"io/ioutil"
	// "strconv"
	// "regexp"
	"strings"
)

type D3 struct {
	X int
	Y int
	Z int
}

type D4 struct {
	X int
	Y int
	Z int
	W int
}

var D3Neighbors26 = [26]D3{
	{X: -1, Y: -1, Z: -1},
	{X: 0, Y: -1, Z: -1},
	{X: 1, Y: -1, Z: -1},
	{X: -1, Y: 0, Z: -1},
	{X: 0, Y: 0, Z: -1},
	{X: 1, Y: 0, Z: -1},
	{X: -1, Y: 1, Z: -1},
	{X: 0, Y: 1, Z: -1},
	{X: 1, Y: 1, Z: -1},
	{X: -1, Y: -1, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 1, Y: -1, Z: 0},
	{X: -1, Y: 0, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: -1, Y: 1, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 1, Y: 1, Z: 0},
	{X: -1, Y: -1, Z: 1},
	{X: 0, Y: -1, Z: 1},
	{X: 1, Y: -1, Z: 1},
	{X: -1, Y: 0, Z: 1},
	{X: 0, Y: 0, Z: 1},
	{X: 1, Y: 0, Z: 1},
	{X: -1, Y: 1, Z: 1},
	{X: 0, Y: 1, Z: 1},
	{X: 1, Y: 1, Z: 1},
}

var D4Neighbors80 = [80]D4{
	{X: -1, Y: -1, Z: -1, W: -1},
	{X: -1, Y: -1, Z: -1, W: 0},
	{X: -1, Y: -1, Z: -1, W: 1},
	{X: -1, Y: -1, Z: 0, W: -1},
	{X: -1, Y: -1, Z: 0, W: 0},
	{X: -1, Y: -1, Z: 0, W: 1},
	{X: -1, Y: -1, Z: 1, W: -1},
	{X: -1, Y: -1, Z: 1, W: 0},
	{X: -1, Y: -1, Z: 1, W: 1},
	{X: -1, Y: 0, Z: -1, W: -1},
	{X: -1, Y: 0, Z: -1, W: 0},
	{X: -1, Y: 0, Z: -1, W: 1},
	{X: -1, Y: 0, Z: 0, W: -1},
	{X: -1, Y: 0, Z: 0, W: 0},
	{X: -1, Y: 0, Z: 0, W: 1},
	{X: -1, Y: 0, Z: 1, W: -1},
	{X: -1, Y: 0, Z: 1, W: 0},
	{X: -1, Y: 0, Z: 1, W: 1},
	{X: -1, Y: 1, Z: -1, W: -1},
	{X: -1, Y: 1, Z: -1, W: 0},
	{X: -1, Y: 1, Z: -1, W: 1},
	{X: -1, Y: 1, Z: 0, W: -1},
	{X: -1, Y: 1, Z: 0, W: 0},
	{X: -1, Y: 1, Z: 0, W: 1},
	{X: -1, Y: 1, Z: 1, W: -1},
	{X: -1, Y: 1, Z: 1, W: 0},
	{X: -1, Y: 1, Z: 1, W: 1},
	{X: 0, Y: -1, Z: -1, W: -1},
	{X: 0, Y: -1, Z: -1, W: 0},
	{X: 0, Y: -1, Z: -1, W: 1},
	{X: 0, Y: -1, Z: 0, W: -1},
	{X: 0, Y: -1, Z: 0, W: 0},
	{X: 0, Y: -1, Z: 0, W: 1},
	{X: 0, Y: -1, Z: 1, W: -1},
	{X: 0, Y: -1, Z: 1, W: 0},
	{X: 0, Y: -1, Z: 1, W: 1},
	{X: 0, Y: 0, Z: -1, W: -1},
	{X: 0, Y: 0, Z: -1, W: 0},
	{X: 0, Y: 0, Z: -1, W: 1},
	{X: 0, Y: 0, Z: 0, W: -1},
	{X: 0, Y: 0, Z: 0, W: 1},
	{X: 0, Y: 0, Z: 1, W: -1},
	{X: 0, Y: 0, Z: 1, W: 0},
	{X: 0, Y: 0, Z: 1, W: 1},
	{X: 0, Y: 1, Z: -1, W: -1},
	{X: 0, Y: 1, Z: -1, W: 0},
	{X: 0, Y: 1, Z: -1, W: 1},
	{X: 0, Y: 1, Z: 0, W: -1},
	{X: 0, Y: 1, Z: 0, W: 0},
	{X: 0, Y: 1, Z: 0, W: 1},
	{X: 0, Y: 1, Z: 1, W: -1},
	{X: 0, Y: 1, Z: 1, W: 0},
	{X: 0, Y: 1, Z: 1, W: 1},
	{X: 1, Y: -1, Z: -1, W: -1},
	{X: 1, Y: -1, Z: -1, W: 0},
	{X: 1, Y: -1, Z: -1, W: 1},
	{X: 1, Y: -1, Z: 0, W: -1},
	{X: 1, Y: -1, Z: 0, W: 0},
	{X: 1, Y: -1, Z: 0, W: 1},
	{X: 1, Y: -1, Z: 1, W: -1},
	{X: 1, Y: -1, Z: 1, W: 0},
	{X: 1, Y: -1, Z: 1, W: 1},
	{X: 1, Y: 0, Z: -1, W: -1},
	{X: 1, Y: 0, Z: -1, W: 0},
	{X: 1, Y: 0, Z: -1, W: 1},
	{X: 1, Y: 0, Z: 0, W: -1},
	{X: 1, Y: 0, Z: 0, W: 0},
	{X: 1, Y: 0, Z: 0, W: 1},
	{X: 1, Y: 0, Z: 1, W: -1},
	{X: 1, Y: 0, Z: 1, W: 0},
	{X: 1, Y: 0, Z: 1, W: 1},
	{X: 1, Y: 1, Z: -1, W: -1},
	{X: 1, Y: 1, Z: -1, W: 0},
	{X: 1, Y: 1, Z: -1, W: 1},
	{X: 1, Y: 1, Z: 0, W: -1},
	{X: 1, Y: 1, Z: 0, W: 0},
	{X: 1, Y: 1, Z: 0, W: 1},
	{X: 1, Y: 1, Z: 1, W: -1},
	{X: 1, Y: 1, Z: 1, W: 0},
	{X: 1, Y: 1, Z: 1, W: 1},
}

func (c D3) Move(t D3) D3 {
	return D3{
		X: c.X + t.X,
		Y: c.Y + t.Y,
		Z: c.Z + t.Z,
	}
}

func (c D4) Move(t D4) D4 {
	return D4{
		X: c.X + t.X,
		Y: c.Y + t.Y,
		Z: c.Z + t.Z,
		W: c.W + t.W,
	}
}


type void = struct{}

type grid map[D3]void

func (g grid) ActiveNeighbors(pos D3) int {
	var active int

	for i := range D3Neighbors26 {
		if _, ok := g[pos.Move(D3Neighbors26[i])]; ok {
			active++
			if active > 3 {
				break
			}
		}
	}

	return active
}

func cube(input []string) int {
	active := make(grid)

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				active[D3{X: x, Y: y, Z: 0}] = void{}
			}
		}
	}

	for i := 0; i < 6; i++ {
		newActive := make(grid)
		visited := make(grid)

		for pos := range active {
			if n := active.ActiveNeighbors(pos); n == 2 || n == 3 {
				newActive[pos] = void{}
			}

			for y := range D3Neighbors26 {
				inactive := pos.Move(D3Neighbors26[y])

				if _, ok := visited[inactive]; ok {
					continue
				}

				visited[inactive] = void{}

				if _, ok := active[inactive]; !ok {
					if active.ActiveNeighbors(inactive) == 3 {
						newActive[inactive] = void{}
					}
				}
			}
		}

		active = newActive
	}

	return len(active)
}

type hyperCube map[D4]void

func (g hyperCube) ActiveNeighbors(pos D4) int {
	var active int

	for i := range D4Neighbors80 {
		if _, ok := g[pos.Move(D4Neighbors80[i])]; ok {
			active++
			if active > 3 {
				break
			}
		}
	}

	return active
}

func cube4d(input []string) int {
	active := make(hyperCube)

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				active[D4{X: x, Y: y, Z: 0, W: 0}] = void{}
			}
		}
	}

	for i := 0; i < 6; i++ {
		newActive := make(hyperCube)
		visited := make(hyperCube)

		for pos := range active {
			if n := active.ActiveNeighbors(pos); n == 2 || n == 3 {
				newActive[pos] = void{}
			}

			for y := range D4Neighbors80 {
				inactive := pos.Move(D4Neighbors80[y])

				if _, ok := visited[inactive]; ok {
					continue
				}

				visited[inactive] = void{}

				if _, ok := active[inactive]; !ok {
					if active.ActiveNeighbors(inactive) == 3 {
						newActive[inactive] = void{}
					}
				}
			}
		}

		active = newActive
	}

	return len(active)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Println(cube(lines))
	fmt.Println(cube4d(lines))
}

