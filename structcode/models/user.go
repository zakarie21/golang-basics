package models

import (
	"encoding/json"
	"fmt"
	"os"

)

// Create a user profile struct, all methods shou,ld be struct methods ,
// read and write to it. ALso delete the username as one of the method and see if it is eorking or not . THere should be no return in any of the struct method

type User struct {
	Name string `json:"title"`
	ID string `json:"description"`
	
}

	

func (u *User) NewUser(name string, id string){

  
	u.Name = name
	u.ID = id
	

}


func (u *User) DeleteUser(name string){

	if name == u.Name{
		u.Name = ""
		u.ID = ""
	}
}


func (u *User) CreateJSONUser() {

	contents, err := os.ReadFile("newuser.txt")

	if err != nil {
		panic("Error when creating contentns of user file")
	}

	byteConverterUserToJson, _ := json.Marshal(u)
	err = os.WriteFile("newuser.txt", []byte(byteConverterUserToJson) + []byte(contents), 0644)
	
	if err != nil{
		fmt.Println(err)
	}
	
}