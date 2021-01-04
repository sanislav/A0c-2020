package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	ingredients := []*Ingredient{}
	ansP1 := 0
	ansP2 := 0

	regex := regexp.MustCompile("(-?\\d+)")
	for _, l := range lines {
		matches := regex.FindAllStringSubmatch(l, -1)

		capacity, _ := strconv.Atoi(matches[0][0])
		durability, _ := strconv.Atoi(matches[1][0])
		flavor, _ := strconv.Atoi(matches[2][0])
		texture, _ := strconv.Atoi(matches[3][0])
		calories, _ := strconv.Atoi(matches[4][0])

		ingredient := &Ingredient{capacity, durability, flavor, texture, calories}
		ingredients = append(ingredients, ingredient)
	}

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100-a; b++ {
			for c := 0; c <= 100-a-b; c++ {
				d := 100 - a - b - c

				amounts := []int{a, b, c, d}
				var capacity, durability, flavor, texture, calories int

				for i, ing := range ingredients {
					amount := amounts[i]

					capacity += ing.capacity * amount
					durability += ing.durability * amount
					flavor += ing.flavor * amount
					texture += ing.texture * amount
					calories += ing.calories * amount
				}

				if capacity < 0 {
					capacity = 0
				}
				if durability < 0 {
					durability = 0
				}
				if flavor < 0 {
					flavor = 0
				}
				if texture < 0 {
					texture = 0
				}

				score := capacity * durability * flavor * texture

				if score > ansP1 {
					ansP1 = score
				}

				if calories == 500 && score > ansP2 {
					ansP2 = score
				}
			}
		}
	}

	println(ansP1)
	println(ansP2)
}
