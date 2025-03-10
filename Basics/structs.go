package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

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

	fmt.Println("\n-----STRUCTS-----")

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

	fmt.Println("SORTING ARRAY OF STRUCT")
	audience := []person{
		{"Alice", 35},
		{"Bob", 45},
		{"James", 25},
	}
	// Sorting slices of structs - using the sort interface.
	sort.Sort(AgeFactor(audience))
	fmt.Println(audience)

	// Sorting slices of structs - using the SortFunc method from the slices package.
	slices.SortFunc(audience,
		func(a, b person) int {
			return cmp.Compare(a.age, b.age)
		})


	//IMPLEMENTING DEPENDENCY INVERSION PRINCIPLE
	Employee1 := Employee{
		emailNotif: EmailNotification{},
	}
	Employee1.emailNotif.notify("Test")
}

type AgeFactor []person

func (a AgeFactor) Len() int {
	return len(a)
}

func (a AgeFactor) Less(i, j int) bool {
	return a[i].age < a[j].age
}
func (a AgeFactor) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}


//IMPLEMENTING DEPENDENCY INVERSION PRINCIPLE
//The Dependency Inversion Principle states that high-level modules should not depend on low-level modules.
//Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.
//In the example below, the Employee struct depends on the NotificationService interface.
//The EmailNotification and SlackNotification structs implement the NotificationService interface.
//The Employee struct can now use any type that implements the NotificationService interface.
//This allows for easy swapping of the notification service without changing the Employee struct.
type NotificationService interface {
	notify(msg string)
}

type EmailNotification struct {}

func (EmailNotification) notify(msg string) {
	fmt.Println("Email notification: ", msg)
}

type SlackNotification struct {}

func (SlackNotification) notify(msg string) {
	fmt.Println("Slack notification: ", msg)
}

type Employee struct {
	emailNotif NotificationService
}

