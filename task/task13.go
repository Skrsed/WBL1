package task

import (
	"fmt"
)

// Поменять местами два числа без создания временной переменной.
func UntemporarlySwap() {
	a := 7
	b := 8

	a = a - b
	b = a + b
	a = b - a

	fmt.Println(a, b)
}
