package map1

import "fmt"

func Map1() {
	myMap := map[string]string{
		"Javascript": "Is Populer Language ",
		"Golang":     "Is Fastest Language",
		"Java":       "Is protected Programming Language",
		"Ruby":       "Is Beatiful Language",
	}

	for key, value := range myMap {
		fmt.Println("Key :", key, "value:", value)
	}

	fmt.Println("=====================================================")

	delete(myMap, "Java")

	for key, value := range myMap {
		fmt.Println("Key :", key, "value:", value)
	}

}
