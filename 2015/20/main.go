package main

import (
	"fmt"
)

func main() {
	target := 36000000
	ansP1 := 0
	ansP2 := 0
	house := [100000000]int{}

	for i := 1; i <= target / 10; i++ {
		for j := i; j < target / 10; j+=i {
			house[j] += i * 10
			if house[j] >= target && (ansP1 == 0 || j < ansP1) {
				ansP1 = j
			}
		}
	}
	fmt.Println(ansP1)

	house = [100000000]int{}

	for i := 1; i <= target / 10; i++ {
		count := 0
		for j := i; j <= target / 10; j+=i {
			house[j] += i * 11
			if house[j] >= target && (ansP2 == 0 || j < ansP2) {
				ansP2 = j
			}

			count++
			if count == 50 {
				break
			}
		}
	}
	fmt.Println(ansP2)
}