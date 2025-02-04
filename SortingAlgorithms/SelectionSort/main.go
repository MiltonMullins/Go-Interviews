package main

import "fmt"

/*
The Selection Sort algorithm finds the lowest value in an array and moves it to the front of the array.

Time complexity: O(n^2)
Space complexity: O(1)
*/

func main(){

	var n = []int{1, 39, 2, 9, 7, 54, 11}

	if len(n) == 1 {
		fmt.Println(n)
	}

	fmt.Print("Unsorted Array: ")
	fmt.Println(n)


	for i := 0; i < len(n); i++ {
		min := i
		for j := i + 1; j < len(n); j++ {
		 if n[min] > n[j] {
		  min = j
		 }
		}
		n[min],n[i] = n[i],n[min]
		fmt.Printf("Swap %v with %v -->", n[i], n[min])
		fmt.Println(n)
	   }
	   fmt.Println("The selection sorted array is : ")
	   fmt.Println(n)
}