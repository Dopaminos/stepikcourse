package basictypes

import (
	"fmt"
	"strconv"
)

var priceStr string = "22.22"
var quantityStr string = "5"

// Принимает две строки, в переменных priceStr и quantityStr. Внутри этих строк записаны числа:
//     priceStr — цена продукта (дробное число).
//     quantityStr — количество продуктов (целое число).

// Сервис должен вычислить общую сумму, которая равна произведению цены на количество.

// Результат (общую сумму) необходимо вывести в консоль.

func StringsToFloat() {
	price, _ := strconv.ParseFloat(priceStr, 64)
	quantity, _ := strconv.Atoi(quantityStr)

	fmt.Printf("%.2f\n", price*float64(quantity))
}
