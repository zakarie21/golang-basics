package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Spaceship struct with fields Name, Crew, and Launched
2. Make a slice called fleet to store multiple spaceships
3. Write a method Launch() that sets Launched to true or returns an error if already launched
4. Write a method Land() that sets Launched to false
5. Write a function ShowFleet(fleet []Spaceship) that loops and prints all ships' details
6. CountLaunched: return how many ships in the fleet are launched
7. FindShipByName: return a pointer to a ship by name or an error if not found
8. ListUnlaunchedShips: return all ships that have not been launched
9. GetTotalCrew: return the total crew count across all ships
10. GetLargestCrewShip: return the ship with the highest crew count
*/

type Spaceship struct {
	Name     string
	Crew     int
	Launched bool
}

func (s *Spaceship) Launch() error {
	if s.Launched {
		return errors.New(s.Name + " has already been launched")
	}
	s.Launched = true
	return nil
}

func (s *Spaceship) Land() {
	s.Launched = false
}

func ShowFleet(fleet []Spaceship) {
	for _, ship := range fleet {
		status := "on standby"
		if ship.Launched {
			status = "launched"
		}
		fmt.Println(ship.Name, "with", ship.Crew, "crew members is currently", status)
	}
}

func CountLaunched(fleet []Spaceship) int {
	total := 0
	for _, ship := range fleet {
		if ship.Launched {
			total++
		}
	}
	return total
}

// Question 7
func FindShipByName(fleet []Spaceship, name string) (*Spaceship, error) {
	for i := range fleet {
		if fleet[i].Name == name {
			return &fleet[i], nil
		}
	}
	return nil, errors.New("ship not found: " + name)
}

// Question 8
func ListUnlaunchedShips(fleet []Spaceship) []Spaceship {
	var result []Spaceship
	for _, ship := range fleet {
		if !ship.Launched {
			result = append(result, ship)
		}
	}
	return result
}

// Question 9
func GetTotalCrew(fleet []Spaceship) int {
	total := 0
	for _, ship := range fleet {
		total += ship.Crew
	}
	return total
}

// Question 10
func GetLargestCrewShip(fleet []Spaceship) Spaceship {
	if len(fleet) == 0 {
		return Spaceship{}
	}

	largest := fleet[0]
	for _, ship := range fleet {
		if ship.Crew > largest.Crew {
			largest = ship
		}
	}
	return largest
}

func main() {
	fleet := []Spaceship{
		{Name: "Apollo", Crew: 13, Launched: false},
		{Name: "Houston", Crew: 5, Launched: false},
		{Name: "Doomsday", Crew: 6, Launched: false},
	}

	err := fleet[0].Launch()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = fleet[0].Launch()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fleet[0].Land()

	fmt.Println("\nFleet Status:")
	ShowFleet(fleet)

	fmt.Println("\nLaunched ships:", CountLaunched(fleet))

	ship, err := FindShipByName(fleet, "Houston")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFound ship:", ship.Name)
	}

	fmt.Println("\nUnlaunched Ships:")
	for _, s := range ListUnlaunchedShips(fleet) {
		fmt.Println("-", s.Name)
	}

	fmt.Println("\nTotal crew across all ships:", GetTotalCrew(fleet))

	largest := GetLargestCrewShip(fleet)
	fmt.Println("\nShip with largest crew:", largest.Name)
}
