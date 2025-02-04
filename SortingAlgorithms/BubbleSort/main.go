package main

import "fmt"

/*
Bubble sort is an algorithm that sort an array from the lowest to the highest value.
It works by comparing each element in the array with the next element,
if the current element is greater than the next element, it swaps the two elements.
The algorithm continues to compare and swap the elements until no more swaps are needed.

Time complexity: O(n^2)
Space complexity: O(1)
*/

func main() {
	var n = []int{100, 39, 2, 9, 7, 54, 11}

	/*
	if len(n) {
		return n
	}
	*/

	fmt.Print("Unsorted Array: ")
	fmt.Println(n)

	var isDone = false

	for !isDone {
		isDone = true
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i] // swap line
				fmt.Printf("Swap %v with %v -->",n[i], n[i+1])
				fmt.Println(n)
				isDone = false
			}
		}
	}

	fmt.Println("Bubble Sort:")
	fmt.Println(n)
}
