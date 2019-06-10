package main

import "fmt"

func main() {

	//init a slice

	nums := []int{3, 4, 6, 8, 9, 777, 33}

	//loop through nums
	for i, id := range nums {
		fmt.Printf("%d -ID: %d\n", i, id)
	}

	//not using index
	for _, id := range nums {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range nums {
		sum += id
	}

	fmt.Println("sum is :", sum)

	//range with map
	teams := map[string]string{"walcott": "Everton", "luca jovic": "Real Madrid"}

	for k, v := range teams {
		fmt.Printf("%s : %s", k, v)
	}
}
