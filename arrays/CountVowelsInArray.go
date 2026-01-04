package arrays

import (
	"fmt"
	"strings"
	"unicode"
)

// Необходимо реализовать функцию CountVowelsInArray, принимающую массив из 3 строк на русском языке.
// Функция должна вывести в консоль 3 целых числа через пробел, каждое число это результат подсчета количество гласных букв в соответствующей строке исходного массива.
// Гласными буквами считаются: а, е, ё, и, о, у, ы, э, ю, я (в нижнем и верхнем регистре).

func CountVowelsInArray(input [3]string) {
	vowels := "аеёиоуыэюя"

	for _, str := range input {
		count := 0

		for _, r := range str {
			rLower := unicode.ToLower(r)

			if strings.Contains(vowels, string(rLower)) {
				count++
			}
		}
		fmt.Print(count, " ")
	}
}
