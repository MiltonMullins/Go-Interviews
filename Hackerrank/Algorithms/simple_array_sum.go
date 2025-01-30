package main

import (
	"fmt"
)

func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0

/* 	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	} */

	for _, v := range ar {
		sum += v
	}
	return sum
}

func main() {

	ar := []int32{1,2,3,4,5}
	fmt.Println(simpleArraySum(ar))
}
