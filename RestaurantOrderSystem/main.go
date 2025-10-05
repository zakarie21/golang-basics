package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Meal struct with Name, Price, and Available
2. Make a slice of Meal called menu
3. Write AddMeal function to add a meal to the menu
4. Write OrderMeal function to mark a meal as unavailable or return error if not found
5. Write ListMenu function to show all meals
*/

type Meal struct {
	Name      string
	Price     float64
	Available bool
}

func AddMeal(menu []Meal, meal Meal) []Meal {
	menu = append(menu, meal)
	return menu
}

func OrderMeal(menu []Meal, name string) ([]Meal, error) {
	for i := 0; i < len(menu); i++ {
		if menu[i].Name == name {
			if menu[i].Available {
				menu[i].Available = false
				return menu, nil
			}
			return menu, errors.New("meal already ordered")
		}
	}
	return menu, errors.New("meal not found")
}

func ListMenu(menu []Meal) {
	for _, meal := range menu {
		fmt.Println("Meal:", meal.Name, "| Price:", meal.Price, "| Available:", meal.Available)
	}
}

func main() {
	//Create an empty menu (slice of meals)
	menu := []Meal{}

	// Step 2: Add meals to the menu
	menu = AddMeal(menu, Meal{Name: "Burger", Price: 5.99, Available: true})
	menu = AddMeal(menu, Meal{Name: "Pizza", Price: 8.49, Available: true})
	menu = AddMeal(menu, Meal{Name: "Salad", Price: 4.25, Available: true})

	//Show the menu
	fmt.Println("Initial Menu:")
	ListMenu(menu)

	//Order a meal
	var err error
	menu, err = OrderMeal(menu, "Pizza")
	if err != nil {
		fmt.Println("Error:", err)
	}

	//Show the updated menu
	fmt.Println("\nMenu After Ordering:")
	ListMenu(menu)
}
