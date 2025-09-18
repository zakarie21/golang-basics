package main

import "fmt"

/*
Create a Book struct

Fields: Title (string), Author (string), Year (int), CheckedOut (bool)

Create a method CheckOut

Updates CheckedOut to true and prints "Book checked out"

Create a function SwapAuthors

Takes two *Book and swaps their Author fields

Create a function FindNewestBook

Takes a slice of Book and returns the one with the most recent Year

*/

type Book struct{
	Title string
	Author string
	Year int
	CheckedOut bool
}

func(b *Book) CheckOut(){
	if b.CheckedOut == true {
		fmt.Println("You have already checked out")
	} else if b.CheckedOut == false{
		b.CheckedOut = true
		fmt.Println("Book has now been updated to check out")
	}
}

func SwapAuthors(book1, book2 *Book) {
	temp:= book1.Author
	book1.Author = book2.Author
	book2.Author = temp
}

func FindNewestBook(books []Book) string{
	if len(books) == 0{
		fmt.Println("You have no books in your database")
	}

	newestBook:= books[0]
	

	for _, book:= range books{
		if book.Year > newestBook.Year{
			newestBook = book
		}
	}

	return newestBook.Title
}

func main() {

		book1:= Book{"Heights", "Brad", 2023, false}
		book2:= Book{"Soldiers", "Katie", 2025, true}

		book1.CheckOut()
		book2.CheckOut()

		fmt.Printf("Author of book 1 is: %s, and auhtor of book 2 is: %s\n", book1.Author, book2.Author)
		SwapAuthors(&book1,&book2)
		fmt.Printf("Author of book 1 is now: %s, and auhtor of book 2 is now: %s\n", book1.Author, book2.Author)

		check:= []Book{book1,book2}
		result:= FindNewestBook(check)
		fmt.Println("Most recent book is", result)
}


/*
Create a Book struct

Fields: Title (string), Author (string), Year (int), CheckedOut (bool)

Create a method CheckOut

Updates CheckedOut to true and prints "Book checked out"

Create a function SwapAuthors

Takes two *Book and swaps their Author fields

Create a function FindNewestBook

Takes a slice of Book and returns the one with the most recent Year

*/