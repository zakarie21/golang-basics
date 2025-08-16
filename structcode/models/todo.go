package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct{
	Description string `json:"description"`
}

func (t *Todo) New(description string) {
	t.Description = description
}

func CreateTodoJSON(printTodo *Todo) {

	byteConvertTodoToJSON,_:= json.Marshal(printTodo)
	err:= os.WriteFile("newtodo.txt", []byte(byteConvertTodoToJSON), 0644)

	if err!= nil{
		fmt.Println(err)
	}
}