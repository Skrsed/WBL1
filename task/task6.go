package task

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
func StopGoroutines() {
	wg := sync.WaitGroup{}
	timer := time.NewTimer(time.Second)

	quit := make(chan struct{})
	wg.Add(1)
	go stopWithChannel(quit, &wg)

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go stopWithContext(ctx, &wg)

	wg.Add(1)
	go stopWithChannelClose(quit, &wg)

	wg.Add(1)
	go stopWithChannelClose2(quit, &wg)

	<-timer.C

	// using chan
	quit <- struct{}{}
	// using ctx
	cancel()
	// using range and closed chan
	close(quit)

	wg.Wait()
	fmt.Println("All goroutines are done")
}

func stopWithContext(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("stopWithContext closed, iterations - " + strconv.Itoa(i))
			wg.Done()

			return
		default:
		}
	}
}

func stopWithChannel(quit chan struct{}, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-quit:
			fmt.Println("stopWithChannel closed, iterations - " + strconv.Itoa(i))
			wg.Done()

			return
		default:
		}
	}
}

func stopWithChannelClose(ch chan struct{}, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		// there is
		_, ok := <-ch
		if !ok {
			fmt.Println("stopWithChannelClose closed, iterations - " + strconv.Itoa(i))
			wg.Done()

			return
		}
	}
}

func stopWithChannelClose2(ch chan struct{}, wg *sync.WaitGroup) {
	for range ch {
	}

	fmt.Println("stopWithChannelClose2 closed")
	wg.Done()
}
