package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

/*
	When you have multiple consecutive parameters of the same type,
	you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
*/
func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func functionsFun() {

	fmt.Println("\n-----FUNCTIONS AND MULTIPLE RETURNS-----")

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// Here we use the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// If you only want a subset of the returned values, use the blank identifier _.
	_, c := vals()
	fmt.Println(c)
}
