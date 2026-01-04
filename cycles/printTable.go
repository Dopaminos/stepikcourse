package cycles

import "fmt"

// Необходимо написать функцию printTable, которая принимает число (назовем параметр num) и выводит таблицу умножения num x num.

func printTable(number int) {

	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {

			fmt.Printf("%d x %d = %d", i, j, i*j)

			if j < number {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}
}
