package task

import (
	"fmt"
	"math/big"
)

var justString string

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?

// Ответ: символ не обязательно представляется в виде одного байта,
// поэтому мы можем случайно обрезать его так, что он перестанет отображаться,
// к тому же считать символы в байтах не удобно. Используем []rune slice, чтобы
// отделить нужное количество из строки.
func someFunc() {
	bigBufferWithHeader := []byte{}
	for i := 0; i < 1e15; i++ {
		bigBufferWithHeader = append(bigBufferWithHeader, 0)
	}
	bigBufferWithHeader[0] = 0
	v := createHugeString(string(1 << 100))
	n := 5
	brokenString := v[:n]

	// Допустим, что мы не хотим копировать всю строку
	runes := makeRuneSliceOfN(v, n)

	//alocateSomeMem()

	fmt.Println(len(runes), cap(runes))
	justString = string(runes[:n])

	fmt.Println(brokenString)
	fmt.Println(justString)
}

func makeRuneSliceOfN(v string, n int) []rune {
	runes := make([]rune, n)
	bigN := big.NewInt(int64(n))
	for i, symbol := range v {
		if bigN.Cmp(big.NewInt(int64(i))) > 0 {
			break
		}
		runes[i] = symbol
	}

	return runes
}

func SomeBadThing() {
	someFunc()
}

func createHugeString(n string) string {

	var s string
	one := big.NewInt(1)
	bigN, _ := new(big.Int).SetString(n, 10)
	for i := big.NewInt(0); i.Cmp(bigN) < 0; i.Add(i, one) {
		s += "🥰 - OwO"
	}
	return s
}
