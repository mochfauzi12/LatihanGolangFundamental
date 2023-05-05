package perulanggan

import "fmt"

func Looping() {
	// count := 10000
	// for i := 1; i <= count; i++ {
	// 	fmt.Println("Cetak nilai ini :", i)
	// }

	//for range dalam golang beserta aturannya
	squad := "Hinokami Kagura Api firewall"

	for index, letter := range squad {

		fmt.Println("index:", index, "Letter :", string(letter))
	}

	// title := "Golang The best Language"

	// for _, letter := range title {
	// 	fmt.Println("letter:", string(letter))
	// }
}
