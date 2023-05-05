package logikaifelse

import "fmt"

func Score() {

	score := 65
	var grade string

	if score <= 50 {
		grade = "E"

	} else if score <= 60 {
		grade = "D"

	} else if score <= 70 {
		grade = "C"

	} else {
		grade = "NULL"
	}

	fmt.Println(grade)

}
