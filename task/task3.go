package task

import "fmt"

const N = 10

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
