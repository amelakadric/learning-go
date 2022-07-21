package main

import (
	"github.com/amelakadric/learning-go/09-monster-slayer-game/actions"
	"github.com/amelakadric/learning-go/09-monster-slayer-game/interaction"
)

var currentRound = 0

func main() {
	winner := ""

	startGame()

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)

}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := (currentRound%3 == 0)

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	var playerHealth int
	var monsterHealth int

	switch userChoice {
	case "ATTACK":
		actions.AttackMonster(false)
	case "HEAL":
		actions.HealPlayer()
	case "SPECIAL_ATTACK":
		actions.AttackMonster(true)
	}

	actions.AttackPlayer()
	playerHealth, monsterHealth = actions.GetHealthAmount()

	if playerHealth <= 0 {
		return "Monster"

	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
}
