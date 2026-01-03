package basictypes

import (
	"fmt"
	"unicode/utf8"
)

// У нас уже создана переменная str с типом string. Необходимо вывести в одном fmt.Println() несколько значений:

// Значение переменной str.
// Количество байт строки str.
// Количество символов строки str.
var str string = "hello"

func AboutString() {
	fmt.Println(str, len(str), utf8.RuneCountInString(str))
}
