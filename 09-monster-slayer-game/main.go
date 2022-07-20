package main

import "github.com/amelakadric/learning-go/09-monster-slayer-game/interaction"

var currentRound = 0

func main() {
	winner := ""

	startGame()

	for winner == "" {
		winner = executeRound()
	}

	endGame()

}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := (currentRound%3 == 0)

	interaction.ShowAvailableActions(isSpecialRound)

	return ""
}

func endGame() {}
