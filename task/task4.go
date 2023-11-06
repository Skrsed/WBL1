package task

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func ChanelDataStream(n int) {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	ch := make(chan string)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt)

	go func() {
		fmt.Println("...ctrl + C to interrupt")
		<-sigs

		cancel()

		wg.Wait()
		fmt.Println("workerers was ends gracefully")

		os.Exit(0)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go listenMainThread(i, ch, ctx, &wg)
	}

	for {
		ch <- randStringRunes(10)
	}
}

func listenMainThread(chanI int, ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case mes := <-ch:
			fmt.Printf("listner %v received %s\n", chanI, mes)
		case <-ctx.Done():
			wg.Done()
			return
		}
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
