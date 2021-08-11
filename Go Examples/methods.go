package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) modify(name string, age int) {
	p.age = age
	p.name = name
}

func main() {

	person1 := person{name: "abc", age: 12}
	fmt.Println(person1.age, person1.name)

	person2 := &person1

	person1.modify("Rana", 21)

	fmt.Println(person1.age, person1.name)
	fmt.Println(person2.age, person2.name)
}
