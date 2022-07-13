// Shelby Simpson
// Module 6 Activity: Goblin Tower
package main

import (
	"math/rand"
	"time"
)

type game struct {
	player hero
}

type hero struct {
	health  int
	attack  int
	defense int
	potions [5]int
	gold    int
	stepNum int
}

type goblin struct {
	health  int
	attack  int
	defense int
}

func (g *game) generateHero() {
	g.player = hero{
		health:  rand.Intn(11) + 20,
		attack:  rand.Intn(3) + 1,
		defense: rand.Intn(5) + 1,
		potions: [5]int{2, 2, 2, 2, 2},
		gold:    0,
	}
}

func fight(player *hero, goblin *goblin) bool {
	for player.health > 0 && goblin.health > 0 {
		playerAttack := player.attack
		playerDefense := player.defense
		goblinAttack := goblin.attack
		goblinDefense := goblin.defense
		for playerDefense > 0 && goblinAttack > 0 {
			playerDefense--
			goblinAttack--
		}
		if goblinAttack > 0 {
			for player.health > 0 && goblinAttack > 0 {
				player.health--
				goblinAttack--
			}
		}
		for goblinDefense > 0 && playerAttack > 0 {
			goblinDefense--
			playerAttack--
		}
		if playerAttack > 0 {
			for goblin.health > 0 && playerAttack > 0 {
				goblin.health--
				playerAttack--
			}
		}
	}
	if player.health > 0 {
		player.gold += 2
		return true // The hero won
	} else {
		return false // The hero lost
	}
}

func generateGoblin() goblin {
	var newGoblin = goblin{
		health:  rand.Intn(6) + 5,
		attack:  rand.Intn(2) + 2,
		defense: rand.Intn(2) + 1,
	}
	return newGoblin
}

func (h *hero) takeStep() {
	h.stepNum++
	var goblin goblin
	goblinChance := rand.Intn(2)
	if goblinChance == 1 {
		goblin = generateGoblin()
		winner := fight(h, &goblin)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
}
