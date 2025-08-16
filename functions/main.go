package main


//create a function that takes a function as an argument and return the function as an argument
//execute function your getting as an argument on each value of the num array you're receving as an argmument

import "fmt"


func main() {

	numbers:= []int{20,40,60,80}

	fnNewNumbers, result:= TakeArgument(Add, numbers) 

	

	fmt.Println(fnNewNumbers(3))
	fmt.Println(result)

}

func Add(num int) int {
	return num + 5
}

func TakeArgument(addfunction func(int) int, num []int) (func(int) int, []int){

	for index, value := range num{
		num[index] = addfunction(value)
	}

	return func(multipliedValue int) int{
		return multipliedValue*3
	} ,num

}

