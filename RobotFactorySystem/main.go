package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Robot struct with fields ID, Model, and Active
2. Make a slice called factory to store all robots
3. Write a method Activate() that sets Active to true or returns an error if already active
4. Write a method Deactivate() that sets Active to false
5. Write a function ShowFactory(factory []Robot) that loops and prints robot details
6. Write a function AddRobot(factory []Robot, newRobot Robot) []Robot to add a new robot to the factory
7. Write a function FindRobot(factory []Robot, id int) (Robot, error) that returns a robot by ID or an error if not found
8. Write a function CountActiveRobots(factory []Robot) int to count the active robots in the factory
*/

type Robot struct {
	ID     int
	Model  string
	Active bool
}

func (r *Robot) Activate() error {
	if r.Active {
		return errors.New("robot " + r.Model + " is already active")
	}
	r.Active = true
	return nil
}

func (r *Robot) Deactivate() {
	r.Active = false
}

func ShowFactory(factory []Robot) {
	for _, r := range factory {
		status := "inactive"
		if r.Active {
			status = "active"
		}
		fmt.Println("Robot ID:", r.ID, "| Model:", r.Model, "| Status:", status)
	}
}

func AddRobot(factory []Robot, newRobot Robot) []Robot {
	factory = append(factory, newRobot)
	return factory
}

func FindRobot(factory []Robot, id int) (Robot, error) {
	for _, r := range factory {
		if r.ID == id {
			return r, nil
		}
	}
	return Robot{}, errors.New("robot not found")
}

func CountActiveRobots(factory []Robot) int {
	count := 0
	for _, r := range factory {
		if r.Active {
			count = count + 1
		}
	}
	return count
}

func main() {

	// Step 1: Create a factory with a few robots
	factory := []Robot{
		{ID: 1, Model: "Robot-1", Active: false},
		{ID: 2, Model: "Robot-2", Active: false},
		{ID: 3, Model: "Robot-3", Active: false},
	}

	// Step 2: Activate a robot
	err := factory[1].Activate()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 3: Try activating the same robot again (to trigger error)
	err = factory[1].Activate()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Step 4: Deactivate a robot
	factory[1].Deactivate()

	// Step 5: Show current factory status
	fmt.Println("\nFactory Robot Status:")
	ShowFactory(factory)

	// Step 6: Add a new robot to the factory
	newRobot := Robot{ID: 4, Model: "Robot-4", Active: false}
	factory = AddRobot(factory, newRobot)
	fmt.Println("\nAdded new robot to factory:")
	ShowFactory(factory)

	// Step 7: Find a robot by ID
	found, err := FindRobot(factory, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFound robot:", found.Model)
	}

	// Step 8: Count active robots
	activeCount := CountActiveRobots(factory)
	fmt.Println("\nNumber of active robots:", activeCount)
}

