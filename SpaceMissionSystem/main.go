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

// NEW FUNCTION (Question 7)
func FindShipByName(fleet []Spaceship, name string) (*Spaceship, error) {
	for i := range fleet {
		if fleet[i].Name == name {
			return &fleet[i], nil
		}
	}
	return nil, errors.New("ship not found: " + name)
}

// NEW FUNCTION (Question 8)
func ListUnlaunchedShips(fleet []Spaceship) []Spaceship {
	var result []Spaceship
	for _, ship := range fleet {
		if !ship.Launched {
			result = append(result, ship)
		}
	}
	return result
}

func main() {
	// Step 1: Create a fleet of spaceships
	fleet := []Spaceship{
		{Name: "Apollo", Crew: 13, Launched: false},
		{Name: "Houston", Crew: 5, Launched: false},
		{Name: "Doomsday", Crew: 6, Launched: false},
	}

	// Step 2: Launch the Apollo
	err := fleet[0].Launch()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 3: Try launching Apollo again (should show an error)
	err = fleet[0].Launch()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 4: Land the Apollo
	fleet[0].Land()

	// Step 5: Show the current fleet status
	fmt.Println("\nFleet Status:")
	ShowFleet(fleet)

	// count launched ships
	fmt.Println("\nLaunched ships:", CountLaunched(fleet))

	// use new FindShipByName function
	ship, err := FindShipByName(fleet, "Houston")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFound ship:", ship.Name)
	}

	// use new ListUnlaunchedShips function
	fmt.Println("\nUnlaunched Ships:")
	for _, s := range ListUnlaunchedShips(fleet) {
		fmt.Println("-", s.Name)
	}
}
