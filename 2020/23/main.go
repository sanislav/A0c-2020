package main

import (
	"strconv"
	"fmt"
	// "io/ioutil"
	// "strings"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) *Node {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return l.head
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return &n
		}
		ptr = ptr.next
	}

	return &n
}

func (l *LinkedList) Search(val int) *Node {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return ptr
		}
		ptr = ptr.next
	}

	return ptr
}

// The crab picks up the three cups that are immediately clockwise of the current cup. They are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
// The crab selects a destination cup: the cup with a label equal to the current cup's label minus one. If this would select one of the cups that was just picked up, the crab will keep subtracting one until it finds a cup that wasn't just picked up. If at any point in this process the value goes below the lowest value on any cup's label, it wraps around to the highest value on any cup's label instead.
// The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.
// The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
func shuffle(l LinkedList, count int) LinkedList {
	current := l.head

	for i:= 0; i < count; i++ {
		extractedMap := make(map[int]bool, 0)
		var extracted []*Node

		extractedMap[current.next.value] = true
		extractedMap[current.next.next.value] = true
		extractedMap[current.next.next.next.value] = true
		extracted = append(extracted, current.next, current.next.next, current.next.next.next)

		destination := current.value - 1
		if destination == 0 {
			destination = 9
		}
		_, justExtracted := extractedMap[destination]

		for justExtracted {
			destination--
			if destination == 0 {
				destination = 9
			}

			_, justExtracted = extractedMap[destination]
		}

		destinationElement := l.Search(destination)
		current.next = extracted[2].next
		extracted[2].next = destinationElement.next
		destinationElement.next = extracted[0]

		current = current.next
	}

	return l
}

func main() {
	input := []int{2, 4, 7, 8, 1, 9, 3, 5, 6}

	list := LinkedList{}
	listP2 := LinkedList{}

	for i, v := range input {
		node := list.Insert(v)
		_ = listP2.Insert(v)

		if i == len(input) - 1 {
			node.next = list.head

		}
	}

	for i:= 10; i <= 1000000; i++ {
		node := listP2.Insert(i)

		if i == 1000000 {
			node.next = listP2.head
		}
	}

	fmt.Println(listP2.len)
	l := shuffle(list, 100)

	ptr := l.head
	ansP1 := ""
	start := false
	finish := false
	for ! finish {
		if len(ansP1) == l.len - 1 {
			finish = true
			break
		}

		if (ptr.value == 1) {
			start = true
		} else {
			if start {
				ansP1 += strconv.Itoa(ptr.value)
			}
		}

		ptr = ptr.next
	}

	fmt.Println(ansP1)

	l2 := shuffle(listP2, 10000000)
	ptr = l2.head
	ansP2 := 1

	for {
		if (ptr.value == 1) {
			ansP2 = ptr.next.value * ptr.next.next.value
			break
		}

		ptr = ptr.next
	}

	fmt.Println(ansP2)
}
