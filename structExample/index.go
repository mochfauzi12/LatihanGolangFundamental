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

func Tampil() {

	student := Students{
		Id:        1,
		Firstname: "Taufik",
		Lastname:  "Fauzi",
		Email:     "Taufiksavalas@gmail.com",
		Address:   "Salatiga Kota ",
		IsActive:  false,
	}

	student1 := Students{}
	student1.Id = 2
	student1.Firstname = "Taufik Bala Salatiga"
	student1.Lastname = "Salatiga "
	student1.Email = "taufik@gmail.com"
	student1.Address = "Jln.Raya salatiga, pasar kembangsari "
	student1.IsActive = true

	fmt.Println(student)
	fmt.Println(student1)
}
