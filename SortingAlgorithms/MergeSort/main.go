package main

import "fmt"

/*
Merge sort is a divide-and-conquer sorting algorithm that works by recursively dividing the array into smaller halves,
sorting them individually, and then merging the sorted halves back together to obtain the final sorted array.

Time complexity: O(n log n)
*/

func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11, 8}

	fmt.Println(mergeSort(n))
}

func mergeSort(array []int) []int {
	length := len(array)
	if length <= 1 {
		return array
	}

	mid := length / 2
	fmt.Print("Left: ")
	fmt.Println(array[:mid])
	fmt.Print("Right: ")
	fmt.Println(array[mid:])
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			fmt.Print("Merged left: ")
			fmt.Println(merged)
			i++
		} else {
			merged = append(merged, right[j])
			fmt.Print("Merged right: ")
			fmt.Println(merged)
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	fmt.Print("Merged total: ")
	fmt.Println(merged)

	return merged
}
