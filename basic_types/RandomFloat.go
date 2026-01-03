package basictypes

import (
	"fmt"
	"math/rand/v2"
)

var min float64 = 1.0
var max float64 = 10.0

// Ваша задача — сгенерировать случайное число с плавающей точкой в диапазоне от min до max, не включая max, и сохранить его в переменной random типа float64.
func RandomFloat() {
	var random float64 = rand.Float64()*(max-min) + min
	fmt.Print(random)
}
