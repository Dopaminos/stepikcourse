package basictypes

import "fmt"

// Мы находимся в функции main. У нас уже есть переменная message типа string, необходимо выполнить следующие действия с данной строкой:

// Удалить пробелы слева и справа от строки message, если такие имеются, и вывести полученную строку.
// Преобразовать строку, полученную в пункте 1, в нижний регистр и вывести её.
// Вывести true, если строка, полученная в пункте 2, начинается с букв "go", в противном случае вывести false.
var message string = " Go - это не просто язык, это СТИЛЬ ЖИЗНИ! "

func StringProcessing() {
	runes := []rune(message)
	start, end := 0, len(runes)-1

	for start <= end && runes[start] == ' ' {
		start++
	}

	for end >= start && runes[end] == ' ' {
		end--
	}

	trimmedMessage := string(runes[start : end+1])
	fmt.Println(trimmedMessage)

	// Конвертируем в нижний регистр вручную
	var lowerRunes []rune
	for _, r := range trimmedMessage {
		if r >= 'А' && r <= 'Я' || r == 'Ё' { // Добавляем проверку на Ё
			r += 'а' - 'А'
		} else if r >= 'A' && r <= 'Z' {
			r -= 'A' - 'a'
		}
		lowerRunes = append(lowerRunes, r)
	}

	lowercaseMessage := string(lowerRunes)
	fmt.Println(lowercaseMessage)

	// Проверяем, начинается ли строка с "go"
	beginsWithGo := len(lowercaseMessage) >= 2 && lowercaseMessage[:2] == "go"
	fmt.Println(beginsWithGo)
}
