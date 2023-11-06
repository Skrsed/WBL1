package task

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)

	<-timer.C
}

// Реализовать собственную функцию sleep.
func MySleep() {
	fmt.Println("Starts mine sleeping")
	start := time.Now()
	sleep(time.Second * 3)
	fmt.Println("Done in ", time.Since(start))
}
