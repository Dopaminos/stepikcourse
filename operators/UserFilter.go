package operators

import "fmt"

// Необходимо написать код, который на основе возраста пользователя, роли и его статуса подписки (активный или неактивный) выводит, может ли пользователь получить доступ к определенному контенту.

// В программе у нас уже есть полученные данные, которые лежат в переменных:

// age - целое число, переменная хранит возраст пользователя.
// role - строка, роль пользователя на сайте, значения могут быть: "admin", "moderator", "user".
// status - строка, статус подписки, значения могут быть: "active", "inactive", "paused".
var age int = 19
var role string = "admin"
var status string = "active"

func UserFilter() {
	var result bool

	if role == "admin" || role == "moderator" {
		result = true
	} else if status == "active" && age >= 18 && role == "user" {
		result = true
	}
	fmt.Println(result)
}
