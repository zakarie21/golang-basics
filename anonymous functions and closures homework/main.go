package main

import "fmt"

func main() {

    // 38. Anonymous function that squares a number and call it immediately
    // Answer: We define a function without a name and call it immediately
    number := 5
    squareFunction := func(n int) int {
        result := n * n
        return result
    }
    squaredNumber := squareFunction(number)
    fmt.Println("38:", squaredNumber) // 25

    // 39. Explain what a closure is in Go and give an example
    // Answer: A closure is a function that remembers variables from its outside scope
    greeting := "Hello"
    greetFunction := func(name string) {
        message := greeting + " " + name
        fmt.Println(message)
    }
    greetFunction("Alice") // prints "Hello Alice"

    // 40. Closure to generate the next Fibonacci number each time
    fibClosure := func() func() int {
        first := 0
        second := 1
        return func() int {
            next := first
            temp := first + second
            first = second
            second = temp
            return next
        }
    }
    nextFib := fibClosure()
    fmt.Println("40:", nextFib()) // 0
    fmt.Println("40:", nextFib()) // 1
    fmt.Println("40:", nextFib()) // 1
    fmt.Println("40:", nextFib()) // 2

    // 41. Closures can be dangerous in loops with goroutines
    // Answer: The closure captures the loop variable so all goroutines may use the same value
    fmt.Println("41: Demonstrating closure bug in loop")
    for i := 1; i <= 3; i++ {
        current := i // extra step to fix value for this iteration
        go func() {
            fmt.Println(current) // without 'current', all goroutines might print 4
        }()
    }

    // 42. Closure that maintains a running total
    runningTotal := func() func(int) int {
        total := 0
        return func(n int) int {
            total = total + n
            return total
        }
    }
    addToTotal := runningTotal()
    total1 := addToTotal(5)
    fmt.Println("42:", total1) // 5
    total2 := addToTotal(3)
    fmt.Println("42:", total2) // 8
    total3 := addToTotal(10)
    fmt.Println("42:", total3) // 18

    // 43. Can closures capture and modify variables outside their own scope
    // Answer: Yes, they can change outer variables
    count := 0
    increment := func() {
        count = count + 1
    }
    increment()
    increment()
    fmt.Println("43:", count) // 2

    // 44. Closure-based counter that can be reset to zero
    counterClosure := func() func(string) int {
        value := 0
        return func(action string) int {
            if action == "reset" {
                value = 0
            } else {
                value = value + 1
            }
            return value
        }
    }
    myCounter := counterClosure()
    first := myCounter("")      // 1
    fmt.Println("44:", first)
    second := myCounter("")     // 2
    fmt.Println("44:", second)
    reset := myCounter("reset") // 0
    fmt.Println("44:", reset)
    third := myCounter("")      // 1
    fmt.Println("44:", third)
}
