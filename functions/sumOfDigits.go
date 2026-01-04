package functions

import "math"

func sumOfDigits(n int) int {
	n = int(math.Abs(float64(n)))
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
