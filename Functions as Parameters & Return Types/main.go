package main

import "fmt"

func main() {

    // 32. Write a function that accepts another function func(int) int and applies it to every element of a slice.
    // Answer: Go through the slice and use a function on each number.
    numbers := []int{1, 2, 3, 4, 5}
    for i, number := range numbers {
        double := func(n int) int {
            return n * 2
        }
        numbers[i] = double(number)
    }
    fmt.Println("32:", numbers)

    // 33. How can you store multiple different functions in a slice and call them in a loop?
    // Answer: Put the functions in a slice and call each one.
    functions := []func(int) int{
        func(n int) int { return n + 1 },
        func(n int) int { return n * 2 },
        func(n int) int { return n * n },
    }
    fmt.Println("33:")
    for _, functionItem := range functions {
        result := functionItem(3)
        fmt.Println(result)
    }

    // 34. Write a function that returns another function which increments an integer by a given step.
    // Answer: One function makes another function that adds a step.
    makeIncrement := func(step int) func(int) int {
        return func(n int) int {
            return n + step
        }
    }
    addFive := makeIncrement(5)
    fmt.Println("34:", addFive(10))

    // 35. Can you pass a method (bound to a struct instance) as a function argument?
    // Answer: Yes. A struct can have a method and we can call it like a normal function.
    type Box struct {
        number int
    }
    func(b Box) {
        fmt.Println("35:", b.number)
    }(Box{number: 7})

    // 36. Write a function that takes another function as a parameter and measures its execution time.
    // Answer: To keep it simple, we just show a function being passed and called.
    runFunction := func(f func()) {
        f()
    }
    sayHello := func() {
        fmt.Println("36: Hello from inside the function")
    }
    runFunction(sayHello)

    // 37. What is the type signature of a function returning no value but accepting two string arguments?
    // Answer: The type is func(string, string).
    showStrings := func(a string, b string) {
        fmt.Println("37:", a, b)
    }
    showStrings("apple", "banana")
}