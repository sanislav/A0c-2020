package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Matrix struct {
	data [1000][1000]int
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	calculateAnswer(generateMatrix(lines, false))
	calculateAnswer(generateMatrix(lines, true))
}

func calculateAnswer(m Matrix) {
	ans := 0

	for i := 0; i < len(m.data[0]); i++ {
		for j := 0; j < len(m.data[i]); j++ {
			if m.data[i][j] > 1 {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

func generateMatrix(lines []string, p2 bool) Matrix {
	m := Matrix{}

	for _, line := range(lines) {
		parts := strings.Split(line, " -> ")

		point1 := strings.Split(parts[0], ",")
		point2 := strings.Split(parts[1], ",")

		x1, _ := strconv.Atoi(point1[0])
		y1, _ := strconv.Atoi(point1[1])

		x2, _ := strconv.Atoi(point2[0])
		y2, _ := strconv.Atoi(point2[1])
		sameLine := x1 == x2 || y1 == y2
		sameDiag := (x1 != x2 && y1 != y2) && math.Abs(float64(x1) - float64(x2)) == math.Abs(float64(y1) - float64(y2))

		if (sameLine) {
			if x1 < x2 {
				for i := x1; i <= x2; i ++ {
					m.data[y1][i] += 1
				}
			} else if x2 < x1 {
				for i := x2; i <= x1; i ++ {
					m.data[y1][i] += 1
				}
			}
			if y1 < y2 {
				for i := y1; i <= y2; i ++ {
					m.data[i][x1] += 1
				}
			} else if y2 < y1 {
				for i := y2; i <= y1; i ++ {
					m.data[i][x1] += 1
				}
			}
		}

		if (sameDiag && p2) {
			if (x1 < x2) {
				c := 0
				for i := x1; i <= x2; i ++ {
					if (y1 < y2) {
						m.data[y1 + c][i] += 1
					} else {
						m.data[y1 - c][i] += 1
					}
					c++
				}
			} else {
				c := 0
				for i := x1; i >= x2; i -- {
					if (y1 < y2) {
						m.data[y1 + c][i] += 1
					} else {
						m.data[y1 - c][i] += 1
					}
					c++
				}
			}
		}
	}

	return m
}
