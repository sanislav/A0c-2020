package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

type tile struct {
	x, y int
}

// Use Doubled coordinates
// https://www.redblobgames.com/grids/hexagons/
func (t tile) neighbours() []tile {
	return []tile{
		tile{t.x + 2, t.y}, 	// E
		tile{t.x - 2, t.y}, 	// W
		tile{t.x + 1, t.y + 1}, // NE
		tile{t.x + 1, t.y - 1}, // SE
		tile{t.x - 1, t.y + 1}, // NW
		tile{t.x - 1, t.y - 1}, // SW
	}
}

func setup(lines []string) map[tile]bool {
	blacks := map[tile]bool{}
	for _, line := range lines {
		var x, y int
		var part string

		// e, se, sw, w, nw, and ne
		for _, c := range line {
			switch string(c) {
			case "n":
				part = "n"
			case "s":
				part = "s"
			case "e":
				if part == "" {
					x += 2
				} else if part == "n" {
					x++
					y++
				} else {
					x++
					y--
				}
				part = ""

			case "w":
				if part == "" {
					x -= 2
				} else if part == "n" {
					x--
					y++
				} else {
					x--
					y--
				}
				part = ""
			}

		}
		t := tile{x, y}
		if _, exists := blacks[t]; exists {
			delete(blacks, t)
		} else {
			blacks[t] = true
		}
	}

	return blacks
}

// Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
// Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
func solveP2(blackTiles map[tile]bool) map[tile]bool {
	for i := 0; i < 100; i++ {
		newTiles := map[tile]bool{}

		surroundingTiles := []tile{}
		for t := range blackTiles {
			surroundingTiles = append(surroundingTiles, t.neighbours()...)
		}

		for _, t := range surroundingTiles {
			blackNeighboursCount := 0
			for _, neighbour := range t.neighbours() {
				if _, exists := blackTiles[neighbour]; exists {
					blackNeighboursCount++
				}
			}

			if _, exists := blackTiles[t]; exists {
				if blackNeighboursCount == 1 || blackNeighboursCount == 2 {
					newTiles[t] = true
				}
			} else {
				if blackNeighboursCount == 2 {
					newTiles[t] = true
				}
			}
		}

		blackTiles = newTiles
	}

	return blackTiles
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	blackTiles := setup(lines)
	fmt.Println("P1: ", len(blackTiles))

	blackTiles = solveP2(blackTiles)
	fmt.Println("P2: ", len(blackTiles))
}