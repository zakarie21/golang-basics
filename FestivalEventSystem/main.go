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
6. Write a function FindPerformer(lineup []Performer, name string) error that searches for a performer by name
7. Write a function RemovePerformer(lineup []Performer, name string) ([]Performer, error) that removes a performer or returns an error if not found
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

func RemovePerformer(lineup []Performer, name string) ([]Performer, error) {
	for i := 0; i < len(lineup); i++ {
		if lineup[i].Name == name {
			lineup = append(lineup[:i], lineup[i+1:]...)
			return lineup, nil
		}
	}
	return lineup, errors.New("cannot remove performer: " + name + " not found")
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

	// Step 3: Try checking in the same performer again
	err = lineup[0].CheckIn()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 4: Mark performer as done performing
	lineup[0].Perform()

	// Step 5: Display the lineup
	fmt.Println("\nCurrent Festival Lineup:")
	ShowLineup(lineup)

	// Step 6: Search for a performer
	fmt.Println("\nSearching for Skrillex:")
	err = FindPerformer(lineup, "Skrillex")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 7: Remove a performer
	fmt.Println("\nRemoving performer Drake:")
	lineup, err = RemovePerformer(lineup, "Drake")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show updated lineup
	fmt.Println("\nLineup After Removal:")
	ShowLineup(lineup)
}
