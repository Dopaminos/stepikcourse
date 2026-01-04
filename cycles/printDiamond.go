package cycles

import "fmt"

// Хочется увидеть ромбик в терминале, для этого необходимо создать функцию func printDiamond(n int), которая выводит текст ]
// Мой бриллиант:, а с новой строки ромб (смотри примеры), где n - это ребро ромба.

func printDiamond(n int) {
	fmt.Println("Мой бриллиант:")

	// функция формирования строку ромба для уровня i
	printRow := func(i int) {
		leftSpaces := ""
		for s := 0; s < n-1-i; s++ {
			leftSpaces += " "
		}

		if i == 0 {
			fmt.Println(leftSpaces + "#")
		} else {
			middleSpaces := ""
			for s := 0; s < i*2-1; s++ {
				middleSpaces += " "
			}
			fmt.Println(leftSpaces + "#" + middleSpaces + "#")
		}
	}

	// верхняя часть
	for i := 0; i < n; i++ {
		printRow(i)
	}
	// нижняя часть
	for i := n - 2; i >= 0; i-- {
		printRow(i)
	}
}
