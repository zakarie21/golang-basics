package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Make a slice called soldiers that stores the names of soldiers
2. Write AddSoldier(soldiers []string, name string) []string that adds a new soldier
3. Write RemoveSoldier(soldiers []string, name string) ([]string, error) that removes a soldier or errors
4. Write ListSoldiers(soldiers []string) that prints all the soldiers
5. Write FindSoldier(soldiers []string, name string) error that checks if a soldier exists
6. Write CountSoldiers(soldiers []string) that prints how many soldiers are in the army
7. Write ClearArmy(soldiers []string) []string that removes all soldiers from the army
8. Write PromoteSoldier(soldiers []string, name string) error that prints a message if soldier found, otherwise returns an error
*/

func AddSoldier(soldiers []string, name string) []string {
	soldiers = append(soldiers, name)
	return soldiers
}

func RemoveSoldier(soldiers []string, name string) ([]string, error) {
	for i := 0; i < len(soldiers); i++ {
		if soldiers[i] == name {
			soldiers = append(soldiers[:i], soldiers[i+1:]...)
			return soldiers, nil
		}
	}
	return soldiers, errors.New("soldier not found")
}

func ListSoldiers(soldiers []string) {
	if len(soldiers) == 0 {
		fmt.Println("No soldiers in the army")
		return
	}
	fmt.Println("Soldiers in the army:")
	for i := 0; i < len(soldiers); i++ {
		fmt.Println("-", soldiers[i])
	}
}

func FindSoldier(soldiers []string, name string) error {
	for i := 0; i < len(soldiers); i++ {
		if soldiers[i] == name {
			fmt.Println("Found soldier:", name)
			return nil
		}
	}
	return errors.New("soldier not found")
}

func CountSoldiers(soldiers []string) {
	fmt.Println("Total number of soldiers in the army:", len(soldiers))
}

func ClearArmy(soldiers []string) []string {
	soldiers = []string{}
	fmt.Println("All soldiers have been removed from the army.")
	return soldiers
}

func PromoteSoldier(soldiers []string, name string) error {
	for i := 0; i < len(soldiers); i++ {
		if soldiers[i] == name {
			fmt.Println("Soldier", name, "has been promoted!")
			return nil
		}
	}
	return errors.New("soldier not found")
}

func main() {
	// Start with an empty slice
	soldiers := []string{}

	// Add some soldiers
	soldiers = AddSoldier(soldiers, "John")
	soldiers = AddSoldier(soldiers, "Mike")
	soldiers = AddSoldier(soldiers, "Sara")

	// List them
	ListSoldiers(soldiers)

	// Count them
	CountSoldiers(soldiers)

	// Try to find a soldier
	err := FindSoldier(soldiers, "Mike")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Try to promote a soldier
	err = PromoteSoldier(soldiers, "Sara")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Try to promote someone not in the list
	err = PromoteSoldier(soldiers, "Alex")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Try to remove a soldier
	soldiers, err = RemoveSoldier(soldiers, "John")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// List again after removal
	ListSoldiers(soldiers)

	// Clear the army
	soldiers = ClearArmy(soldiers)

	// Show after clearing
	ListSoldiers(soldiers)

	// Final count
	CountSoldiers(soldiers)
}

