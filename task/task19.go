package task

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
func ReverseString() {
	str := "🐹 🐅🐅🐅"

	out := []rune(str)
	lastI := len(out) - 1
	for i := 0; i <= lastI/2; i++ {
		out[i], out[lastI-i] = out[lastI-i], out[i]
	}

	fmt.Println(string(out))
}
