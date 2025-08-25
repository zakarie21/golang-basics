package main

import (
	"fmt"
	"time"
	"goroutines/models"

)

func SlowGreet(name string, done chan string) {
	//time.Sleep(5 * time.Second)
	fmt.Println("Slow greet", name)
	done <- "slow greet complete"
}

func FastGreet(name string, done chan string) {
	time.Sleep(3 * time.Second)
	fmt.Println("fast greet", name)
	done <- "fast greet complete"
}

func main() {

	// done:= make(chan string)

	// go SlowGreet("Lucy", done)
	// go SlowGreet("Lucy", done)
	// go SlowGreet("Lucy", done)
	// go SlowGreet("Lucy", done)
	// fmt.Println("edws")
	// go  FastGreet("Jack", done)

	// for i := 0; i < 5; i++ {
	// 	<- done
	// }

	done := make(chan string)
	err := make(chan error)

	orderStatusObject:= models.Order{
		OrderID: "123456zak7890",
	}

	go orderStatusObject.IsShippingAvailable(done, err)
	go orderStatusObject.IsProductAvailable(done, err)
	go orderStatusObject.IsPaymentSuccessful(done, err)

	for i := 0; i < 3; i++ {

		select {
		case success := <-done:
			{
				fmt.Println(success)
			}
		case failure := <-err:
			{
				fmt.Println(failure)
			}
		}
	}

}
