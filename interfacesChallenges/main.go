package main

import (
	"fmt"
	models "interface/models"
)


// type Shape interface{
// 	Perimeter() float64
// }

type PaymentProcessor interface{
	Pay(amount float64) string
}

func Checkout(p PaymentProcessor, amount float64){
	fmt.Println(p.Pay(amount))
}

func main() {
	// var circle Shape
	// var rectangle Shape

	creditCard := models.CreditCard{CardNumber: "12345678"}
	upi := models.Upi{UpiNumber: "45678"}
	crypto:= models.CryptoWallet{ID: "111111"}

	
	Checkout(creditCard, 456.78)
	Checkout(upi, 2340.83)
	Checkout(crypto, 234500.76)



	// rectangle = models.Rectangle{Width: 4.7, Height: 6.7}
	// circle = models.Circle{Radius: 4.7}
	// fmt.Println(rectangle.Perimeter())
	// fmt.Println(circle.Perimeter())
	
}

/*
Create a PaymentProcessor interface:
type PaymentProcessor interface {
    Pay(amount float64) string
}
Implement it for CreditCard, UPI, and CryptoWallet.
Write a Checkout function that accepts a PaymentProcessor 
and uses it without knowing the exact type.

*/