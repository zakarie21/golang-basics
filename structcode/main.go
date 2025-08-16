package main

import (
	// "bufio"
	// "os"
	//"strings"
	// divide "structPointersNotes/divide"
	// book "structPointersNotes/models"
	// todo "structPointersNotes/models"
	user "structPointersNotes/models"


	"fmt"
)



func main(){
	// var newTodoPtr *todo.Todo
	// newTodoPtr = &todo.Todo{}
	// var newBookPtr *book.Book
	// newBookPtr = &book.Book{}
	var newUserPtr *user.User
	newUserPtr = &user.User{}
	var user2 *user.User
	user2 = &user.User{}

	// var todoPtr *todo.Todo
	// todoPtr= &todo.Todo{Description : "dinner"}
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("PLease enter the notes title")
	

	// description, err := reader.ReadString('\n')
	// if err != nil {
	// 	panic("description note could  be created")
	// }

	// newBookPtr.NewBook("Brad", "Breakfast", &todo.Todo{Description : "dinner"} )
	// fmt.Println(newBookPtr)
	// newTodoPtr.New(description)
	// todo.CreateTodoJSON(newTodoPtr)


	newUserPtr.NewUser("Zak", "99" )
	fmt.Println(newUserPtr)

	user2.NewUser("Jack", "21" )
	fmt.Println(user2)

	newUserPtr.CreateJSONUser()

	newUserPtr.DeleteUser("Zak")
	
}

// Create a user profile struct, all methods shou,ld be struct methods 
// , read and write to it. ALso delete the username as one of the method and see if it is eorking or not . 
// THere should be no return in any of the struct method