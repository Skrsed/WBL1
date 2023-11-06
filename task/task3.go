package task

import "fmt"

const N = 10

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.
func Concurency2() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 2; x <= N; x += 2 {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	res := 0
	for x := range squares {
		res += x
	}

	fmt.Println("r=", res)
}
