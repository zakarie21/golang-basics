package main

import (
	"fmt"
)

//
// 1. Write a function that takes two integers and returns their sum.
//
func addNumbers(a int, b int) int {
	// just return the sum
	return a + b
}

//
// 2. Write a function that divides two integers and returns both the result and an error if dividing by zero.
//
func divideNumbers(a int, b int) (int, string) {
	if b == 0 {
		// return error message as string
		return 0, "cannot divide by zero"
	}
	return a / b, ""
}

//
// 3. Write a function that takes a slice of integers and returns the largest number.
// If the slice is empty, return an error.
//
func findLargest(numbers []int) (int, string) {
	if len(numbers) == 0 {
		return 0, "slice is empty"
	}
	largest := numbers[0]
	for _, number := range numbers {
		if number > largest {
			largest = number
		}
	}
	return largest, ""
}

//
// 4. Write a function that checks if a string is empty.
// If it is empty, return an error.
//
func checkString(text string) (string, string) {
	if text == "" {
		return "", "string is empty"
	}
	return text, ""
}

//
// 5. Write a function that takes a number and says if it is positive.
// If it is zero or negative, return an error.
//
func checkPositive(num int) (int, string) {
	if num <= 0 {
		return 0, "number is not positive"
	}
	return num, ""
}

func main() {
	// Q1
	fmt.Println("1. Sum:", addNumbers(3, 7))

	// Q2
	result, err := divideNumbers(10, 0)
	if err != "" {
		fmt.Println("2. Error:", err)
	} else {
		fmt.Println("2. Division:", result)
	}

	// Q3
	numbers := []int{5, 10, 3, 8}
	largest, err := findLargest(numbers)
	if err != "" {
		fmt.Println("3. Error:", err)
	} else {
		fmt.Println("3. Largest number:", largest)
	}

	// Q4
	value, err := checkString("")
	if err != "" {
		fmt.Println("4. Error:", err)
	} else {
		fmt.Println("4. String is:", value)
	}

	// Q5
	number, err := checkPositive(-5)
	if err != "" {
		fmt.Println("5. Error:", err)
	} else {
		fmt.Println("5. Number is positive:", number)
	}
}
