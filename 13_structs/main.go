package main

import (
	"fmt"
	"strconv"
)

//structs are like classes
//define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//there are two types of methods : value and pointer receivers

//value receiver method-you cannot change values in a value receiver method

func (person Person) greet() string {

	return "Hello " + person.firstName + person.lastName + " aged " + strconv.Itoa(person.age)
}

//pointer receiver method- you can change something in a pointer receiver

func (person *Person) changeAge() {
	person.age++
}

func (person *Person) getMarried(spouseLastName string) {
	person.lastName = spouseLastName
}

func main() {

	person1 := Person{firstName: "kevin", lastName: "mirera", city: "Nairobi", gender: "m", age: 23}

	//alternative
	// person1 := Person{ "kevin",  "mirera",  "Nairobi",  "m",  23}

	// fmt.Println(person1.gender)

	person1.changeAge()
	person1.getMarried("Aniston")
	fmt.Println(person1.greet())

}
