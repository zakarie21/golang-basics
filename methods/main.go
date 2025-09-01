package main

import "fmt"

// 1. Book struct with a method to print info
type Book struct {
	title  string
	author string
}

// prints the book info
func (book Book) printInfo() {
	fmt.Println("1: Book title:", book.title, "Author:", book.author)
}

// 2. method with pointer receiver so I can change the title
func (book *Book) updateTitle(newTitle string) {
	book.title = newTitle
}

// 3. method to check if the book is by a certain author
func (book Book) isByAuthor(authorName string) bool {
	if book.author == authorName {
		return true
	} else {
		return false
	}
}

// 4. another struct with its own method
type Student struct {
	name  string
	grade int
}

// checks if grade is >= 50
func (student Student) passed() bool {
	if student.grade >= 50 {
		return true
	} else {
		return false
	}
}

func main() {
	// Q1
	myBook := Book{title: "Go Basics", author: "Alice"}
	myBook.printInfo()
	// I call the method on myBook, it prints the book info

	// Q2
	fmt.Println("2: Before title update:", myBook.title)
	// I print the title before changing it
	myBook.updateTitle("Go for Beginners")
	// I call the method to update the title
	fmt.Println("2: After title update:", myBook.title)
	// I print the new title to check that it worked

	// Q3
	fmt.Println("3: Is myBook by Alice?", myBook.isByAuthor("Alice"))
	// this prints true because the author is Alice
	fmt.Println("3: Is myBook by Bob?", myBook.isByAuthor("Bob"))
	// this prints false because the author is not Bob

	// Q4
	studentA := Student{name: "John", grade: 75}
	studentB := Student{name: "Mary", grade: 40}

	fmt.Println("4:", studentA.name, "passed?", studentA.passed())
	// true because grade is 75 which is >= 50
	fmt.Println("4:", studentB.name, "passed?", studentB.passed())
	// false because grade is 40 which is < 50

	// I notice methods belong to the struct they are written for
	myBook.printInfo()
	// Book method only works on Book, Student method only works on Student
}
