package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Player struct with Name, Position, and Goals
2. Make a slice of Player called team
3. Write AddPlayer function to add new players to the team
4. Write UpdateGoals function to increase a player’s goals or return an error if not found
5. Use a map called teamStats to track total goals and players, and write ShowStats to print it
*/

type Player struct {
	Name     string
	Position string
	Goals    int
}

func AddPlayer(team []Player, player Player) []Player {
	team = append(team, player)
	return team
}

func UpdateGoals(team []Player, name string, goals int) ([]Player, error) {
	for i := 0; i < len(team); i++ {
		if team[i].Name == name {
			newGoalTotal := team[i].Goals + goals
			team[i].Goals = newGoalTotal
			return team, nil
		}
	}
	return team, errors.New("player not found")
}

func ShowStats(team []Player) {
	teamStats := make(map[string]int)
	totalGoals := 0

	for _, player := range team {
		totalGoals = totalGoals + player.Goals
	}

	teamStats["TotalGoals"] = totalGoals
	teamStats["PlayerCount"] = len(team)

	fmt.Println("Team Stats:", teamStats)
}

func main() {
	// Step 1: Create an empty slice for the team
	team := []Player{}

	// Step 2: Add players to the team
	team = AddPlayer(team, Player{Name: "Ronaldo", Position: "Forward", Goals: 5})
	team = AddPlayer(team, Player{Name: "Messi", Position: "Forward", Goals: 7})
	team = AddPlayer(team, Player{Name: "Modric", Position: "Midfielder", Goals: 3})

	// Step 3: Show team list
	fmt.Println("Initial Team:")
	for _, player := range team {
		fmt.Println("Name:", player.Name, "| Position:", player.Position, "| Goals:", player.Goals)
	}

	// Step 4: Update goals for a player
	var err error
	team, err = UpdateGoals(team, "Ronaldo", 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 5: Show updated team stats
	fmt.Println("\nAfter Updating Goals:")
	for _, player := range team {
		fmt.Println("Name:", player.Name, "| Position:", player.Position, "| Goals:", player.Goals)
	}
	ShowStats(team)
}