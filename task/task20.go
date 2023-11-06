package task

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	lastWordIndex := len(words) - 1

	for i := 0; i < lastWordIndex/2; i++ {
		words[i], words[lastWordIndex-i] = words[lastWordIndex-i], words[i]
	}

	return strings.Join(words, " ")
}

// Разработать программу, которая переворачивает слова в строке.
func SwapWords() {
	sentece := "🤔 🤔 日本語 snow dog sun 🤔"

	res := reverseWords(sentece)

	fmt.Println(res)
}
