package operators

import "fmt"

// Вам дано целое число, в переменной num, необходимо выполнить с ним следующие действия (результат каждого действия необходимо вывести в консоль):

// Сдвинуть число влево на 2 бита.
// Сдвинуть число вправо на 1 бит.
// Использовать оператор побитового "И" с числом 3.
// Использовать оператор побитового "ИЛИ" с числом 2.
// Использовать оператор побитового "XOR" с числом 2.
// Инвертировать биты.
var num int = 2222

func BitOperation() {
	shiftLeft := num << 2
	shiftRight := num >> 1
	bitAnd := num & 3
	bitOr := num | 2
	bitXor := num ^ 2
	invertBits := ^num

	fmt.Println(shiftLeft)
	fmt.Println(shiftRight)
	fmt.Println(bitAnd)
	fmt.Println(bitOr)
	fmt.Println(bitXor)
	fmt.Println(invertBits)
}
