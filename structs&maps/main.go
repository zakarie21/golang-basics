package main

import "fmt"


//
// 1. Create a struct called Student with fields name and age.
// Store three students in a map where the key is the student’s ID (an int). Print all students.
//
type Student struct {
	name string
	age  int
}

//
// 2. Write a function that takes the map of students and prints only the names.
//
func printNames(students map[int]Student) {
	for id, student := range students {
		fmt.Println("ID:", id, "Name:", student.name)
	}
}

//
// 3. Add a new student to the map. Show what happens if you add another student with the same ID.
//
func addStudent(students map[int]Student, id int, student Student) {
	students[id] = student
}

//
// 4. Write a function that updates the age of a student in the map given the ID.
// If the ID doesn’t exist, return an error message.
//
func updateAge(students map[int]Student, id int, newAge int) string {
	student, ok := students[id]
	if !ok {
		return "student ID not found"
	}
	student.age = newAge
	students[id] = student
	return ""
}

//
// 5. Write a function that deletes a student from the map by ID.
// Show what happens if you try deleting an ID that doesn’t exist.
//
func deleteStudent(students map[int]Student, id int) string {
	_, ok := students[id]
	if !ok {
		return "student ID not found"
	}
	delete(students, id)
	return ""
}

func main() {
	// Q1
	students := make(map[int]Student)
	students[1] = Student{name: "Alice", age: 20}
	students[2] = Student{name: "Bob", age: 21}
	students[3] = Student{name: "Charlie", age: 19}

	fmt.Println("1. All students:", students)

	// Q2
	fmt.Println("2. Names of students:")
	printNames(students)

	// Q3
	addStudent(students, 4, Student{name: "David", age: 22})
	fmt.Println("3. Added David:", students)

	addStudent(students, 2, Student{name: "Eve", age: 23})
	fmt.Println("3. Added Eve with same ID as Bob:", students)

	// Q4
	err := updateAge(students, 3, 25)
	if err != "" {
		fmt.Println("4. Error:", err)
	} else {
		fmt.Println("4. Updated Charlie’s age:", students[3])
	}

	err = updateAge(students, 10, 30)
	if err != "" {
		fmt.Println("4. Error:", err)
	}

	// Q5
	err = deleteStudent(students, 1)
	if err != "" {
		fmt.Println("5. Error:", err)
	} else {
		fmt.Println("5. Deleted student with ID 1")
	}

	err = deleteStudent(students, 100)
	if err != "" {
		fmt.Println("5. Error:", err)
	}

	fmt.Println("5. Final students:", students)
}
