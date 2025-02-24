package main

import "fmt"

type User struct {
	name string
	lastName string
}

func (u User) showName(){
	fmt.Println("Name: " + u.name)
	fmt.Println("Last Name: " + u.lastName)
}

type Teacher struct {
	salary float32
	User
}

func (t Teacher) show(){
	t.showName()
	fmt.Printf("Salary: %v\n", t.salary)
}

type Student struct {
	average float32
	User
}

func (s Student) show(){
	s.showName()
	fmt.Printf("average: %v\n", s.average)
}

func main(){

	teacher := Teacher{salary: 1000, User: User{name: "John", lastName: "Doe"}}
	teacher.show()

	student := Student{average: 8.5, User: User{name: "Alice", lastName: "Smith"}}
	student.show()
}