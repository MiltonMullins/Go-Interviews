package main

import "fmt"

/*
Insertion Sort: It builds the final sorted list one item at a time.
It takes an element from the unsorted portion and inserts it into its correct position in the sorted portion, shifting the larger elements to the right.

Time complexity: O(n^2)
Space complexity: O(1)
*/

func main() {
	var n = []int{10, 39, 2, 9, 7, 54, 11}

	if len(n) == 1 {
		fmt.Println(n)
	}

	fmt.Print("Unsorted Array: ")
	fmt.Println(n)

	i := 1
	for i < len(n) {
		j := i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]
			fmt.Printf("Swap %v with %v -->", n[j-1], n[j])
			fmt.Println(n)
			j--
		}
		i++
	}

	fmt.Println(n)
}
