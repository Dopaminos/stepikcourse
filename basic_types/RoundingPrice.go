package basictypes

import (
	"fmt"
	"math"
)

//     discountPercent — тип float64, процент скидки, который нужно применить к цене товара.
//     productPrice — тип float64, представляющий цену текущего товара.

// Необходимо будет округлить значение, чтобы избежать длинной дробной части, округление должно быть в пользу покупателя (то есть округляйте в меньшую сторону),
// максимально может быть лишь два знака в дробной части (смотри примеры в тестовых данных).
var discountPercent float64 = 10.0
var productPrice float64 = 99.9

func RoundingPrice(float64, float64) {

	finalPrice := math.Floor((productPrice-((discountPercent/100)*productPrice))*100) / 100

	fmt.Println(finalPrice)
}
