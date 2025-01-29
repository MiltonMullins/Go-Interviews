package main

/* In Go, an array is a numbered sequence of elements of a specific length. 
In typical Go code, slices are much more common; arrays are useful in some special scenarios.*/

import "fmt"

func arrays() {

	fmt.Println("\n-----ARRAYS-----")

    var a [5]int
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	//You can also have the compiler count the number of elements for you with ...
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	//If you specify the index with :, the elements in between will be zeroed.
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}