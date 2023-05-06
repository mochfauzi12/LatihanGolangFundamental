package slice

import "fmt"

func Slice() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "Playstation 4")

	gamingConsoles = append(gamingConsoles, "XBOX")

	gamingConsoles = append(gamingConsoles, "Nitendo Switch")

	//fmt.Println(gamingConsoles)

	for _, index := range gamingConsoles {
		fmt.Println(index)
	}
}
