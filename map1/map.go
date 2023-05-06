package map1

import "fmt"

func Map() {

	// var myMap map[string]int

	// myMap = map[string]int{}

	// myMap["Javascript"] = 10
	// myMap["Golang"] = 12
	// myMap["Java"] = 13
	// myMap["Laravel"] = 15

	// fmt.Println(myMap)

	// fmt.Println(myMap["Java"])

	myMap := map[string]string{
		"Javascript": "10",
		"Golang":     "12",
		"Java":       "13",
	}

	fmt.Println(myMap)

}
