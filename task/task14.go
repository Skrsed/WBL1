package task

import (
	"fmt"
	"reflect"
)

func checkType[T any](val T) reflect.Type {
	return reflect.TypeOf(val)
}

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}
func CheckType() {
	intV := 11
	strV := "STR"
	boolV := true
	chanV := make(chan interface{})

	fmt.Println(checkType(intV))
	fmt.Println(checkType(strV))
	fmt.Println(checkType(boolV))
	fmt.Println(checkType(chanV))
}
