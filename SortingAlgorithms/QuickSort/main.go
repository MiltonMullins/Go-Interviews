package main

import "fmt"

/*
Quick sort is another efficient divide-and-conquer sorting algorithm.
It works by selecting a pivot element from the array and partitioning the other elements into two sub-arrays based on
whether they are less than or greater than the pivot. The sub-arrays are then recursively sorted using the same process.
*/

func main() {
	var n = []int{10, 39, 2, 9, 7, 54, 11}

	fmt.Println(quickSort(n))
}

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	fmt.Print("Array: ")
	fmt.Println(array)

	pivot := array[len(array)-1]
	fmt.Printf("Pivot: %v\n", pivot)
	var left, right []int

	for i := 0; i < len(array)-1; i++ {
		if array[i] <= pivot {
			left = append(left, array[i])
			fmt.Printf("Left: %v\n", left)
		} else {
			right = append(right, array[i])
			fmt.Printf("Right: %v\n", right)
		}
	}

	left = quickSort(left)
	fmt.Printf("Left Total: %v\n", left)
	right = quickSort(right)
	fmt.Printf("Right Total: %v\n", right)

	return append(append(left, pivot), right...)
}
