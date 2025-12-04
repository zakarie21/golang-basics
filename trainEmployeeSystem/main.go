package main

import (
	"errors"
	"fmt"
)

// 1. employees slice to hold names of train employees
// 2. AddEmployee: add new employee
// 3. RemoveEmployee: remove employee or error if not found
// 4. ListEmployees: print all employees
// 5. FindEmployee: check if employee exists or error
// 6. CountEmployees: return how many employees are in the list
// 7. ClearEmployees: remove all employees from the list

func AddEmployee(employees []string, name string) []string {
	employees = append(employees, name)
	return employees
}

func RemoveEmployee(employees []string, name string) ([]string, error) {
	for i := 0; i < len(employees); i++ {
		if employees[i] == name {
			employees = append(employees[:i], employees[i+1:]...)
			return employees, nil
		}
	}
	return employees, errors.New("employee not found")
}

func ListEmployees(employees []string) {
	if len(employees) == 0 {
		fmt.Println("No employees registered")
		return
	}
	fmt.Println("Train Employees:")
	for i := 0; i < len(employees); i++ {
		fmt.Println("-", employees[i])
	}
}

func FindEmployee(employees []string, name string) error {
	for i := 0; i < len(employees); i++ {
		if employees[i] == name {
			fmt.Println("Found employee:", name)
			return nil
		}
	}
	return errors.New("employee not found")
}

func CountEmployees(employees []string) int {
	return len(employees)
}

func ClearEmployees(employees []string) []string {
	return []string{}
}

func main() {
	employees := []string{}

	// Add employees
	employees = AddEmployee(employees, "Alice")
	employees = AddEmployee(employees, "Bob")
	employees = AddEmployee(employees, "Charlie")

	// List employees
	ListEmployees(employees)

	// Find one
	err := FindEmployee(employees, "Bob")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Remove one
	employees, err = RemoveEmployee(employees, "Alice")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// List again
	ListEmployees(employees)

	// Try to remove someone not there
	employees, err = RemoveEmployee(employees, "David")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Total employees:", CountEmployees(employees))

	// Clear all employees
	employees = ClearEmployees(employees)
	fmt.Println("After clearing:")
	ListEmployees(employees)
}
