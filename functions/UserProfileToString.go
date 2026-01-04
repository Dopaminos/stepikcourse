package functions

import (
	"fmt"
	"strings"
)

// Напишите функцию UserProfileToString, которая принимает имя пользователя name (string) и его возраст age (int).
// Функция должна возвращать строку с сообщением (string) и ошибку (error), если таковая была.
// Сообщение должно быть в формате: Имя человека: [ИМЯ], возраст: [ВОЗРАСТ].
// Если имя было передано в функцию и имеет пробелы слева/справа - эти пробелы нужно убрать.
func UserProfileToString(name string, age int) (string, error) {
	if len(name) == 0 {
		return "", fmt.Errorf("empty name")
	}

	name = strings.TrimSpace(name)

	if len(name) == 0 {
		return "", fmt.Errorf("name cannot contain only spaces")
	}

	if age < 0 {
		return "", fmt.Errorf("negative age")
	}

	profile := fmt.Sprintf("Имя человека: %s, возраст: %d.", name, age)
	return profile, nil
}
