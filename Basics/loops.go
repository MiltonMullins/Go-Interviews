package main

import "fmt"

/*
Unlike other languages like C, Java, or JavaScript
there are no parentheses surrounding the three components
of the for statement and the braces { } are always required.
*/
func loops() {

	fmt.Println("\n-----LOOPS-----")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}
	//You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 { // skip even numbers
			continue
		}
		fmt.Println(n)
	}

	//Basic for-each loop (slice or array)
	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
	}

	//String Iteration: runes or bytes
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}

	//To loop over individual bytes, simply use a normal for loop and string indexing:
	const s = "日本語"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	//Map Iteration -> Keys and Values
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	//Channel iteration
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
