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

func miniMaxSumV2(arr []int32) {
	var sum, min, max int32
	for _, v := range arr {
		sum += v
	}
	min = arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		} else if arr[i] > max {
			max = arr[i]
		}
	}
	totalX := sum - min
	totalY := sum - max
	fmt.Println(totalY, totalX)
}

func main() {

	ar := []int32{1,3,5,7,9}
	miniMaxSum(ar)
}
