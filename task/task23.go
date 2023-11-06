package task

import "fmt"

func deleteFromSlice[T any](slice []T, i int) []T {
	// controversial step
	newSlice := make([]T, 0, len(slice)-1)
	leftHalf := append(newSlice, slice[:i]...)
	return append(leftHalf, slice[i+1:]...)
}

// some kind of fast deletion
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Удалить i-ый элемент из слайса.
func DeleteItem() {
	someSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13}
	sliceWithoutItem := deleteFromSlice(someSlice, 3)

	fmt.Println(sliceWithoutItem[:], len(sliceWithoutItem), cap(sliceWithoutItem))
}
