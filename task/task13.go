package task

type Swap struct {
	a int
	b string
}

// func untemporarlySwap() {
// 	p1 := Swap{a: 1, b: "test"}
// 	p2 := Swap{a: 2, b: "tset"}

// 	pp1 := (*(*[unsafe.Sizeof(p1)]byte)(unsafe.Pointer(&p1)))[:]
// 	pp2 := (*(*[unsafe.Sizeof(p2)]byte)(unsafe.Pointer(&p2)))[:]

// 	p1 = Swap(pp1 ^ pp2)
// }
