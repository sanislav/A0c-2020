package main

import (
	"strconv"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	sections := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	player1 := make([]int, 0)
	player2 := make([]int, 0)

	p1lines := strings.Split(sections[0], "\n")
	p2lines := strings.Split(sections[1], "\n")
	for _, l := range p1lines {
		if l == "Player 1:" {
			continue
		}

		card, _ := strconv.Atoi(l)
		player1 = append(player1, card)
	}

	for _, l := range p2lines {
		if l == "Player 2:" {
			continue
		}
		card, _ := strconv.Atoi(l)
		player2 = append(player2, card)
	}

    p1, deck := combat(player1, player2, false)
    score := 0
    for i, c := range deck {
        score += (len(deck)-i) * c
	}
	fmt.Println(score, p1)

	p1, deck = combat(player1, player2, true)
    score = 0
    for i, c := range deck {
        score += (len(deck)-i) * c
	}
	fmt.Println(score, p1)
}

func combat(player1 []int, player2 []int, isP2 bool) (bool, []int) {
	seen := make(map[string]bool, 0)
	for len(player1) != 0 && len(player2) != 0 {
		hash := ""
		for _, i := range player1 {
			hash += strconv.Itoa(i)
		}
		for _, j := range player2 {
			hash += strconv.Itoa(j)
		}

		_, hashed := seen[hash]
		if hashed && isP2 {
            return true, player1
		}
		seen[hash] = true

		c1 := player1[0]
		c2 := player2[0]
		p1Wins := false

		if isP2 && len(player1) > c1 && len(player2) > c2 {
			newDeck1 := make([]int, 0)
			newDeck2 := make([]int, 0)
			for i := 1; i <= c1; i++ {
				newDeck1 = append(newDeck1, player1[i])
			}
			for i := 1; i <= c2; i++ {
				newDeck2 = append(newDeck2, player2[i])
			}
			// fmt.Println(player1, c1, newDeck1)
			// fmt.Println(player2, c2, newDeck2)
			p1Wins, _ = combat(newDeck1, newDeck2, isP2)
		} else {
            p1Wins = c1 > c2
		}

        if p1Wins {
			player1 = append(player1[1:], c1, c2)
			player2 = player2[1:]
		} else {
			player1 = player1[1:]
			player2 = append(player2[1:], c2, c1)
		}
	}

    if len(player1) > len(player2) {
        return true, player1
	}

	return false, player2
}