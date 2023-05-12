package main

import (
	basicfunction "LatihanGolangFundamental/BasicFunction"
	"LatihanGolangFundamental/map1"
	sliceandmap "LatihanGolangFundamental/sliceAndmap"
	structexample "LatihanGolangFundamental/structExample"
	"fmt"
)

// "LatihanGolangFundamental/calculation"
// "LatihanGolangFundamental/perkalian"

func main() {
	// fmt.Println("Bahasa Pemograman Golang")

	// result := calculation.Add(100, 3000000000)

	// fmt.Println("Ini adalah hasil perkalian")

	//result1 := perkalian.AddPerkalian(1000, 2000000)

	// fmt.Println(result)
	// fmt.Println(result1)

	// Muzan := pembagian.Pembagian(1200000, 300)

	// fmt.Println(Muzan)
	// var Age int64 = 400
	// Age = 500

	// fmt.Println(Age)

	// logikaifelse.Logic()
	// logikaifelse.Score()

	//switchcase.Switchcase1()

	//perulanggan.Looping()

	//perulanggan.Kuis()
	//array.Array1()

	//slice.Slice()

	//map1.Map()
	map1.Map1()

	//fmt.Println("====================================================================================")

	//sliceandmap.Slice_Map()

	sliceandmap.Quiz()
	fmt.Println("====================================================================================")

	//basicfunction.PrintMyResult("Saya Senang")
	sentence := basicfunction.PrintMyResult("Saya Senang ")
	fmt.Println(sentence)

	fmt.Println("====================================================================================")

	result1 := basicfunction.Add(45, 67)
	fmt.Println(result1)

	fmt.Println("====================================================================================")

	// luas1, keliling1 := basicfunction.Calculation(10, 2)

	// fmt.Println(luas1)
	// fmt.Println(keliling1)

	//basicfunction.Sum()

	result, err := basicfunction.Calculate(10, 2, "%")
	if err != nil {

		fmt.Println("Terjadi Error")
		fmt.Println(err.Error())

	}

	fmt.Println(result)

	fmt.Println("====================================================================================")

	structexample.Tampil()

}
