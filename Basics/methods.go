package main

import "fmt"

type rect struct {
    width, height int
}
// This area method has a receiver type of *rect.
func (r *rect) area() int {
    return r.width * r.height
}
// Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

type geometricForm struct {
    area int
    rectangle rect // This is Inheritance in Go
}

func methodsFun() {

	fmt.Println("\n-----METHODS-----")

    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())

    g := geometricForm{
        12,
        r,
    }

    fmt.Println(g.area)
    fmt.Println("perim: ", g.rectangle.perim())
}