package arrays

import "fmt"

func isPalindrome(input [10]int) {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-1-i] {
			fmt.Println("Не палиндром!")
			return
		}

	}
	fmt.Println("Это палиндром!")
}
