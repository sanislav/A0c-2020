package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func minMax(is []int) (min, max int) {
	min, max = is[0], is[0]
	for _, i := range is {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return
}

func shuffle(input []int, nTurns int) map[int]*Node {
	min, max := minMax(input)

	nodeLookup := map[int]*Node{}
	for _, i := range input {
		nodeLookup[i] = &Node{value: i}
	}
	for idx, i := range input {
		nodeLookup[i].next = nodeLookup[input[(idx+1)%len(input)]]
	}

	current := nodeLookup[input[0]]
	for i := 0; i < nTurns; i++ {
		threeCupStart := current.next
		threeCupEnd := threeCupStart.next.next
		current.next = threeCupEnd.next

		destination := current.value - 1
		for {
			if destination < min {
				destination = max
			}

			if (destination != threeCupStart.value) &&
				(destination != threeCupStart.next.value) &&
				(destination != threeCupEnd.value) {
				break
			}

			destination--
		}
		destinationCup := nodeLookup[destination]

		t := destinationCup.next
		destinationCup.next = threeCupStart
		threeCupEnd.next = t

		current = current.next
	}

	return nodeLookup
}

func solveP1(input []int) string {
	nodeLookup := shuffle(input, 100)

	s := ""
	cup := nodeLookup[1].next
	for cup != nodeLookup[1] {
		s += fmt.Sprint(cup.value)
		cup = cup.next
	}
	return s
}

func solveP2(input []int) int {
	_, max := minMax(input)
	length := len(input)
	for i := 0; i < 1000000-length; i++ {
		input = append(input, max+1+i)
	}

	nodeLookup := shuffle(input, 10000000)

	return nodeLookup[1].next.value * nodeLookup[1].next.next.value
}

func main() {
	input := []int{2, 4, 7, 8, 1, 9, 3, 5, 6}
	fmt.Println("Part 1 =", solveP1(input))
	fmt.Println("Part 2 =", solveP2(input))
}