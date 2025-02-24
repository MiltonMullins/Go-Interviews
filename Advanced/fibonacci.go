package main

import (
	"fmt"
	//"time"
)

/* func main() {

	slice := []int{1, 2, 3, 4, 5, 6}

	wg := sync.WaitGroup{}
	ch := make(chan int, len(slice))
	for _, v := range slice {
		wg.Add(1)
		go workker(v, &wg, ch)
	}

	wg.Wait()
	close(ch)
	fmt.Println("Hello World")

/* 	for v := range ch {
		fmt.Println(v)
	}

	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}

}

func workker(num int, wg *sync.WaitGroup, ch chan int) {

	fmt.Println(num)
	time.Sleep(time.Second)
	ch <- num
	wg.Done()
}
*/

 func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	//time.Sleep(10*time.Second)
	fibonacci(c, quit)
}