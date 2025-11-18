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
8. Write a function CountReady(lineup []Performer) int that returns how many performers are checked in
9. Write a function AddPerformer(lineup []Performer, p Performer) []Performer that appends a new performer to the lineup
10. Write a function CountByGenre(lineup []Performer, genre string) int that returns how many performers belong to a given genre
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
			status := "checked in"
			if !p.Ready {
				status = "not ready"
			}
			fmt.Println("Found:", p.Name, "| Genre:", p.Genre, "| Status:", status)
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

func CountReady(lineup []Performer) int {
	count := 0
	for _, p := range lineup {
		if p.Ready {
			count++
		}
	}
	return count
}

func AddPerformer(lineup []Performer, p Performer) []Performer {
	return append(lineup, p)
}


func CountByGenre(lineup []Performer, genre string) int {
	count := 0
	for _, p := range lineup {
		if p.Genre == genre {
			count++
		}
	}
	return count
}

func main() {

	// Step 1: Create the festival lineup
	lineup := []Performer{
		{Name: "Artic Monkey", Genre: "Brit Pop", Ready: false},
		{Name: "Skrillex", Genre: "EDM", Ready: false},
		{Name: "Drake", Genre: "Rap", Ready: false},
		{Name: "Taylor Swift", Genre: "Pop", Ready: false},
		{Name: "Kanye West", Genre: "Rap", Ready: false},
		{Name: "Yung Blud", Genre: "Rock", Ready: false},
	}

	// Step 2: Check in multiple performers
	fmt.Println("\nChecking in Taylor Swift and Skrillex:")
	lineup[3].CheckIn() // Taylor Swift
	lineup[1].CheckIn() // Skrillex

	// Step 3: Try checking in someone already checked in
	fmt.Println("\nTrying to check in Skrillex again:")
	err := lineup[1].CheckIn()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 4: Performer finishes performing
	fmt.Println("\nTaylor Swift performs:")
	lineup[3].Perform()

	// Step 5: Display the lineup
	fmt.Println("\nCurrent Lineup:")
	ShowLineup(lineup)

	// Step 6: Search for a performer
	fmt.Println("\nSearching for Kanye West:")
	err = FindPerformer(lineup, "Kanye West")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 7: Remove a performer
	fmt.Println("\nRemoving performer Drake:")
	lineup, err = RemovePerformer(lineup, "Drake")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 8: Count ready performers
	fmt.Println("\nNumber of performers checked in:")
	fmt.Println("Ready performers:", CountReady(lineup))

	// Step 9: Add a new performer
	fmt.Println("\nAdding a new performer: Billie Eilish")
	lineup = AddPerformer(lineup, Performer{Name: "Billie Eilish", Genre: "Pop", Ready: true})

	// Step 10: Count by genre
	fmt.Println("\nNumber of Rap performers:", CountByGenre(lineup, "Rap"))

	// Show final lineup
	fmt.Println("\nFinal Lineup:")
	ShowLineup(lineup)
}

