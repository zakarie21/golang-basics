package main

import "fmt"

/*
Questions:
1. Create a map called powerGrid with city name as key and power usage as value
2. Add three cities with example power usage
3. Write a function ShowPower() to print cities and their usage
4. Write a function AddUsage() to increase power usage for a city
5. Write a function FindHighestUsage() to find the city with the highest power usage
6. In main, call all these functions and show results
7. Write a function RemoveCity() to remove a city from the map if it exists
*/

func ShowPower(powerGrid map[string]int) {
	for city, usage := range powerGrid {
		fmt.Println(city, ":", usage, "MW")
	}
}

func AddUsage(powerGrid map[string]int, city string, amount int) {
	powerGrid[city] = powerGrid[city] + amount
}

func FindHighestUsage(powerGrid map[string]int) {
	highestCity := ""
	highestUsage := 0

	for city, usage := range powerGrid {
		if usage > highestUsage {
			highestUsage = usage
			highestCity = city
		}
	}

	fmt.Println("\nCity with highest power usage:", highestCity, "-", highestUsage, "MW")
}

func RemoveCity(powerGrid map[string]int, city string) {
	if _, exists := powerGrid[city]; exists {
		delete(powerGrid, city)
		fmt.Println("\nRemoved city:", city)
	} else {
		fmt.Println("\nCity not found:", city)
	}
}

func main() {
	// Step 1: Create a map for city power usage
	powerGrid := map[string]int{
		"London":     900,
		"Manchester": 750,
		"Birmingham": 820,
	}

	// Step 2: Show starting power usage
	fmt.Println("Initial power usage:")
	ShowPower(powerGrid)

	// Step 3: Increase London's power usage
	AddUsage(powerGrid, "London", 150)

	// Step 4: Show updated power usage
	fmt.Println("\nAfter adding usage:")
	ShowPower(powerGrid)

	// Step 5: Find and show the city with the highest usage
	FindHighestUsage(powerGrid)

	// Step 6: Remove a city from the map
	RemoveCity(powerGrid, "Manchester")

	// Step 7: Show final power grid
	fmt.Println("\nFinal power grid:")
	ShowPower(powerGrid)
}

