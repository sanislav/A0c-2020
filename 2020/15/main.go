package main

import (
	"fmt"
)


func solveP1(input []int, turns int) int {
	mem := make(map[int][]int, 0)
	order := []int{15,12,0,14,3,1}

	turnNo := 0
	// last spoken
	num := order[len(order)-1]

	for turnNo < turns {
		num = order[len(order)-1]
		if (turnNo < len(input)) {
			num = input[turnNo]
		} else {
			if val, ok := mem[num]; ok {
				if (len(mem[num]) < 2) {
					num = 0
				} else {
					num = val[len(val) - 1] - val[len(val) - 2]
				}
			} else {
				num = 0
			}
		}

		mem[num] = append(mem[num], turnNo)

		order = append(order, num)

		turnNo++
	}

	return num
}

func main() {
	input := []int{15,12,0,14,3,1}

	ans := solveP1(input, 2020)
	fmt.Println("P1", ans)

	ans = solveP1(input, 30000000)
	fmt.Println("P1", ans)
}
