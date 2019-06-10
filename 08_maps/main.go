package main

import "fmt"

func main() {

	//maps-this is a key value pair

	emails := make(map[string]string)

	//assign key value

	emails["kevin"] = "kevmirera@gmail.com"

	emails["sharon"] = "sharon@gmail.com"

	fmt.Println(emails["kevin"])

	//delete a value
	delete(emails, "kevin")

	fmt.Println(emails)

	//declare map and add kev value

	teams := map[string]string{"walcott": "Everton", "luca jovic": "Real Madrid"}

	fmt.Println(teams)

}
