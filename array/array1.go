package array

import "fmt"

func Array1() {

	//BENTUK TERDAHULU
	// var language [6]string
	// language[0] = "Go"
	// language[1] = "Ruby"
	// language[2] = "Python"
	// language[3] = "C#"
	// language[4] = "Javascript"
	// language[5] = "Java"

	//BENTUK DENGAN MEDIFINISIKAN BERAPA JUMLAH ARRAY NYA
	// language := [6]string{"Golang",
	// 	"Ruby",
	// 	"Javascript",
	// 	"Java",
	// 	"PHP",
	// 	"Python",
	// }

	language := [...]string{"Golang",
		"Ruby",
		"Javascript",
		"Java",
		"PHP",
		"Python",
		"Mega Chan",
	}

	//   FOR RANGE UNTUK MENGAMBIL NILAI ARRAY
	for index, lang := range language {
		fmt.Println("Index :", index, "Language :", string(lang))
	}

	// fmt.Println(language)
	// length := len(language)
	// fmt.Println(length)
}
