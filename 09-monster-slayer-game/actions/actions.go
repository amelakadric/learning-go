package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}
	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue

}

func HealPlayer() {

	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)
	currentPlayerHealth += healValue
	if currentPlayerHealth > PLAYER_HEALTH {
		currentPlayerHealth = PLAYER_HEALTH
	}

}

func AttackPlayer() {

	dmgValue := generateRandBetween(MONSTER_ATTACK_MIN_DMG, MONSTER_ATTACK_MAX_DMG)
	currentPlayerHealth -= dmgValue
}

func GetHealthAmount() (int, int) {
	return currentMonsterHealth, currentPlayerHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
