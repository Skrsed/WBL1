package task

import "fmt"

func binSearch[T Number](numSlice []T, needle T) int {
	low := 0
	high := len(numSlice) - 1

	for low <= high {
		mid := low + (high-low)/2

		if needle == numSlice[mid] {
			return mid
		}

		if needle < numSlice[mid] {
			high = mid - 1
			continue
		}

		low = mid + 1
	}

	return -1
}

// Реализовать бинарный поиск встроенными методами языка.
func BinarySearch() {
	numSlice := []int{1, 2, 4, 5, 6, 8, 9, 10, 13, 17}
	res := []int{}

	for _, val := range numSlice {
		res = append(res, binSearch(numSlice, val))
	}

	fmt.Println(res)

}
