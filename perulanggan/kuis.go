package perulanggan

import "fmt"

func Kuis() {

	title := "Golang is the best language"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index:", index, "Letter:", string(letter))
	// 	}
	// }

	//mencari huruf vokal pada sebuah kalimat
	for index, letter := range title {
		letterString := string(letter)

		// 	if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 		fmt.Println("index :", index, "Letter:", string(letter))
		// 	}

		// }

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("Index:", index, "Letter:", string(letter))
		}

	}
}
