package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Raindeer struct {
	Name string
	Speed, Duration, Rest int
	CalcDistance int
}

func (r Raindeer) distanceAfter(sec int) int {
	intervalDuration := r.Duration + r.Rest
	intervalDistance := r.Duration * r.Speed
	part :=	sec / intervalDuration
	ans := part * intervalDistance
	rest := sec - part * intervalDuration

	if rest > 0 {
		if rest < r.Duration {
			ans += rest * r.Speed
		} else {
			ans += r.Duration * r.Speed
		}
	}

	return ans
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	ansP1 := 0
	ansP2 := 0
	deers := []Raindeer{}

	for _, s := range lines {
		parts := strings.Split(s, " ")
		speed, _ := strconv.Atoi(parts[3])
		duration, _ := strconv.Atoi(parts[6])
		rest, _ := strconv.Atoi(parts[13])
		raindeer := Raindeer{
			Name: parts[0],
			Speed: speed,
			Duration: duration,
			Rest: rest,
		}
		deers = append(deers, raindeer)
		travel := raindeer.distanceAfter(2503)

		if travel > ansP1 {
			ansP1 = travel
		}
	}

	bonusPoints := map[string]int{}

	for s := 1; s <= 2503; s++ {
		bestDeers := []Raindeer{}
		best := 0

		for _, deer := range deers {
			distance := deer.distanceAfter(s)
			if distance >= best {
				best = distance
				deer.CalcDistance = distance
				bestDeers = append(bestDeers, deer)
			}
		}

		for _, deer := range bestDeers {
			if deer.CalcDistance == best {

				bonusPoints[deer.Name]++

				if ansP2 < bonusPoints[deer.Name] {
					ansP2 = bonusPoints[deer.Name]
				}
			}
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}
