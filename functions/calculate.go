package functions

import (
	"errors"
)

func calculate(a, b float64, op string) (result float64, err error) {
	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		if b == 0 {
			err = errors.New("division by zero")
			break
		}
		result = a / b
	default:
		err = errors.New("unknown operation")
	}

	return result, err
}
