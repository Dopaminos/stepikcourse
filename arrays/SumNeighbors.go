package arrays

func SumNeighbors(input [10]int) [10]int {
	var newArray [10]int

	for i := 0; i < len(input); i++ {
		var left, right int

		if i > 0 {
			left = input[i-1]
		} else {
			left = 0
		}

		if i < len(input)-1 {
			right = input[i+1]
		} else {
			right = 0
		}

		newArray[i] = left + right
	}
	return newArray
}
