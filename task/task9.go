package task

import (
	"fmt"
	"os"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные из
// второго канала должны выводиться в stdout.
func NumsConveyor() {
	arr := [8]int{1, 3, 5, 6, 7, 9, 11, 13}
	arrValues := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for _, v := range arr {
			arrValues <- v
		}
		close(arrValues)
	}()

	// Squarer
	go func() {
		for x := range arrValues {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Fprintln(os.Stdout, x)
	}
}
