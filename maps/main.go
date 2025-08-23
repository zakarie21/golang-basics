package main

import "fmt"

// Q25 needs its own helper function
func countWords(text string) map[string]int {
	counts := make(map[string]int)
	word := ""

	for i := 0; i < len(text); i++ {
		letter := string(text[i])

		if letter == " " {
			if word != "" {
				if _, found := counts[word]; found {
					counts[word] = counts[word] + 1
				} else {
					counts[word] = 1
				}
				word = "" // reset after finishing a word
			}
		} else {
			word = word + letter // build word letter by letter
		}
	}

	// add the last word
	if word != "" {
		if _, found := counts[word]; found {
			counts[word] = counts[word] + 1
		} else {
			counts[word] = 1
		}
	}

	return counts
}

func main() {

	// 24. How do you check if a key exists in a map?
	// Answer: You use value, ok := map[key].
	// "value" gives the value, "ok" is true if the key exists, false if it does not.
	students := map[string]int{"bob": 10, "sue": 8}
	value1, exists1 := students["bob"]
	fmt.Println("24:", value1, exists1) // 10 true
	value2, exists2 := students["tom"]
	fmt.Println("24:", value2, exists2) // 0 false
	// Here value1 = 10 and exists1 = true because "bob" exists.
	// value2 = 0 (zero value) and exists2 = false because "tom" does not exist.

	// 25. Write a function to count the frequency of words in a string using a map.
	// Answer: We loop through the string, build words, and count how many times each word appears.
	sentence := "go go play now go play"
	wordCounts := countWords(sentence)
	fmt.Println("25:", wordCounts) // map[go:3 play:2 now:1]

	// 26. Can you compare two maps directly using ==? Why or why not?
	// Answer: No, you cannot compare two maps with == because maps are reference types.
	// You can only compare if a map is nil.
	mapOne := map[string]int{"a": 1}
	mapTwo := map[string]int{"a": 1}
	// fmt.Println(mapOne == mapTwo) // not allowed
	fmt.Println("26:", mapOne == nil, mapTwo == nil)

	// 27. What happens if you try to read a value from a nil map?
	// Answer: Reading from a nil map does not crash. It gives the zero value of the type.
	var grades map[string]int         // nil map
	fmt.Println("27:", grades["sam"]) // prints 0, zero value for int

	// 28. What happens if you try to write to a nil map?
	// Answer: Writing to a nil map causes a runtime panic.
	// grades["sam"] = 99 // causes panic if uncommented
	fmt.Println("28: writing to nil map causes panic")

	// 29. How do you delete a key from a map? What happens if the key doesn't exist?
	// Answer: You use delete(map, key). If the key exists, it is removed.
	// If it doesn’t exist, nothing happens.
	pets := map[string]int{"a": 1, "b": 2}
	delete(pets, "a")        // removes key "a"
	delete(pets, "c")        // does nothing
	fmt.Println("29:", pets) // map[b:2]

	// 30. Write a program that swaps keys and values in a map (assuming values are unique).
	// Answer: Make a new map, and for each pair, put the value as the new key and the key as the new value.
	animals := map[string]string{"cat": "meow", "dog": "bark"}
	swapped := make(map[string]string)
	for key, value := range animals {
		swapped[value] = key
	}
	fmt.Println("30:", swapped) // map[meow:cat bark:dog]

	// 31. Can map iteration order be guaranteed in Go? Why?
	// Answer: No, Go does not guarantee the order of maps.
	// The order is random each time you loop over a map.
	colors := map[string]string{"r": "red", "g": "green", "b": "blue"}
	fmt.Println("31: map iteration order example:")
	for key, value := range colors {
		fmt.Println(key, value)
	}
}
