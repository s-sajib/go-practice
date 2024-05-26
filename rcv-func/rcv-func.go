//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:

package main

import "fmt"

//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name

type Player struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

func (player *Player) modifyHealth(health int) {
	player.health += health
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
}

func (player *Player) modifyEnergy(energy int) {
	player.energy += energy
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	}
}

func main() {

	myPlayer := Player{
		name:      "Sajib",
		health:    80,
		maxHealth: 100,
		energy:    1000,
		maxEnergy: 1500,
	}

	fmt.Println("Initial Stat: ", myPlayer)

	myPlayer.modifyHealth(10)

	fmt.Println("After eating: ", myPlayer)

	myPlayer.modifyEnergy(1500)

	fmt.Println("After Motivation Video: ", myPlayer)
}
