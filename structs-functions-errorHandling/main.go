package main

import (
	"errors"
	"fmt"
)

type Movie struct {
	Title    string
	Director string
	Year     int
	Rented   bool
	Rating   float64
}


func (m *Movie) Rent() error {
	if m.Rented {
		return errors.New("movie already rented")
	}
	m.Rented = true
	fmt.Println("You rented:", m.Title)
	return nil
}


func (m *Movie) Return() error {
	if !m.Rented {
		return errors.New("movie was not rented")
	}
	m.Rented = false
	fmt.Println("You returned:", m.Title)
	return nil
}


func SwapDirectors(m1, m2 *Movie) {
	temp := m1.Director
	m1.Director = m2.Director
	m2.Director = temp
}


func FindHighestRatedMovie(movies []Movie) (Movie, error) {
	if len(movies) == 0 {
		return Movie{}, errors.New("no movies available")
	}

	highest := movies[0]
	for _, m := range movies {
		if m.Rating > highest.Rating {
			highest = m
		}
	}
	return highest, nil
}

func main() {
	
	movie1 := Movie{"Inception", "Christopher Nolan", 2010, false, 8.8}
	movie2 := Movie{"Interstellar", "Christopher Nolan", 2014, false, 8.6}
	movie3 := Movie{"The Dark Knight", "Christopher Nolan", 2008, false, 9.0}

	
	if err := movie1.Rent(); err != nil {
		fmt.Println("Error:", err)
	}

	
	if err := movie1.Rent(); err != nil {
		fmt.Println("Error:", err)
	}

	
	if err := movie1.Return(); err != nil {
		fmt.Println("Error:", err)
	}

	
	SwapDirectors(&movie1, &movie2)
	fmt.Printf("After swap: movie1 director = %s, movie2 director = %s\n", movie1.Director, movie2.Director)


	movies := []Movie{movie1, movie2, movie3}
	best, err := FindHighestRatedMovie(movies)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Highest rated movie: %s (%.1f)\n", best.Title, best.Rating)
	}
}
