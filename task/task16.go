package task

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// https://ru.wikipedia.org/wiki/%D0%91%D1%8B%D1%81%D1%82%D1%80%D0%B0%D1%8F_%D1%81%D0%BE%D1%80%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B0
// Схема Хоара
func partition[T Number](A []T, low int, high int) int {
	pivot := A[(low+high)/2]
	i := low
	j := high

	for {
		for ; A[i] < pivot; i++ {
		}
		for ; A[j] > pivot; j-- {
		}
		if i >= j {
			return j
		}

		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}

func quickSort[T Number](A []T, low int, high int) {
	if low >= high {
		return
	}

	p := partition[T](A, low, high)
	quickSort[T](A, low, p)
	quickSort[T](A, p+1, high)
}

// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.
func QuickSort() {
	numSlice := []int{1, 5, 3, 7, 6, 0, 1}
	sliceLen := len(numSlice)

	sorted := make([]int, sliceLen)
	copy(sorted, numSlice)
	quickSort(sorted, 0, sliceLen-1)

	fmt.Println("sorted", sorted)
}
