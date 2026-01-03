package basictypes

import "fmt"

// Поскольку наш магазин не взимает копейки с покупателей, ваша задача заключается в том, чтобы вывести в консоль только целую часть значения из переменной totalPrice.
func PrintIntegerPart() {
	var totalPrice float64 = 552.64
	fmt.Println(int(totalPrice))
}
