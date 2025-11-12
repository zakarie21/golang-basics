package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Performer struct with fields Name, Genre, and Ready
2. Make a slice called lineup to store all performers
3. Write a method CheckIn() that sets Ready to true or returns an error if already checked in
4. Write a method Perform() that sets Ready to false
5. Write a function ShowLineup(lineup []Performer) that loops and prints performer details
6. Write a function FindPerformer(lineup []Performer, name string) error that searches for a performer by name and prints their details or returns an error if not found
*/

type Performer struct {
	Name  string
	Genre string
	Ready bool
}

func (p *Performer) CheckIn() error {
	if p.Ready {
		return errors.New(p.Name + " is already checked in")
	}
	p.Ready = true
	return nil
}

func (p *Performer) Perform() {
	p.Ready = false
}

func ShowLineup(lineup []Performer) {
	for _, p := range lineup {
		status := "not ready"
		if p.Ready {
			status = "checked in"
		}
		fmt.Println(p.Name, "-", p.Genre, ":", status)
	}
}

func FindPerformer(lineup []Performer, name string) error {
	for _, p := range lineup {
		if p.Name == name {
			status := "not ready"
			if p.Ready {
				status = "checked in"
			}
			fmt.Println("Found performer:", p.Name, "| Genre:", p.Genre, "| Status:", status)
			return nil
		}
	}
	return errors.New("performer not found: " + name)
}

func main() {
	// Step 1: Create a lineup of performers
	lineup := []Performer{
		{Name: "Artic Monkey", Genre: "Brit Pop", Ready: false},
		{Name: "Skrillex", Genre: "EDM", Ready: false},
		{Name: "Drake", Genre: "Rap", Ready: false},
	}

	// Step 2: Check in a performer
	err := lineup[0].CheckIn()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 3: Try checking in the same performer again (should trigger an error)
	err = lineup[0].CheckIn()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 4: Mark a performer as done performing
	lineup[0].Perform()

	// Step 5: Display all performer statuses
	fmt.Println("\nCurrent Festival Lineup:")
	ShowLineup(lineup)

	// Step 6: Find a performer in the lineup
	fmt.Println("\nSearching for performer 'Artic Monkey':")
	err = FindPerformer(lineup, "Artic Monkey")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nSearching for performer 'Beyonce':")
	err = FindPerformer(lineup, "Beyonce")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
