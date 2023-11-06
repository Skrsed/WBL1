package task

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable] struct {
	acc map[K]int
	m   sync.Mutex
	wg  *sync.WaitGroup
}

func (smap *SafeMap[K]) WriteMap(key K) {
	defer smap.wg.Done()

	smap.m.Lock()
	defer smap.m.Unlock()

	if _, ok := smap.acc[key]; !ok {
		smap.acc[key] = 0
	}

	smap.acc[key]++
}

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func ParallelMapWriter() {
	wg := sync.WaitGroup{}
	n := 5

	safeMap := SafeMap[string]{
		acc: map[string]int{},
		wg:  &wg,
	}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go safeMap.WriteMap("123")
	}

	wg.Wait()
	fmt.Println(safeMap.acc)
}
