package main

import "github.com/amelakadric/learning-go/09-monster-slayer-game/interaction"

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

	return ""
}

func endGame() {}
