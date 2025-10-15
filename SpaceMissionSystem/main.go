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
}
