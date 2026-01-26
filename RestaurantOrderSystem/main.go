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
6. Write FindMeal function to check if a meal exists in the menu and print its details
7. Write UpdatePrice function to change the price of a meal if it exists
8. Write RestockMeal function to mark a meal as available again or return an error if not found
9. CountAvailableMeals: return the number of meals that are currently available
10. GetTotalMenuValue: return the total price of all available meals
11. GetFirstAvailableMeal: return the first available meal or an error if none exist
12. HasMeal: return true if a meal exists in the menu
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

func FindMeal(menu []Meal, name string) error {
	for _, meal := range menu {
		if meal.Name == name {
			fmt.Println("Found Meal:", meal.Name, "| Price:", meal.Price, "| Available:", meal.Available)
			return nil
		}
	}
	return errors.New("meal not found")
}

func UpdatePrice(menu []Meal, name string, newPrice float64) ([]Meal, error) {
	for i := 0; i < len(menu); i++ {
		if menu[i].Name == name {
			menu[i].Price = newPrice
			fmt.Println("Price for", name, "updated to", newPrice)
			return menu, nil
		}
	}
	return menu, errors.New("meal not found")
}

func RestockMeal(menu []Meal, name string) ([]Meal, error) {
	for i := 0; i < len(menu); i++ {
		if menu[i].Name == name {
			if !menu[i].Available {
				menu[i].Available = true
				fmt.Println("Meal", name, "has been restocked and is now available again.")
				return menu, nil
			}
			return menu, errors.New("meal is already available")
		}
	}
	return menu, errors.New("meal not found")
}

// Question 9
func CountAvailableMeals(menu []Meal) int {
	count := 0
	for _, meal := range menu {
		if meal.Available {
			count++
		}
	}
	return count
}

// Question 10
func GetTotalMenuValue(menu []Meal) float64 {
	total := 0.0
	for _, meal := range menu {
		if meal.Available {
			total += meal.Price
		}
	}
	return total
}

// Question 11
func GetFirstAvailableMeal(menu []Meal) (Meal, error) {
	for _, meal := range menu {
		if meal.Available {
			return meal, nil
		}
	}
	return Meal{}, errors.New("no available meals")
}

// Question 12
func HasMeal(menu []Meal, name string) bool {
	for _, meal := range menu {
		if meal.Name == name {
			return true
		}
	}
	return false
}

func main() {
	menu := []Meal{}

	menu = AddMeal(menu, Meal{Name: "Burger", Price: 5.99, Available: true})
	menu = AddMeal(menu, Meal{Name: "Pizza", Price: 8.49, Available: true})
	menu = AddMeal(menu, Meal{Name: "Salad", Price: 4.25, Available: true})

	fmt.Println("Initial Menu:")
	ListMenu(menu)

	var err error
	menu, err = OrderMeal(menu, "Pizza")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nMenu After Ordering:")
	ListMenu(menu)

	fmt.Println("\nFinding a meal:")
	err = FindMeal(menu, "Burger")
	if err != nil {
		fmt.Println("Error:", err)
	}

	menu, err = UpdatePrice(menu, "Salad", 5.50)
	if err != nil {
		fmt.Println("Error:", err)
	}

	menu, err = RestockMeal(menu, "Pizza")
	if err != nil {
		fmt.Println("Error:", err)
	}

	menu, err = RestockMeal(menu, "Sushi")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nFinal Menu:")
	ListMenu(menu)

	fmt.Println("\nAvailable meals count:", CountAvailableMeals(menu))
	fmt.Printf("Total value of available meals: $%.2f\n", GetTotalMenuValue(menu))

	// Demo Question 11
	meal, err := GetFirstAvailableMeal(menu)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("First available meal:", meal.Name)
	}

	// Demo Question 12
	fmt.Println("Does Burger exist?", HasMeal(menu, "Burger"))
	fmt.Println("Does Sushi exist?", HasMeal(menu, "Sushi"))
}
