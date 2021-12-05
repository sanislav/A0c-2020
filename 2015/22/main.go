package main

import "fmt"

// "strconv"
// "sort"

// Magic Missile costs 53 mana. It instantly does 4 damage.
// Drain costs 73 mana. It instantly does 2 damage and heals you for 2 hit points.
// Shield costs 113 mana. It starts an effect that lasts for 6 turns. While it is active, your armor is increased by 7.
// Poison costs 173 mana. It starts an effect that lasts for 6 turns. At the start of each turn while it is active, it deals the boss 3 damage.
// Recharge costs 229 mana. It starts an effect that lasts for 5 turns. At the start of each turn while it is active, it gives you 101 new mana.
type boss struct {
	hp int
	damage int
}

type player struct {
	hp int
	mana int
}

type abbilityProperties struct {
	name string
	manaCost int
	manaGain int
	damage int
	healing int
	increaseArmour int
	turnsActive int
}

func main() {
	playerWin := false
	bossDamage := 9

	for (!playerWin) {
		boss := boss{51, bossDamage}
		player := player{50, 500}

		playerOptions := []abbilityProperties{}

		playerOptions = append(playerOptions, abbilityProperties{name: "mm", manaCost: 53, damage: 4})
		playerOptions = append(playerOptions, abbilityProperties{name: "drain", manaCost: 73, damage: 2, healing: 2})
		playerOptions = append(playerOptions, abbilityProperties{name: "shield", manaCost: 113, increaseArmour: 2, turnsActive: 6})
		playerOptions = append(playerOptions, abbilityProperties{name: "poison", manaCost: 173, damage: 3, turnsActive: 6})
		playerOptions = append(playerOptions, abbilityProperties{name: "recharge", manaCost: 229, manaGain: 101, turnsActive: 5})

		tickers := []abbilityProperties{}
		nextTickers := []abbilityProperties{}
		solutions := map[string]bool{}

		solutionHash := ""

		for player.hp > 0 && boss.hp > 0 {
			// ticker effects
			for _, tick := range(tickers) {
				boss, player, tick = applySpell(boss, player, tick)
	
				if tick.turnsActive > 0 {
					nextTickers = append(nextTickers, tick)
				}
			}

			tickers = nextTickers

			// choose next ability
			castIndex := -1
			for i, ability := range(playerOptions) {
				if player.mana >= ability.manaCost &&  {
					castIndex = i
					break
				}
			}

			solutionHash += playerOptions[castIndex].name
			fmt.Println("Player casts: " + playerOptions[castIndex].name)

			if (castIndex >= 0) {
				player.hp = 0
				solutions[solutionHash] = false
			}

			if val, ok := solutions[solutionHash]; ok {
				if (val == false) {
					break
				}
			}

			if (playerOptions[castIndex].turnsActive == 0) {
				boss, player, _ = applySpell(boss, player, playerOptions[castIndex])
			}

			player.mana -= playerOptions[castIndex].manaCost
			player.hp -= boss.damage
			boss.damage = bossDamage

			if (boss.hp <= 0) {
				playerWin = true

				fmt.Println(player)
				fmt.Println(boss)

			}
			solutions[solutionHash] = true

			if (player.hp <= 0) {
				solutions[solutionHash] = false
			}
		}

		fmt.Println(solutions)

	}
}

// apply effects of spell on boss and player
func applySpell(boss boss, player player, spell abbilityProperties) (boss, player, abbilityProperties) {
	boss.hp -= spell.damage
	boss.damage -= spell.increaseArmour

	player.mana += spell.manaGain
	player.hp += spell.healing

	if (spell.turnsActive > 0) {
		spell.turnsActive--;
	}

	return boss, player, spell
}