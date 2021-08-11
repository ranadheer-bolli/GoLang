package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	//	p.name = "Ranadheer Bolli"
	p := person{name: name}
	p.age = 22
	return &p
}

func main() {

	fmt.Println(person{"Ranadheer", 21})

	fmt.Println(person{name: "javeed", age: 22})

	fmt.Println(person{name: "pavan"})

	fmt.Println(&person{"nikhil", 21})

	fmt.Println(newPerson("Faisal"))

	s := person{name: "praneeth", age: 21}

	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 22
	fmt.Println(sp.age)

}
