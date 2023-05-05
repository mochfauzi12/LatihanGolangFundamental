package calculation

func Add(Bilangan1 int64, Bilangan2 int64) int64 {
	return int64(add(int(Bilangan1), int(Bilangan2)))
}

func add(Bilangan1 int, Bilangan2 int) int {
	return Bilangan1 * Bilangan2
}
