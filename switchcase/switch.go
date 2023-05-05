package switchcase

import "fmt"

func Switchcase1() {
	number := 40

	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("Default tidak ada nilainya ")
	}

}
