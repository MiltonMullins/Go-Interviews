package main

import (
	"fmt"
	"slices"
)

func miniMaxSum(arr []int32) {
	var min,max int64

	slices.Sort(arr)

	for i := 0 ; i < len(arr)-1 ; i++ {
		min+= int64(arr[i])
	}

	for i := 1 ; i < len(arr) ; i++ {
		max+= int64(arr[i])
	}
	
	fmt.Printf("%v %v", min, max)
}

func main() {

	ar := []int32{1,3,5,7,9}
	miniMaxSum(ar)
}
