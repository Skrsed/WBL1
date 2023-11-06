package task

import (
	"encoding/binary"
	"fmt"
	"math"
	"unsafe"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func XorInt64(n int) {
	int64ByteLen := 8
	x := 42354332
	littleEndianMask := int64(math.Pow(2, float64(n)))

	leMaskRep := *(*[8]byte)(unsafe.Pointer(&littleEndianMask))
	xRepresentation := *(*[8]byte)(unsafe.Pointer(&x))

	applyBE := make([]byte, len(leMaskRep))
	applyLE := make([]byte, len(leMaskRep))
	for i := 0; i < int64ByteLen; i++ {
		applyBE[i] = leMaskRep[int64ByteLen-i-1] ^ xRepresentation[i]
		applyLE[i] = leMaskRep[i] ^ xRepresentation[i]
	}

	fmt.Printf("%08b\n", xRepresentation)
	fmt.Printf("%08b\n", applyBE)
	fmt.Printf("%08b\n", applyLE)

	fmt.Println(binary.LittleEndian.Uint64(applyBE))
	fmt.Println(binary.LittleEndian.Uint64(applyLE))
}
