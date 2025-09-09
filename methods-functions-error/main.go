package main

import (
	"fmt"
)

//
// 1. Create a struct Player with fields name, position, and stamina.
// Write a method Train that increases stamina.
//
type Player struct {
	name     string
	position string
	stamina  int
}

// pointer receiver so the stamina of the original player changes
func (p *Player) Train() {
	p.stamina = p.stamina + 10
}

//
// 2. Add a method PlayMatch that decreases stamina.
// If stamina goes below 0, return an error message.
//
func (p *Player) PlayMatch() string {
	p.stamina = p.stamina - 20
	if p.stamina < 0 {
		return "error: stamina cannot be negative"
	}
	return ""
}

//
// 3. Write a standalone function ComparePlayers that takes two Player pointers
// and returns the one with higher stamina. If stamina is equal, return an error.
//
func ComparePlayers(p1 *Player, p2 *Player) (*Player, string) {
	if p1.stamina > p2.stamina {
		return p1, ""
	} else if p2.stamina > p1.stamina {
		return p2, ""
	} else {
		return nil, "error: both players have equal stamina"
	}
}

//
// 4. Write a method CheckStamina that shows the current stamina of a player.
//
func (p *Player) CheckStamina() int {
	return p.stamina
}

//
// 5. Write a function that takes a slice of Player and prints only players with stamina greater than 50.
//
func PrintStrongPlayers(players []Player) {
	for _, player := range players {
		if player.stamina > 50 {
			fmt.Println("Player:", player.name, "Position:", player.position, "Stamina:", player.stamina)
		}
	}
}

func main() {
	// Q1
	player1 := Player{name: "John", position: "Midfielder", stamina: 40}
	fmt.Println("1. Before training:", player1.stamina)
	player1.Train()
	fmt.Println("1. After training:", player1.stamina)

	// Q2
	player2 := Player{name: "Mike", position: "Striker", stamina: 10}
	fmt.Println("2. Before match:", player2.stamina)
	err := player2.PlayMatch()
	if err != "" {
		fmt.Println("2. Error:", err)
	} else {
		fmt.Println("2. After match:", player2.stamina)
	}

	// Q3
	player3 := Player{name: "Sam", position: "Defender", stamina: 60}
	winner, err := ComparePlayers(&player1, &player3)
	if err != "" {
		fmt.Println("3. Error:", err)
	} else {
		fmt.Println("3. Player with more stamina:", winner.name, "-", winner.stamina)
	}

	// Q4
	fmt.Println("4. Checking stamina of", player1.name, ":", player1.CheckStamina())

	// Q5
	allPlayers := []Player{
		player1,
		player2,
		player3,
		{name: "Alex", position: "Goalkeeper", stamina: 55},
	}
	fmt.Println("5. Players with stamina > 50:")
	PrintStrongPlayers(allPlayers)
}
