package main

import "fmt"

//
// 1. Account struct and function to add 50 to balance
//
type Account struct {
    owner   string
    balance int
}

// function takes a pointer, so it changes the original account
func addBonus(account *Account) {
    account.balance = account.balance + 50
}

//
// 2. Player struct and function to update score
//
type Player struct {
    username string
    score    int
}

// function takes a pointer, so the score gets updated in the original struct
func updateScore(player *Player, newScore int) {
    player.score = newScore
}

//
// 3. Book struct with pointer receiver method
//
type Book struct {
    title string
    pages int
}

// method with pointer receiver so I can change pages inside the original book
func (book *Book) addPages() {
    book.pages = book.pages + 10
}

//
// 4. Teacher struct with two pointers to the same struct
//
type Teacher struct {
    name    string
    subject string
}

//
// 5. Function returning a pointer to a Movie struct
//
type Movie struct {
    name   string
    rating int
}

func newMovie(name string, rating int) *Movie {
    return &Movie{name: name, rating: rating}
}

func main() {
    // Q1
    myAccount := Account{owner: "Sam", balance: 100}
    fmt.Println("Before bonus:", myAccount.balance)
    addBonus(&myAccount) 
    // I pass a pointer so the function can edit the balance
    fmt.Println("After bonus:", myAccount.balance)

    // Q2
    myPlayer := Player{username: "Gamer123", score: 10}
    fmt.Println("Before score update:", myPlayer.score)
    updateScore(&myPlayer, 99)
    // I pass a pointer and the new score I want
    fmt.Println("After score update:", myPlayer.score)

    // Q3
    myBook := Book{title: "Learning Go", pages: 200}
    fmt.Println("Pages before:", myBook.pages)
    myBook.addPages() 
    // calling the method adds 10 pages
    fmt.Println("Pages after:", myBook.pages)

    // Q4
    teacher1 := &Teacher{name: "Mr. Lee", subject: "History"}
    teacher2 := teacher1 
    // teacher2 points to the same Teacher as teacher1
    fmt.Println("teacher1 subject:", teacher1.subject)
    fmt.Println("teacher2 subject:", teacher2.subject)
    teacher1.subject = "Math" 
    // I change subject through teacher1
    fmt.Println("teacher2 subject after change:", teacher2.subject)
    // teacher2 shows the change because they point to the same struct

    // Q5
    movie1 := newMovie("Inception", 9)
    movie2 := newMovie("Interstellar", 10)
    // function returned pointers to two different Movie structs
    fmt.Println("Movie1:", movie1.name, "Rating:", movie1.rating)
    fmt.Println("Movie2:", movie2.name, "Rating:", movie2.rating)
}
