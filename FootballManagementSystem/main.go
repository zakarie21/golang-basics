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
8. Write a function RemovePlayer that removes a player by name from the team.
9. Write a function GetTotalGoalsByPosition(team []Player, position string) int that returns the total goals scored by players in a given position.
10. Write a function CountPlayers(team []Player) int that returns how many players are in the team.
11. AverageGoalsPerPlayer: return the average goals scored per player
12. GetPlayerByName: return a player by name or an error if not found
13. HasPlayer: return true if a player exists on the team
14. GetPlayersWithMinGoals: return players with at least N goals
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

	for _, player := range team {
		if player.Position == position {
			result = append(result, player)
		}
	}
	return result
}

func RemovePlayer(team []Player, name string) []Player {
	for i, player := range team {
		if player.Name == name {
			return append(team[:i], team[i+1:]...)
		}
	}
	return team
}

func GetTotalGoalsByPosition(team []Player, position string) int {
	total := 0
	for _, player := range team {
		if player.Position == position {
			total += player.Goals
		}
	}
	return total
}

func CountPlayers(team []Player) int {
	return len(team)
}

// Question 11
func AverageGoalsPerPlayer(team []Player) float64 {
	if len(team) == 0 {
		return 0
	}

	totalGoals := 0
	for _, player := range team {
		totalGoals += player.Goals
	}
	return float64(totalGoals) / float64(len(team))
}

// Question 12
func GetPlayerByName(team []Player, name string) (*Player, error) {
	for i := range team {
		if team[i].Name == name {
			return &team[i], nil
		}
	}
	return nil, errors.New("player not found")
}

// Question 13
func HasPlayer(team []Player, name string) bool {
	for _, player := range team {
		if player.Name == name {
			return true
		}
	}
	return false
}

// Question 14
func GetPlayersWithMinGoals(team []Player, minGoals int) []Player {
	var result []Player
	for _, player := range team {
		if player.Goals >= minGoals {
			result = append(result, player)
		}
	}
	return result
}

func main() {
	team := []Player{}

	team = AddPlayer(team, Player{Name: "Ronaldo", Position: "Forward", Goals: 5})
	team = AddPlayer(team, Player{Name: "Messi", Position: "Forward", Goals: 7})
	team = AddPlayer(team, Player{Name: "Modric", Position: "Midfielder", Goals: 3})
	team = AddPlayer(team, Player{Name: "Ramos", Position: "Defender", Goals: 1})

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

	fmt.Println("\nTop Scorer:", GetTopScorer(team))

	fmt.Println("\nMidfielders:", FindPlayersByPosition(team, "Midfielder"))

	team = RemovePlayer(team, "Ramos")
	fmt.Println("\nAfter Removing Ramos:")
	for _, player := range team {
		fmt.Println(player)
	}

	fmt.Println("\nTotal goals by forwards:", GetTotalGoalsByPosition(team, "Forward"))

	fmt.Println("\nTotal players in team:", CountPlayers(team))

	// Demo Question 11
	fmt.Printf("\nAverage goals per player: %.2f\n", AverageGoalsPerPlayer(team))

	// Demo Question 12
	player, err := GetPlayerByName(team, "Messi")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nRetrieved player:", player.Name, "-", player.Position, "- Goals:", player.Goals)
	}

	// Demo Question 13
	fmt.Println("\nDoes Ronaldo exist?", HasPlayer(team, "Ronaldo"))
	fmt.Println("Does Neymar exist?", HasPlayer(team, "Neymar"))

	// Demo Question 14
	fmt.Println("\nPlayers with at least 5 goals:")
	for _, p := range GetPlayersWithMinGoals(team, 5) {
		fmt.Println(p.Name, "-", p.Goals)
	}
}
