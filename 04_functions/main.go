package main

import "fmt"

func printname(name string) string {

	return name
}

func getsum(int1 int, int2 int) int {

	return int1 + int2
}

func main() {

	fmt.Println(printname("keveaux"))
	fmt.Println(getsum(3, 3))

}
