/*
1 - Create a map that contains the number of the months as key and a string that contains the name of the months as value.
2 - Create a slice of 20 random numbers between 0 and 20.
3 - Spawn 20 go routines to call the function of the next item.
4 - Create a function to check if a number exist on the map. The number will be from the slice of item 2.
	a - if the number is equal to a key, print the name of the month.
	b - if not, send those numbers on a channel and print them on screen.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	//Crear map
	calendar := map[int]string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "May",
		6:  "Jun",
		7:  "Jul",
		8:  "Agu",
		9:  "Sep",
		10: "Oct",
		11: "Nov",
		12: "Dec",
	}

	//Crear Slice
	var randomNum = make([]int, 20)

	for i := 0; i < 20; i++ {
		randomNum[i] = rand.Intn(20)
	}

	fmt.Println(randomNum)
	fmt.Println(calendar)

	//Crear Funcion
	ch := make(chan int, 20)

	var wg sync.WaitGroup
	//Crear goroutines
	for _, num := range randomNum {
		wg.Add(1)
		go checkCalendar(num, calendar, ch, &wg)
	}
	wg.Wait()
	close(ch)

	for value := range ch {
		println(value)
	}
}

func checkCalendar(num int, calendar map[int]string, ch chan int, wg *sync.WaitGroup) {
	if month, exist := calendar[num]; exist {
		fmt.Println(month)
	} else {
		ch <- num
	}
	wg.Done()
}
