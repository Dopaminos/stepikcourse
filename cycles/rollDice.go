package cycles

import (
	"fmt"
	"math/rand/v2"
)

// Необходимо симулировать броски двух шестигранных кубиков, для этого необходимо реализовать функцию rollDice.
// Данная функция должна принимать одно целое число, ничего при этом не возвращая.

// В функцию будет передаваться число в диапазоне от 2 до 12 включительно (число всегда передается в правильном диапазоне, проверять не нужно),
// которое будет представлять целевую сумму, которую необходимо получить при бросках кубиков.

func rollDice(dice int) {
	rollCounter := 0

	for {
		diceOne := rand.IntN(6) + 1
		diceTwo := rand.IntN(6) + 1
		rollCounter++

		sum := diceOne + diceTwo

		if sum == dice {
			fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d %s.\n",
				diceOne, diceTwo, sum, rollCounter, pluralRolls(rollCounter))
			break
		} else {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n",
				diceOne, diceTwo, sum)
		}
	}
}

func pluralRolls(n int) string {
	if n >= 10 && n <= 20 {
		return "бросков"
	}

	lastDigit := n % 10
	switch lastDigit {
	case 1:
		return "бросок"
	case 2, 3, 4:
		return "броска"
	default:
		return "бросков"
	}
}
