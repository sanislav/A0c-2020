package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0
	g := [][]string{}
	p1g := [][]string{}

	foldInstructions := []map[string]int{}
	dots := []map[int]int{}
	maxx := 0
	maxy := 0

	for _, line := range(lines) {
        if line == "" {
            continue
		} else if strings.Contains(line, "fold") {
			parts := strings.Split(line, " ")
			p := strings.Split(parts[2], "=")
			c, _ := strconv.Atoi(p[1])

			foldInstructions = append(foldInstructions, map[string]int{
				p[0]: c,
			})
        } else {
            parts := strings.Split(line, ",")
            y, _ := strconv.Atoi(parts[0])
            x, _ := strconv.Atoi(parts[1])

			if x > maxx {
				maxx = x
			}

			if y > maxy {
				maxy = y
			}

			dots = append(dots, map[int]int{x: y})
        }
	}

	for i := 0; i <= maxx; i++ {
		l := []string{}
		for j := 0; j <= maxy; j++ {
			l = append(l, ".")
		}

		g = append(g, l)
	}

	for _, m := range dots {
		for x, y := range m {
			g[x][y] = "#"
		}
	}

	c := 0
	for _, instruction := range foldInstructions {
		c++
		for d, v := range instruction {
			gf := [][]string{}

			if d == "y" {
				// flip horizontally
				for i := v + 1; i < len(g); i++ {
					l := []string{}
					for j := 0; j < len(g[0]); j++ {
						newV := "."

						if g[i][j] == "#" || g[v-(i-v)][j] == "#" {
							newV = "#"
						}

						l = append(l, newV)
					}
					gf = append(gf, l)
				}

				for m, n := 0, len(gf)-1; m < n; m, n = m+1, n-1 {
					gf[m], gf[n] = gf[n], gf[m]
				}
			} else {
				// flip vertically
				for i := 0; i < len(g); i++ {
					l := []string{}
					for j := v + 1; j < len(g[0]); j++ {
						newV := "."

						if g[i][j] == "#" || g[i][v-(j-v)] == "#" {
							newV = "#"
						}

						l = append(l, newV)
					}

					for m, n := 0, len(l)-1; m < n; m, n = m+1, n-1 {
						l[m], l[n] = l[n], l[m]
					}

					gf = append(gf, l)
				}
			}

			g = gf
		}

		if c == 1 {
			p1g = g
		}
		// break
	}

	for i := 0; i < len(p1g); i++ {
		for j := 0; j < len(p1g[0]); j++ {
			if p1g[i][j] == "#" {
				ansP1 +=1
			}
		}
	}
	fmt.Println(ansP1)

	for i := 0; i < len(g); i++ {
		fmt.Println(g[i])
	}
}
