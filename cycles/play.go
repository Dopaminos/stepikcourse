package cycles

// https://stepik.org/lesson/1500232/step/5?unit=1520340
// desc is too long
import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

func game() {
	for range 100 {
		random = rand.IntN(100) + 1
		guesses = 0
		result := play()
		if result != random {
			fmt.Printf("Неверный ответ. Было загадано число %d, а в ответе получили число %d", random, result)
			os.Exit(-1)
		}
	}
}

var guesses int
var random int

func guess(num int) (int, error) {
	if guesses >= 6 {
		return 0, errors.New("too many attempts")
	}
	guesses++
	if num > random {
		return -1, nil
	}
	if num < random {
		return 1, nil
	}
	return 0, nil
}

func play() int {
	var leftPointer int = 1
	var rightPointer int = 100

	for leftPointer <= rightPointer {
		var midPointer = (leftPointer + rightPointer) / 2

		result, err := guess(midPointer)

		if err != nil {
			return midPointer
		}

		switch result {
		case 0:
			return midPointer
		case -1:
			rightPointer = midPointer - 1
		case 1:
			leftPointer = midPointer + 1
		}
	}

	return leftPointer
}
