package task

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
func MarcoPolo(n int) {
	timer := time.NewTimer(time.Second * time.Duration(n))
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(1)
	go sayPolo(ctx, ch, &wg)

	i := 0
loop:
	for {
		select {
		case <-timer.C:
			cancel()
			wg.Wait()
			break loop
		default:
		}

		iStr := strconv.Itoa(i)
		fmt.Println(iStr + " Marco!")
		ch <- iStr + " Polo"
		i++
	}

	fmt.Println("done!")
}

func sayPolo(ctx context.Context, ch chan string, wg *sync.WaitGroup) {
	for {
		select {
		case mes := <-ch:
			//time.Sleep(time.Second * 2)
			fmt.Println(mes)
		case <-ctx.Done():
			wg.Done()

			return
		}
	}
}
