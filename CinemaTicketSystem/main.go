package main

import (
	"errors"
	"fmt"
)

/*
Questions:
1. Create a Movie struct with Title, TicketsAvailable, and TicketPrice
2. Make a slice of Movie called cinema
3. Write AddMovie function to add movies to the cinema
4. Write BuyTicket function that reduces TicketsAvailable or returns an error if sold out
5. Write ShowRevenue function that loops through all movies and calculates total revenue
6. Write FindMovie function that prints movie details or an error if not found
7. CountMovies: return the total number of movies in the cinema
8. HasAvailableTickets: return true if at least one movie has tickets available
9. GetFirstAvailableMovie: return the first movie with available tickets or an error if none
10. CountAvailableTickets: return the total number of available tickets across all movies
*/

type Movie struct {
	Title            string
	TicketsAvailable int
	TicketPrice      float64
}

func AddMovie(cinema []Movie, movie Movie) []Movie {
	cinema = append(cinema, movie)
	return cinema
}

func BuyTicket(cinema []Movie, title string, tickets int) ([]Movie, error) {
	for i := 0; i < len(cinema); i++ {
		if cinema[i].Title == title {
			if cinema[i].TicketsAvailable >= tickets {
				cinema[i].TicketsAvailable = cinema[i].TicketsAvailable - tickets
				return cinema, nil
			}
			return cinema, errors.New("not enough tickets available")
		}
	}
	return cinema, errors.New("movie not found")
}

func ShowRevenue(cinema []Movie, sold map[string]int) {
	totalRevenue := 0.0
	for _, movie := range cinema {
		revenue := float64(sold[movie.Title]) * movie.TicketPrice
		fmt.Println(movie.Title, "revenue:", revenue)
		totalRevenue = totalRevenue + revenue
	}
	fmt.Println("Total revenue:", totalRevenue)
}

func FindMovie(cinema []Movie, title string) error {
	for _, movie := range cinema {
		if movie.Title == title {
			fmt.Println("\nMovie Found:")
			fmt.Println("Title:", movie.Title)
			fmt.Println("Tickets Available:", movie.TicketsAvailable)
			fmt.Println("Ticket Price:", movie.TicketPrice)
			return nil
		}
	}
	return errors.New("movie not found in cinema")
}

// Question 7
func CountMovies(cinema []Movie) int {
	return len(cinema)
}

// Question 8
func HasAvailableTickets(cinema []Movie) bool {
	for _, movie := range cinema {
		if movie.TicketsAvailable > 0 {
			return true
		}
	}
	return false
}

// Question 9
func GetFirstAvailableMovie(cinema []Movie) (*Movie, error) {
	for i := range cinema {
		if cinema[i].TicketsAvailable > 0 {
			return &cinema[i], nil
		}
	}
	return nil, errors.New("no movies with available tickets")
}

// Question 10
func CountAvailableTickets(cinema []Movie) int {
	total := 0
	for _, movie := range cinema {
		total += movie.TicketsAvailable
	}
	return total
}

func main() {
	// Step 1: Create an empty cinema (slice of movies)
	cinema := []Movie{}

	// Step 2: Add movies
	cinema = AddMovie(cinema, Movie{Title: "Inception", TicketsAvailable: 50, TicketPrice: 10.0})
	cinema = AddMovie(cinema, Movie{Title: "Interstellar", TicketsAvailable: 40, TicketPrice: 12.5})
	cinema = AddMovie(cinema, Movie{Title: "Tenet", TicketsAvailable: 30, TicketPrice: 11.0})

	// Step 3: Create a map to track tickets sold
	sold := make(map[string]int)

	// Step 4: Try buying some tickets
	var err error
	cinema, err = BuyTicket(cinema, "Inception", 3)
	if err == nil {
		sold["Inception"] = sold["Inception"] + 3
	} else {
		fmt.Println("Error:", err)
	}

	cinema, err = BuyTicket(cinema, "Tenet", 5)
	if err == nil {
		sold["Tenet"] = sold["Tenet"] + 5
	} else {
		fmt.Println("Error:", err)
	}

	// Step 5: Show total revenue
	fmt.Println("\nCinema Revenue:")
	ShowRevenue(cinema, sold)

	// Step 6: Find a movie in the cinema
	fmt.Println()
	err = FindMovie(cinema, "Interstellar")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = FindMovie(cinema, "Avatar")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Demo Question 7
	fmt.Println("\nTotal movies in cinema:", CountMovies(cinema))

	// Demo Question 8
	fmt.Println("Any tickets available?", HasAvailableTickets(cinema))

	// Demo Question 9
	firstMovie, err := GetFirstAvailableMovie(cinema)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("First available movie:", firstMovie.Title)
	}

	// Demo Question 10
	fmt.Println("Total available tickets:", CountAvailableTickets(cinema))
}
