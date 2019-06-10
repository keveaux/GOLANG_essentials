package main

import "fmt"

func main() {

	var name = "keveaux"
	var age = 23

	//constants - cannot be redefined
	const iscool = true

	//shorthand way of defining variables can only be used inside a function
	date := 123

	fmt.Println(name, age, iscool, date)
	fmt.Printf("%T\n", age)

}
