package basicfunction

func PrintMyResult(sentence string) string {
	newSentence := sentence + "Belajar Golang di Udemy"
	return newSentence
}

func Add(number, numberTwo int) int {
	return number * numberTwo
	//return result
}

//function multiple return

func Calculation(panjang int64, lebar int64) (int64, int64) {
	//var con int
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// function dengan PREDIFINED RETURN VALUE
func Calculation1(panjang int64, lebar int64) (luas int64, keliling int64) {
	//var con int
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}
