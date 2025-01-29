package main

import "fmt"

type person struct {
    name string
    age  int
}
/* 
	Go is a garbage collected language; you can safely return a pointer to a local variable - 
	it will only be cleaned up by the garbage collector when there are no active references to it.
 */
func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func structsFun() {

    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})

    fmt.Println(newPerson("Jon"))

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)

    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}