package main

import "fmt"

func main() {

	//pointers point to the memory address or location of a value

	a := 5
	b := &a

	fmt.Println(a, b)

	//use * to read val from address

	fmt.Println(*b)
}
