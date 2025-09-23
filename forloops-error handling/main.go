package main


import (
	"errors"
	"fmt"
)

func PrintEvenNumbers(n int) {
	fmt.Println("Even numbers up to", n, ":")
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func CountPositiveNumbers(numbers []int) int {
	count := 0
	for _, num := range numbers {
		if num > 0 {
			count++
		}
	}
	return count
}

func DivideNumbers(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func FindMax(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty slice")
	}

	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max, nil
}

func FilterNegativeNumbers(numbers []int) []int {
	var filtered []int
	for _, num := range numbers {
		if num >= 0 {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

func main() {
	PrintEvenNumbers(10)

	values := []int{-1, 3, 0, 5, -7, 8}
	count := CountPositiveNumbers(values)
	fmt.Println("Number of positive numbers:", count)

	result, err := DivideNumbers(10, 2)
	if err != nil {
		fmt.Println("Error dividing numbers:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	_, err2 := DivideNumbers(5, 0)
	if err2 != nil {
		fmt.Println("Expected error:", err2)
	}

	max, err3 := FindMax(values)
	if err3 != nil {
		fmt.Println("Error finding max:", err3)
	} else {
		fmt.Println("Max value is:", max)
	}

	filtered := FilterNegativeNumbers(values)
	fmt.Println("Filtered (no negatives):", filtered)
}
