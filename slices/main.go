package main

import "fmt"




func main() {

	numbers := []int{20, 60, 30, 15, 75}
	fmt.Println("Original slice:", numbers)

	numbers = removeElement(numbers, 2) 
	fmt.Println("After removal:", numbers)



	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", nums)
	fmt.Println("Length:", len(nums))   // 5
	fmt.Println("Capacity:", cap(nums)) // 5

	part := nums[:3]
	fmt.Println("Sub-slice:", part)
	fmt.Println("Length:", len(part))   // 3
	fmt.Println("Capacity:", cap(part)) // 5 (shares the same underlying array)

	mySlice := []int{1, 2, 3, 4, 5}
	newSlice := multiplyByTwo(mySlice)
	fmt.Println("Original slice:", mySlice)
	fmt.Println("New slice:", newSlice)

	original := []int{1, 2, 3, 4, 5}
	backup := make([]int, len(original)) // make a new slice
	copy(backup, original)               // copy values
	fmt.Println("Original slice:", original)
	fmt.Println("Copied slice:", backup)
}


// 11. How do slices differ from arrays in Go in terms of memory and length?
// In Go, arrays have a fixed size that cannot change and copying an array makes a full copy of all its elements.
// Slices are more flexible because they point to an underlying array and keep track of both their length and capacity. When you append beyond the capacity,
// Go automatically creates a bigger array and moves the data. Arrays stay fixed while slices can grow, which makes slices much more common in Go.


// 12. Write a function that removes an element at index i from a slice without breaking order.
func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice 
	}
	return append(slice[:i], slice[i+1:]...)
}


// 13. What is the difference between slice capacity and length? Show an example.
//The length of a slice is how many elements it currently has, 
// and the capacity is how many elements it can hold in the underlying array before Go needs to create a new one.


// 14. How does slicing affect the underlying array?
// When you make a slice from another slice or array, it doesn’t make a new array, 
// it just points to the same one. So if you change something in the new slice, 
// it also changes in the original slice because they share the same data in memory. 
// The slice just keeps track of where it starts, its length, and capacity.

// 15. What happens if you append to a slice that has reached its capacity?
// If you append to a slice and it has already reached its capacity, Go automatically creates a new bigger array, copies the existing elements into it, 
// and then adds the new element. The slice then points to this new array, so the original underlying array is no longer affected.

// 16. Write a function that takes a slice of integers and returns a new slice with each element doubled.

func multiplyByTwo(nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = v * 2
	}
	return result
}


// 17. Explain the copy function for slices and give a use case.
// In Go, copy is used to take all or some elements from one slice and put them into another slice. 
// You need to make the second slice first, and copy will fill it with values from the first slice.


// 18. What happens if you slice beyond the slice's length but within its capacity?
// If you slice beyond a slice’s length but stay within its capacity, Go lets you do it and the new slice will include more elements. 
// It still uses the same underlying array, so changes in the new slice can affect the original slice.



