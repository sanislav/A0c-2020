package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func solve(inputString []string, withWaypoint bool) int {
	// N:1 E:2 S:3 W:4
	shipDirection := 2
	horizontal := 0
	vertical := 0
	waypointHorizontal := 0
	waypointVertical := 0

	if (withWaypoint) {
		waypointHorizontal = 10
		waypointVertical = 1
	}

	for _, line := range(inputString) {
		direction := string(line[0])
		distance, _ := strconv.Atoi(line[1:])
		rotation := 1

		if direction == "L" {
			rotation = -1
		}

		if direction == "R" || direction == "L" {
			if (withWaypoint) {

				// 10 1 -> R

				// #1 1 -10
				// #2 -10 -1
				// #3 -1 10

				// 10 1 -> L

				// #1 -1 10
				// #2 -10 -1
				// #3 1 -10
				if (distance / 90 == 2) {
					waypointHorizontal = -1 * waypointHorizontal
					waypointVertical = -1 * waypointVertical
				} else if (distance / 90 == 1 && rotation == 1 || distance / 90 == 3 && rotation == -1) {
					origHorizontal := waypointHorizontal
					waypointHorizontal = waypointVertical
					waypointVertical = -1 * origHorizontal
				} else if (distance / 90 == 3 && rotation == 1 || distance / 90 == 1 && rotation == -1) {
					origHorizontal := waypointHorizontal
					waypointHorizontal = -1 * waypointVertical
					waypointVertical = origHorizontal
				}
			} else {
				shipDirection += (distance / 90) * rotation
				if (shipDirection > 4) {
					shipDirection -= 4
				} else if (shipDirection < 1) {
					shipDirection += 4
				}
			}

			continue
		}

		if (direction == "F") {
			// move ship
			if (withWaypoint) {
				horizontal += distance * waypointHorizontal
				vertical += distance * waypointVertical
			} else {
				if shipDirection == 2 {
					horizontal += distance
				} else if shipDirection == 4 {
					horizontal -= distance
				} else if shipDirection == 1 {
					vertical += distance
				} else if shipDirection == 3 {
					vertical -= distance
				}
			}
		} else {
			if (withWaypoint) {
				// move waypoint
				if direction == "E" {
					waypointHorizontal += distance
				} else if direction == "W" {
					waypointHorizontal -= distance
				} else if direction == "N" {
					waypointVertical += distance
				} else if direction == "S" {
					waypointVertical -= distance
				}
			} else {
				if direction == "E" {
					horizontal += distance
				} else if direction == "W" {
					horizontal -= distance
				} else if direction == "N" {
					vertical += distance
				} else if direction == "S" {
					vertical -= distance
				}
			}
		}
	}

	if horizontal < 0 {
		horizontal *= -1
	}

	if vertical < 0 {
		vertical *= -1
	}

	return horizontal + vertical
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	inputString := strings.Split(strings.TrimSpace(string(input)), "\n")

	distance := solve(inputString, false)
	fmt.Println("P1", distance)

	distance = solve(inputString, true)
	fmt.Println("P2", distance)
}
