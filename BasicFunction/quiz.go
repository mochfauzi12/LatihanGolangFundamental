package basicfunction

import (
	"errors"
	"fmt"
)

//Soal dan pembahasan

func main() {
	//soal 1
	scores := []int{10, 5, 8, 9, 7}
	total := Sum(scores)
	fmt.Println(total)

	//soal 2
	//result, err := calculate(10,2,"+")
	//result, err := calculate(10,2,"-")
	//result, err := calculate(10,2,"/")
	//result, err := calculate(10,2,"=")

}

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}

	return total
}

func Calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	case "%":
		result = number % numberTwo
	default:
		errorResult = errors.New("invalid operation")
	}
	return result, errorResult
}
