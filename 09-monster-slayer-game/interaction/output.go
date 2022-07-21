package interaction

import "fmt"

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")

}

func ShowAvailableActions(specialAttackIsAvailable bool) {

	fmt.Println("Please choose your action")
	fmt.Println("--------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special attack")
	}

}

func PrintRoundStatistics(roundData *RoundData) {
	switch roundData.Action {
	case "ATTACK":
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	case "SPECIAL_ATTACK":
		fmt.Printf("Player performed a special attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	case "HEAL":
		fmt.Printf("Player healed for %v\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)

	fmt.Printf("Player health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster health: %v\n", roundData.MonsterHealth)

}

func DeclareWinner(winner string) {
	fmt.Println("---------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("---------------------")
	fmt.Printf("%v won!\n", winner)
}
