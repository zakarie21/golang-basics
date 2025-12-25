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
6. Write ReturnBook function to mark a borrowed book as available
7. CountAvailableBooks: return how many books are currently available
8. FindBookByTitle: return a pointer to a book by title or an error if not found
9. ListBorrowedBooks: return all books that are currently borrowed
10. CountBorrowedBooks: return how many books are currently borrowed
11. ListAvailableBooks: return all books that are currently available
12. GetTotalBooks: return the total number of books in the library
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

// Question 6
func ReturnBook(library []Book, title string) ([]Book, error) {
	for i := 0; i < len(library); i++ {
		if library[i].Title == title {
			if !library[i].Available {
				library[i].Available = true
				return library, nil
			}
			return library, errors.New("book was not borrowed")
		}
	}
	return library, errors.New("book not found")
}

// Question 7
func CountAvailableBooks(library []Book) int {
	count := 0
	for _, book := range library {
		if book.Available {
			count++
		}
	}
	return count
}

// Question 8
func FindBookByTitle(library []Book, title string) (*Book, error) {
	for i := range library {
		if library[i].Title == title {
			return &library[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// Question 9
func ListBorrowedBooks(library []Book) []Book {
	var borrowed []Book
	for _, book := range library {
		if !book.Available {
			borrowed = append(borrowed, book)
		}
	}
	return borrowed
}

// Question 10
func CountBorrowedBooks(library []Book) int {
	count := 0
	for _, book := range library {
		if !book.Available {
			count++
		}
	}
	return count
}

// Question 11
func ListAvailableBooks(library []Book) []Book {
	var available []Book
	for _, book := range library {
		if book.Available {
			available = append(available, book)
		}
	}
	return available
}

// Question 12
func GetTotalBooks(library []Book) int {
	return len(library)
}

func main() {
	// Create an empty library (slice of books)
	library := []Book{}

	// Add some books
	library = AddBook(library, Book{Title: "Lord of the Rings", Author: "Tolkien", Available: true})
	library = AddBook(library, Book{Title: "Harry Potter", Author: "JK Rowling", Available: true})

	// Show all books in the library
	fmt.Println("Initial Library:")
	for _, book := range library {
		book.PrintInfo()
	}

	// Try borrowing books
	var err error
	library, err = BorrowBook(library, "Lord of the Rings")
	if err != nil {
		fmt.Println("Error:", err)
	}
	library, err = BorrowBook(library, "Harry Potter")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show library after borrowing
	fmt.Println("\nLibrary After Borrowing:")
	for _, book := range library {
		book.PrintInfo()
	}

	// Return one book
	library, err = ReturnBook(library, "Harry Potter")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show library after returning
	fmt.Println("\nLibrary After Returning:")
	for _, book := range library {
		book.PrintInfo()
	}

	// Use CountAvailableBooks
	fmt.Println("\nAvailable books:", CountAvailableBooks(library))

	// Use FindBookByTitle
	foundBook, err := FindBookByTitle(library, "Lord of the Rings")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\nFound book:")
		foundBook.PrintInfo()
	}

	// Use ListBorrowedBooks
	fmt.Println("\nBorrowed books:")
	for _, book := range ListBorrowedBooks(library) {
		book.PrintInfo()
	}

	// Use CountBorrowedBooks
	fmt.Println("\nTotal borrowed books:", CountBorrowedBooks(library))

	// Use ListAvailableBooks
	fmt.Println("\nAvailable books list:")
	for _, book := range ListAvailableBooks(library) {
		book.PrintInfo()
	}

	// Use GetTotalBooks
	fmt.Println("\nTotal books in library:", GetTotalBooks(library))
}
