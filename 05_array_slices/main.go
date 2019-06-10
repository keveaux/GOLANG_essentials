package main

import "fmt"

func main() {

	//Arrays -have to be a fixed length
	//slice-array without a fixed type

	// var fruitArr [2]string

	// //assign values

	// fruitArr[0] = "apples"
	// fruitArr[1] = "oranges"

	//declare and assign at the same time
	fruitArr := [2]string{"apples", "oranges"}

	fmt.Println(fruitArr)

	fruitArrSlice := []string{"coconuts", "pineapples"}

	fmt.Println(len(fruitArrSlice))

}
