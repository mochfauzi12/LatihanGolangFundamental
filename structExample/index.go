package structexample

import "fmt"

type Students struct {
	Id        int
	Firstname string
	Lastname  string
	Email     string
	Address   string
	IsActive  bool
}

//embed struct example

type Groups struct {
	Name        string
	Admin       Students
	Students    []Students
	IsAvailable bool
}

func Tampil() {

	// student := Students{
	// 	Id:        1,
	// 	Firstname: "Taufik",
	// 	Lastname:  "Fauzi",
	// 	Email:     "Taufiksavalas@gmail.com",
	// 	Address:   "Salatiga Kota ",
	// 	IsActive:  false,
	// }

	// student1 := Students{
	// 	Id:        0,
	// 	Firstname: "ananda",
	// 	Lastname:  "putri",
	// 	Email:     "ananda@gmail.com",
	// 	Address:   "Jalan raya demangan , eromoko",
	// 	IsActive:  false,
	// }

	student := Students{1, "Maman", "Surahman", "maamansurahman@gmail.com", "Demangan City", true}
	student1 := Students{2, "Maman", "Abdul", "Abdulrohman@gmail.com", "Ngungahan City", true}
	displayStudent1 := DisplayStudent(student)
	displayStudent2 := DisplayStudent(student1)
	fmt.Println(displayStudent1)
	fmt.Println(displayStudent2)

}

func DisplayStudent(student Students) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", student.Firstname, student.Lastname, student.Email)

}
