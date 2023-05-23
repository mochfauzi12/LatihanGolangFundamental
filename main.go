package main

import (
	"LatihanGolangFundamental/pointer"
	"fmt"
)

// basicfunction "LatihanGolangFundamental/BasicFunction"
// "LatihanGolangFundamental/map1"
// sliceandmap "LatihanGolangFundamental/sliceAndmap"
// structexample "LatihanGolangFundamental/structExample"
// "fmt"

// "LatihanGolangFundamental/calculation"
// "LatihanGolangFundamental/perkalian"

type Students struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
	Address   string
	IsActive  bool
}

// embed struct example
type Groups struct {
	Name        string
	Admin       Students
	Students    []Students
	IsAvailable bool
}

func (group Groups) DisplayStudent() {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count: %d\n", len(group.Students))

	fmt.Println("Name of Member :")

	for _, student := range group.Students {
		fmt.Println(student.Firstname)
	}

}

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
	//map1.Map1()

	//fmt.Println("====================================================================================")

	//sliceandmap.Slice_Map()

	// sliceandmap.Quiz()
	// fmt.Println("====================================================================================")

	//basicfunction.PrintMyResult("Saya Senang")
	// sentence := basicfunction.PrintMyResult("Saya Senang ")
	// fmt.Println(sentence)

	// fmt.Println("====================================================================================")

	// result1 := basicfunction.Add(45, 67)
	// fmt.Println(result1)

	// fmt.Println("====================================================================================")

	// luas1, keliling1 := basicfunction.Calculation(10, 2)

	// fmt.Println(luas1)
	// fmt.Println(keliling1)

	//basicfunction.Sum()

	// result, err := basicfunction.Calculate(10, 2, "%")
	// if err != nil {

	// 	fmt.Println("Terjadi Error")
	// 	fmt.Println(err.Error())

	// }

	// fmt.Println(result)

	// fmt.Println("====================================================================================")

	// structexample.Tampil()
	//structexample.DisplayStudent(structexample.students)

	// student := Students{1, "Maman", "Surahman", "maamansurahman@gmail.com", "Demangan City", true}
	// student1 := Students{2, "Maman", "Abdul", "Abdulrohman@gmail.com", "Ngungahan City", true}

	// students := []Students{student, student1}

	// group := Groups{"Cracker", student, students, true}

	// DisplayGroup(group)

	pointer.Pointer1()
}

func DisplayGroup(group Groups) {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count: %d\n", len(group.Students))

	fmt.Println("Name of Member :")

	for _, student := range group.Students {
		fmt.Println(student.Firstname)
	}

}

// method
func (student Students) Display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", student.Firstname, student.Lastname, student.Email)

}

// function
func DisplayStudent(student Students) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", student.Firstname, student.Lastname, student.Email)

}
