package main

import "fmt"

//## *3. Arrays (5 Questions)*





func main() {

    numbers := [5]int{1, 2, 3, 4, 5}         
    fmt.Println("Original:", numbers)        
    reversed := reverse(numbers)             
    fmt.Println("Reversed:", reversed)  
	
	//  How do you initialize an array of length 5 with default values of 10?
		
		nums := [5]int{10, 10, 10, 10, 10}  
		fmt.Println(nums)                   


	//  What is the zero value of an array of strings?

	// The zero value means what Go fills in when you do not assign values. For strings, the zero value is an empty string "".
		var array [3]string           
		fmt.Println(array)            


	//  How can you convert an array to a slice?

	// To convert an array to a slice, you have to use the slice operator [:]. This creates a slice view of the array.

		arr := [5]int{1, 2, 3, 4, 5}   
		slice := arr[:]                
		fmt.Println(slice)             


}


//  How are arrays passed to functions in Go — by value or reference?

// In Go, arrays are passed to functions by value. This means the function gets a copy of the array, not the original one. 
// If you change the array inside the function, the original array will not change.


//  Write a function that reverses an array in place.

func reverse(arr [5]int) [5]int {
    for i := 0; i < len(arr)/2; i++ {     
        j := len(arr) - 1 - i             
        arr[i], arr[j] = arr[j], arr[i]   
    }
    return arr                            
}



