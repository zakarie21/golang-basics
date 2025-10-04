package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Book struct with Title, Author, Available
2. Make a slice of Book called library
3. Write AddBook function to add a book to library
4. Write BorrowBook function with error if not found
5. Define Printable interface and implement PrintInfo for Book
*/

type Book struct {
	Title     string
	Author    string
	Available bool
}

type Printable interface {
	PrintInfo()
}

func (b Book) PrintInfo() {
	fmt.Println("Title:", b.Title, "| Author:", b.Author, "| Available:", b.Available)
}

func AddBook(library []Book, book Book) []Book {
	library = append(library, book)
	return library
}

func BorrowBook(library []Book, title string) ([]Book, error) {
	for i := 0; i < len(library); i++ {
		if library[i].Title == title {
			if library[i].Available {
				library[i].Available = false
				return library, nil
			}
			return library, errors.New("book already borrowed")
		}
	}
	return library, errors.New("book not found")
}

func main() {
	// Create an empty library (slice of books)
	library := []Book{}

	// Add some books
	library = AddBook(library, Book{Title: "1984", Author: "George Orwell", Available: true})
	library = AddBook(library, Book{Title: "Moby Dick", Author: "Herman Melville", Available: true})

	// Show all books in the library
	fmt.Println("Initial Library:")
	for _, book := range library {
		book.PrintInfo()
	}

	// Try borrowing a book
	var err error
	library, err = BorrowBook(library, "1984")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show library again after borrowing
	fmt.Println("\nLibrary After Borrowing:")
	for _, book := range library {
		book.PrintInfo()
	}
}
