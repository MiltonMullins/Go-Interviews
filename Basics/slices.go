package main

/*
	Allocation with make
	Back to allocation. The built-in function make(T, args) serves a purpose different from new(T).
	It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).
	The reason for the distinction is that these three types represent, under the covers,
	references to data structures that must be initialized before use. A slice, for example,
	is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity,
	and until those items are initialized, the slice is nil. For slices, maps, and channels,
	make initializes the internal data structure and prepares the value for use.
*/
import (
	"fmt"
	"slices"
)

func slicesFun() {

	fmt.Println("\n-----SLICES-----")

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	/*
		To create an empty slice with non-zero length, use the builtin make.
		Here we make a slice of strings of length 3 (initially zero-valued).
		By default a new slice’s capacity is equal to its length;
		if we know the slice is going to grow ahead of time,
		it’s possible to pass a capacity explicitly as an additional parameter to make.
	*/
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	/*
		In addition to these basic operations, slices support several more that make them richer than arrays.
		One is the builtin append, which returns a slice containing one or more new values.
		Note that we need to accept a return value from append as we may get a new slice value.
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	/*
		Slices support a “slice” operator with the syntax slice[low:high].
		For example, this gets a slice of the elements s[2], s[3], and s[4].
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)

	//This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2:", l)

	//And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
