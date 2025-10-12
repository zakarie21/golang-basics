package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create an Animal struct with fields Name, Type, and Fed
2. Write a method Feed() that sets Fed to true or returns an error if already fed
3. Write a method ResetFeeding() that sets Fed to false
4. Write a function ShowAnimalStatus(a Animal) that prints name, type, and feeding status
5. In main(), create animals, feed one, try feeding again (error), and reset feeding
*/

type Animal struct {
	Name string
	Type string
	Fed  bool
}

func (a *Animal) Feed() error {
	if a.Fed {
		return errors.New(a.Name + " has already been fed")
	}
	a.Fed = true
	return nil
}

func (a *Animal) ResetFeeding() {
	a.Fed = false
}

func ShowAnimalStatus(a Animal) {
	status := "hungry"
	if a.Fed {
		status = "fed"
	}
	fmt.Println(a.Name, "the", a.Type, "is currently", status)
}

func main() {
	// Step 1: Create some animals
	cow := Animal{Name: "Biff", Type: "Cow", Fed: false}
	chicken := Animal{Name: "Chip", Type: "Chicken", Fed: true}
	sheep := Animal{Name: "Kipper", Type: "Sheep", Fed: true}

	// Step 2: Feed the cow
	err := cow.Feed()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 3: Try feeding the cow again (should show error)
	err = cow.Feed()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 4: Show all animal statuses
	fmt.Println("\nAnimal statuses:")
	ShowAnimalStatus(cow)
	ShowAnimalStatus(chicken)
	ShowAnimalStatus(sheep)

	// Step 5: Reset feeding for next day
	cow.ResetFeeding()
	chicken.ResetFeeding()
	sheep.ResetFeeding()

	fmt.Println("\nFeeding reset complete. Current statuses:")
	ShowAnimalStatus(cow)
	ShowAnimalStatus(chicken)
	ShowAnimalStatus(sheep)
}