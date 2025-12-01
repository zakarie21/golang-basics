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
6. Write a function GetTopScorer(team []Player) Player that returns the player with the most goals.
7. Write a function FindPlayersByPosition that returns all players who play in a given position.
8. Write a function RemovePlayer(team []Player, name string) ([]Player, error) that removes a player by name or returns an error if not found.
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
			team[i].Goals += goals
			return team, nil
		}
	}
	return team, errors.New("player not found")
}

func ShowStats(team []Player) {
	teamStats := make(map[string]int)
	totalGoals := 0

	for _, player := range team {
		totalGoals += player.Goals
	}

	teamStats["TotalGoals"] = totalGoals
	teamStats["PlayerCount"] = len(team)

	fmt.Println("Team Stats:", teamStats)
}

func GetTopScorer(team []Player) Player {
	if len(team) == 0 {
		return Player{}
	}

	top := team[0]
	for _, p := range team {
		if p.Goals > top.Goals {
			top = p
		}
	}
	return top
}

func FindPlayersByPosition(team []Player, position string) []Player {
	var result []Player
	for _, p := range team {
		if p.Position == position {
			result = append(result, p)
		}
	}
	return result
}


func RemovePlayer(team []Player, name string) ([]Player, error) {
	for i := 0; i < len(team); i++ {
		if team[i].Name == name {
			team = append(team[:i], team[i+1:]...)
			return team, nil
		}
	}
	return team, errors.New("player not found")
}

func main() {
	team := []Player{}

	team = AddPlayer(team, Player{Name: "Ronaldo", Position: "Forward", Goals: 5})
	team = AddPlayer(team, Player{Name: "Messi", Position: "Forward", Goals: 7})
	team = AddPlayer(team, Player{Name: "Modric", Position: "Midfielder", Goals: 3})

	fmt.Println("Initial Team:")
	for _, player := range team {
		fmt.Println("Name:", player.Name, "| Position:", player.Position, "| Goals:", player.Goals)
	}

	var err error
	team, err = UpdateGoals(team, "Ronaldo", 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nAfter Updating Goals:")
	for _, player := range team {
		fmt.Println("Name:", player.Name, "| Position:", player.Position, "| Goals:", player.Goals)
	}
	ShowStats(team)

	
	fmt.Println("\nRemoving Player: Modric")
	team, err = RemovePlayer(team, "Modric")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nTeam After Removal:")
	for _, player := range team {
		fmt.Println("Name:", player.Name, "| Position:", player.Position, "| Goals:", player.Goals)
	}
}
