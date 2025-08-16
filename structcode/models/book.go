package models

import (
	"fmt"
	"encoding/json"
	"os"
)



type Book struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Todo Todo `json:"todostruct"`
}

	

func (b *Book) NewBook(title string, description string, todo *Todo){

  
	b.Title = title
	b.Description = description
	b.Todo = *todo

}

func CreateJSON(newBook Book) {

	
	byteConverterBookToJson, _ := json.Marshal(newBook)
	err:= os.WriteFile("newbook.txt", []byte(byteConverterBookToJson), 0644)
	
	if err != nil{
		fmt.Println(err)
	}
	newBook.Title = "fad"
}