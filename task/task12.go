package task

import (
	"fmt"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

func MakeSet[T comparable](items []T) Set[T] {
	data := map[T]struct{}{}
	for _, item := range items {
		data[item] = struct{}{}
	}

	return Set[T]{
		data: data,
	}
}

func (s *Set[T]) GetSlice() []T {
	res := make([]T, 0, len(s.data))

	for val := range s.data {
		res = append(res, val)
	}

	return res
}

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func OwnSet() {
	items := []string{"cat", "cat", "dog", "cat", "tree"}
	set := MakeSet(items)

	fmt.Println(set.GetSlice())
}
