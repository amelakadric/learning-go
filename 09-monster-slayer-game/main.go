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

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	switch userChoice {
	case "ATTACK":
		playerAttackDmg = actions.AttackMonster(false)
	case "HEAL":
		playerHealValue = actions.HealPlayer()
	case "SPECIAL_ATTACK":
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()
	playerHealth, monsterHealth := actions.GetHealthAmount()

	roundData := interaction.RoundData{
		Action:           userChoice,
		PlayerHealth:     playerHealth,
		MonsterHealth:    monsterHealth,
		PlayerAttackDmg:  playerAttackDmg,
		PlayerHealValue:  playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interaction.PrintRoundStatistics(&roundData)

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
