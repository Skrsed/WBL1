package task

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	count int
	m     sync.Mutex
	wg    *sync.WaitGroup
}

func (ac *SafeCounter) Increment() {
	defer ac.wg.Done()

	ac.m.Lock()
	defer ac.m.Unlock()

	ac.count++
}

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func ParallelIncrementation() {
	wg := sync.WaitGroup{}
	n := 5

	ac := SafeCounter{
		count: 0,
		wg:    &wg,
	}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go ac.Increment()
	}

	wg.Wait()
	fmt.Println(ac.count)
}
