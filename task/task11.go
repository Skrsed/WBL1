package task

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.
func UGroup() {
	var group1 = []int{1, 7, 4, 3, 6, 8, 5, 2, 9, 10}
	var group2 = []int{8, 5, 2, 9, 10, 13, 14, 11}

	var concatMap = make(map[int]bool, len(group1))
	var res = make([]int, 0, len(group2))
	for _, val := range group1 {
		concatMap[val] = true
	}
	for _, val := range group2 {
		if concatMap[val] {
			res = append(res, val)
		}
	}

	fmt.Println(res)
}
