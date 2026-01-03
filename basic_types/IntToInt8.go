package basictypes

import "fmt"

// Ваша задача — создать новую переменную i типа int и присвоить ей значение, которое хранится в переменной i8.
func IntToInt8() {
	var i8 int8
	var i int = int(i8)

	fmt.Printf("i's (%d) type changed to %T\n", i, i)
}
