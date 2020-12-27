package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
	"regexp"
)

type Pixels []string

func (p Pixels) Rotations() []Pixels {
	rot := make([]Pixels, 8)
	for r := 0; r < 8; r += 2 {
		for _, s := range p {
			rot[r] = append(rot[r], flip(s))
		}
		for i := range p {
			rot[r+1] = append(rot[r+1], flip(p.col(i)))
		}
		p = rot[r+1]
	}
	return rot
}

func (p Pixels) col(i int) (col string) {
	for _, s := range p {
		col += string(s[i])
	}
	return
}

func flip(s string) string {
	r := ""
	for _, c := range s {
		r = string(c) + r
	}

	return r
}

func neighboursMap(sensors map[string][]string) map[string][]string {
	m := map[string][]string{}
	for i1, e1 := range sensors {
		for i2, e2 := range sensors {
			if i1 != i2 {
				for _, v1 := range e1 {
					for _, v2 := range e2 {
						if v1 == v2 {
							alreadyExists := false
							for _, v := range(m[i1]) {
								if v == i2 {
									alreadyExists = true
									break
								}
							}
							if ! alreadyExists {
								m[i1] = append(m[i1], i2)
							}
						}
					}
				}
			}
		}
	}

	return m
}

func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)
	unique := make(map[string]bool)

	if (len(a) == 0) {
		return b
	}
	if (len(b) == 0) {
		return a
	}

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			if _, isUnique := unique[item]; ! isUnique {
				c = append(c, item)
				unique[item] = true
			}
		}
	}

	return
}

func rebuildImage(m map[string][]string) [12][12]string {
	start := ""
	for k, n := range m {
		if len(n) == 2 {
			start = k
			break
		}
	}

	grid := [12][12]string{}
	used := map[string]bool{}
	l1 := m[start][0]
	l2 := m[start][1]
	used[start] = true
	used[l1] = true
	used[l2] = true

	grid[0][0] = start
	grid[0][1] = l1
	grid[1][0] = l2

	for len(used) != 12 * 12 {
		for i := 0; i < 12; i++ {
			for j := 0; j < 12; j ++ {
				if grid[i][j] != "" {
					continue
				}

				possibilitiesLeft := []string{}
				possibilitiesTop := []string{}

				if j > 0 {
					left := grid[i][j-1]

					for _, v := range m[left] {
						if _, alreadyUsed := used[v]; !alreadyUsed {
							possibilitiesLeft = append(possibilitiesLeft, v)
						}
					}
				}
				if i > 0 {
					top := grid[i-1][j]
					for _, v := range m[top] {
						if _, alreadyUsed := used[v]; !alreadyUsed {
							for _, v := range m[top] {
								if _, alreadyUsed := used[v]; !alreadyUsed {
									possibilitiesTop = append(possibilitiesTop, v)
								}
							}
						}
					}
				}
				options := intersection(possibilitiesLeft, possibilitiesTop)

				if len(options) == 1 {
					grid[i][j] = options[0]
					used[options[0]] = true
				}

				if i == 11 && j == 11 {
					i = 0;
					j = 0;
				}
			}
		}
	}

	return grid
}

func generateImgPixels(grid [12][12]string, tileToPixels map[string]Pixels) Pixels {
	image := [120]string{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			sensorId := grid[i][j]

			downId, upId, rightId, leftId := "", "", "", ""
			downR, upR, rightR, leftR := []Pixels{}, []Pixels{}, []Pixels{}, []Pixels{}

			// flip / rotate to match neighbours below and right
			rotations := tileToPixels[sensorId].Rotations()

			if i > 0 {
				upId = grid[i-1][j]
				upR = tileToPixels[upId].Rotations()
			}
			if i < len(grid) - 1 {
				downId = grid[i+1][j]
				downR = tileToPixels[downId].Rotations()
			}
			if j > 0 {
				leftId = grid[i][j-1]
				leftR = tileToPixels[leftId].Rotations()
			}
			if j < len(grid[0]) - 1 {
				rightId = grid[i][j+1]
				rightR = tileToPixels[rightId].Rotations()
			}

			// check rotations matching right neighbour rotations
			position := ""
			for index, r := range rotations {
				matchesL := []string{}
				matchesR := []string{}
				matchesT := []string{}
				matchesB := []string{}

				strIndex := strconv.Itoa(index)

				edgeL := r.col(0)
				edgeR := r.col(len(r) - 1)
				edgeT := r[0]
				edgeB := r[len(r) - 1]

				if len(rightR) > 0 {
					for _, rr := range rightR {
						edgeLR := rr.col(0)

						if (edgeR == edgeLR) {
							matchesL = append(matchesL, strIndex)
						}
					}
				}

				if len(leftR) > 0 {
					for _, lr := range leftR {
						edgeRR := lr.col(len(lr) - 1)

						if (edgeL == edgeRR) {
							matchesR = append(matchesR, strIndex)
						}
					}
				}

				if len(upR) > 0 {
					for _, ur := range upR {
						edgeUU := ur[0]

						if (edgeT == edgeUU) {
							matchesT = append(matchesT, strIndex)

						}
					}
				}

				if len(downR) > 0 {
					for _, dr := range downR {
						edgeRD := dr[len(dr) - 1]

						if (edgeB == edgeRD) {
							matchesB = append(matchesB, strIndex)
						}
					}
				}

				// fmt.Println(matchesB, matchesT, matchesL, matchesR)
				if len(matchesB) == 1 {
					position = matchesB[0]
				} else if len(matchesT) == 1 {
					position = matchesT[0]
				} else if len(matchesL) == 1 {
					position = matchesL[0]
				} else if len(matchesR) == 1 {
					position = matchesR[0]
				}
			}

			posIndex, _ := strconv.Atoi(position)

			if position == "" {
				os.Exit(1)
			}
			sensorPixels := rotations[posIndex]
			for ir := 0; ir < len(sensorPixels); ir++ {
				image[(i*10) + ir] += string(sensorPixels[ir][1:len(sensorPixels[ir]) - 1])
			}
		}
	}

	cutImg := Pixels{}
	for i := 0; i < len(image); i++ {
		if (i % 10) != 0 && (i % 9) != 0 {
			cutImg = append(cutImg, image[i])
		}
	}

	return cutImg
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	tilesToPixels := map[string]Pixels{}
	sensors := map[string][]string{}
	for _, s := range sections {
		parts := strings.Split(s, "\n")
		id := strings.Split(strings.TrimRight(parts[0], ":"), "Tile ")

		top := string(parts[1])
		bottom := string(parts[10])
		left := ""
		right := ""
		pixels := Pixels{}

		for i := 1; i <= 10; i++ {
			left += string(parts[i][0])
			right += string(parts[i][9])
			pixels = append(pixels, parts[i])
		}

		sensors[id[1]] = []string{top, flip(top), bottom, flip(bottom), left, flip(left), right, flip (right)}
		tilesToPixels[id[1]] = pixels
	}

	m := neighboursMap(sensors)

	ans := 1
	for id, l := range m {
		if len(l) == 2 { // each corner piece only has 2 neighbours
			intID, _ := strconv.Atoi(id)
			ans *= intID
		}
	}

	fmt.Println(ans)

	sensorPositions := rebuildImage(m)
	img := generateImgPixels(sensorPositions, tilesToPixels)

	monster := []string{"..................#.", "#....##....##....###", ".#..#..#..#..#..#..."}
	nmonster := 0
	for _, r := range img.Rotations() {
		for y := 0; y < len(r)-len(monster); y++ {
		findMonster:
			for x := 0; x < len(r[0])-len(monster[0]); x++ {
				for i, s := range monster {
					if match, _ := regexp.MatchString(s, r[y+i][x:x+len(s)]); !match {
						continue findMonster
					}
				}
				nmonster++
			}
		}
	}
	fmt.Println(strings.Count(strings.Join(img, ""), "#") - nmonster*strings.Count(strings.Join(monster, ""), "#"))
}