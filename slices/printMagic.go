package slices

import (
	"fmt"
	"strconv"
	"strings"
)

func printMagic(input []int) string {
	if len(input) == 0 {
		fmt.Println("[]")
		return "[]"
	}

	prefix := make([]int, len(input))
	suffix := make([]int, len(input))
	result := make([]int, len(input))

	prefix[0] = 1
	suffix[len(input)-1] = 1

	for i := 1; i < len(input); i++ {
		prefix[i] = prefix[i-1] * input[i-1]
	}

	for i := len(input) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * input[i+1]
	}

	for i := range input {
		result[i] = prefix[i] * suffix[i]
	}

	parts := []string{}
	for _, v := range result {
		parts = append(parts, strconv.Itoa(v))
	}
	fmt.Println("[" + strings.Join(parts, ", ") + "]")

	return "[" + strings.Join(parts, ", ") + "]"
}
