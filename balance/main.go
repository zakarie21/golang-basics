// Create a banking app which will have 4 options,
// 1. ) View account balance
// 2.) withdrawal amount from the bank and the user will enter the amount to withdraw
// 3.) Credit amount to the balance
//  4.) Exit from the bank. The program should run infinitely until and unless user enters
//  4 . Handle it with error scenarios and using for loop for infinitely handling

package main

import (
	"errors"
	"fmt"
	"strconv"

	//"strconv"
	//"errors"
	// "strconv"
	"os"
)

/*
Create a banking app which will have 4 options, 1. ) View account balance
2.) withdrawal amount from the bank and the user will enter the amount to withdraw
3.) Credit amount to the balance 4.) Exit from the bank. The program should
run infinitely until and unless user enters
4 . Handle it with error scenarios and using for loop for infinitely handling
*/

func viewBalance(balance float64) {

	fmt.Println("Your  balance is", balance)
}

func withdrawAmount(balance float64, withdraw float64) float64 {

	balance = balance - withdraw

	fmt.Println("Your new balance is", balance)

	return balance
}

func creditAccount(balance float64, add float64) float64 {

	balance = balance + add
	fmt.Println("Your new balance is", balance)
	return balance
}


func SaveBalance(filename string, balance float64){
	
	//content := []byte("Hello, Go by Example!\nThis is a test file.")
	stringBalance := fmt.Sprint(balance)
	byteBalance := []byte(stringBalance)
	err := os.WriteFile(filename, byteBalance, 0644)
	if err != nil{
		fmt.Println("Write operation failed")
	}
}

func ReadBalance(filename string) float64{
	data, err := os.ReadFile("Balance.txt")
	if err != nil{
		fmt.Errorf("read file failure")
	}
	stringBalance:= string(data)
	
	floatBalance, err:= strconv.ParseFloat(stringBalance, 64)
	if err!= nil{
		errors.New("Float type conversion failure")
	}
	return floatBalance
	

}

func main() {

	var option int
	var balance float64
	var withdraw float64
	var add float64
	

	balance = ReadBalance("Balance.txt")

	for {
		fmt.Println("1 - View Balance")
		fmt.Println("2- Withdraw")
		fmt.Println("3 - Add credit")
		fmt.Println("4- Exit")
		fmt.Println("Please choose an option")
		fmt.Scan(&option)

		if option == 1 {
			viewBalance(balance)

		} else if option == 2 {
			fmt.Println("How much would you like to withdraw")
			fmt.Scan(&withdraw)
			balance = withdrawAmount(balance, withdraw)

		} else if option == 3 {
			fmt.Println("How much would you like to add")
			fmt.Scan(&add)
			balance = creditAccount(balance, add)

		} else if option == 4 {
			SaveBalance("Balance.txt", balance)
			break
		} else {
			fmt.Println("PLease enter a valid number")
		}
	}

	// //create variables to hold the data
	// var principleAmount string
	// var interest float64
	// var time int

	// //collect input from user

	// principleAmount = "20.000"
	// interest = 12.0
	// time = 10

	// // converting string to float64
	// // amount := 0.0
	// amount, err := strconv.ParseFloat(principleAmount, 64)
	// if err != nil {
	// 	fmt.Println("Invalid number")
	// }

	// // calculating the interest
	// result := (amount * interest * float64(time)) / 100
	// fmt.Println("The simple intrestest calculated is:", result)

}
