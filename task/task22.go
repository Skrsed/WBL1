package task

import (
	"fmt"
	"math"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.
func LongArithmetic() {
	num := int64(math.Pow(2, 20)) + 1
	a := new(big.Int).SetInt64(num)
	b := new(big.Int).SetInt64(num*2 + 117)
	x := new(big.Int)
	fx := new(big.Float)
	fx.SetPrec(128) // bits

	fmt.Println("Divide a to b", fx.Quo(
		new(big.Float).SetInt(a),
		new(big.Float).SetInt(b),
	))
	fmt.Println("Multiply ", x.Mul(a, b))
	fmt.Println("Sum ", x.Add(a, b))
	fmt.Println("Diff ", x.Sub(a, b))
}
