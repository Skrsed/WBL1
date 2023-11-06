package task

import (
	"fmt"
	"os"
)

func Concurency1() {
	nums := []int{2, 4, 6, 8, 10}
	res := make([]int, 0, len(nums))

	ch := make(chan int, len(nums))
	for _, num := range nums {
		go square(num, ch)
	}

	for i := 0; i < cap(ch); i++ {
		res = append(res, <-ch)
	}

	fmt.Fprintln(os.Stdout, res)
}

func square(x int, ch chan<- int) {
	ch <- x * x
}
