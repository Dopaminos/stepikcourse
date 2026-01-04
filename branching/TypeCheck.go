package branching

import "fmt"

var val any = 2

func TypeCheck() {
	switch v := val.(type) {
	case int, float64, bool, string:
		fmt.Printf("В переменной val находится тип %T.", v)
	default:
		fmt.Println("В переменной val находится неизвестный тип данных.")
	}
}
