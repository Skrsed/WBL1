package task

import (
	"fmt"
	"unicode"
)

func checkUnique(str string) bool {
	strMap := make(map[rune]struct{}, len(str))

	for _, symbol := range str {
		lowerS := unicode.ToLower(symbol)
		if _, ok := strMap[lowerS]; ok {
			return false
		}
		strMap[lowerS] = struct{}{}
	}

	return true
}

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
func AllUnique() {
	str := "qwertyuioP"
	isUnique := checkUnique(str)

	fmt.Println(isUnique)
}
