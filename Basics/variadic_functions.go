package main

//Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func variadicFun() {

	fmt.Println("\n-----VARIADIC FUNCTION-----")

    sum(1, 2)
    sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}