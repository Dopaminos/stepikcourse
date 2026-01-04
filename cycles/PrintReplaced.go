package cycles

import "fmt"

// Напишите функцию PrintReplaced, которая принимает строку и выводит в консоль новую строку, в которой все буквы "у" заменены на буквы "а" (регистрозависимо, прописные буквы игнорируем).
// Так как мы изучаем циклы, хотелось бы, чтобы и решение было с использованием циклов, а не готовых функций.
func PrintReplaced(sample string) {
	var result string
	for _, char := range sample {
		if char == 'у' {
			result += "а"
		} else {
			result += string(char)
		}
	}

	fmt.Println(result)
}
