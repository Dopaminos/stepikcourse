package functions

// Вам необходимо реализовать функцию userProfile, которая будет обрабатывать информацию о пользователе на основе его идентификатора.

import (
	"fmt"
)

var id string = "2"

func userProfile(id string) (string, error) {
	balanceKopecks, err := fetchUserInfo(id)
	if err != nil {
		return "", fmt.Errorf("fetch error: %w", err)
	}

	balanceRubles := float64(balanceKopecks) / 100.0

	message := fmt.Sprintf("Пользователь с id %s имеет на счету %.2f руб.", id, balanceRubles)
	return message, nil
}

func fetchUserInfo(id string) (int, error) {
	if id == "2" {
		return 3322, nil
	}
	return 3322, nil
}
