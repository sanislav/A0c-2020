package main

import (
	"fmt"
	"math"
)

type item struct {
	cost int
	damage int
	armor int
}

func main() {
	weapons := []item{
		{8, 4, 0},
		{10, 5, 0},
		{25, 6, 0},
		{40, 7, 0},
		{74, 8, 0},
	}

	armours := []item{
		{0, 0, 0},
		{13, 0, 1},
		{31, 0, 2},
		{53, 0, 3},
		{75, 0, 4},
		{102, 0, 5},
	}

	rings := []item{
		{0, 0, 0},
		{0, 0, 0},
		{25, 1, 0},
		{50, 2, 0},
		{100, 3, 0},
		{20, 0, 1},
		{40, 0, 2},
		{80, 0, 3},
	}

	myHP := 100
	ansP1 := 999
	ansP2 := 0

	for _, w := range weapons {
		myDamage := w.damage
		myArmour := 0
		cost := w.cost

		ansP1, ansP2 = setAnswers(myDamage, myArmour, myHP, cost, ansP1, ansP2)

		for _, a := range armours {
			myArmour = a.armor
			cost = w.cost + a.cost

			ansP1, ansP2 = setAnswers(myDamage, myArmour, myHP, cost, ansP1, ansP2)

			for i := 0; i < len(rings) - 1; i++ {
				costR := cost + rings[i].cost
				myArmourR := myArmour + rings[i].armor
				myDamageR := myDamage + rings[i].damage

				ansP1, ansP2 = setAnswers(myDamageR, myArmourR, myHP, costR, ansP1, ansP2)

				for j := i + 1; j < len(rings); j++ {
					ansP1, ansP2 = setAnswers(myDamageR + rings[j].damage, myArmourR + rings[j].armor, myHP, costR + rings[j].cost, ansP1, ansP2)
				}
			}
		}
	}

	fmt.Println(ansP1)
	fmt.Println(ansP2)
}

func setAnswers(myDamage int, myArmour int, myHP int, cost int, ansP1 int, ansP2 int) (int, int) {
	canWin := canWin(myDamage, myArmour, myHP)

	if canWin && cost < ansP1 {
		ansP1 = cost
	}

	if !canWin && cost > ansP2 {
		ansP2 = cost
	}

	return ansP1, ansP2
}

// bossHP/(myDammage - bossArmour) <= myHP/(bossDamage - myArmour)
func canWin(dmg int, arm int, hp int) bool {
	bossHP := 104
	bossDamage := 8
	bossArmour := 1

	myD := dmg - bossArmour
	if myD <= 0 {
		myD = 1
	}

	bossD := bossDamage - arm
	if bossD <= 0 {
		bossD = 1
	}

	return math.Ceil(float64(bossHP)/float64(myD)) <= math.Ceil(float64(hp)/float64(bossD))
}
