package sliceandmap

import "fmt"

func Slice_Map() {
	students := []map[string]string{

		{"name": "Tanjiro Kamado", "Score": "A"},
		{"name": "Uzui Tengen", "Score": "B"},
		{"name": "Moichiro Tokito", "Score": "C"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["Score"])

	}

	//fmt.Printf("====================================================================================")

}
