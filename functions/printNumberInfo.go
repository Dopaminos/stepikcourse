package functions

// Необходимо лучше изучить число. Необходимо создать функцию printNumberInfo, которая принимает целое число и выводит информацию об этом числе.
import (
	"fmt"
	"math"
)

func printNumberInfo(num int) {
	if num < 0 {
		fmt.Printf("Число %d отрицательное.\n", num)
	} else if num > 0 {
		fmt.Printf("Число %d положительное.\n", num)
	} else {
		fmt.Println("Число равно 0.")
	}

	if num%2 == 0 {
		fmt.Printf("Число %d четное.\n", num)
	} else {
		fmt.Printf("Число %d нечетное.\n", num)
	}

	if num > 0 {
		sqrt := math.Sqrt(float64(num))
		intSqrt := int(sqrt)
		if sqrt == float64(intSqrt) {
			fmt.Printf("Квадратный корень числа %d является целым числом и равен %d.\n", num, intSqrt)
		} else {
			fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, sqrt)
		}
	}
}
